package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var connections int32 = 0

func delayHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%d connections", atomic.AddInt32(&connections, 1))
	defer func() {
		log.Printf("%d connections", atomic.AddInt32(&connections, -1))
	}()

	values, ok := r.URL.Query()["delay"]
	if ok {
		if dur, err := time.ParseDuration(values[0]); err == nil {
			time.Sleep(dur)
		}
	}
	fmt.Fprintf(w, "Hello, World!\n")
}

func inspectHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "UNKNOWN"
	}
	fmt.Fprintf(w, "%d open connections on %s!\n", connections, hostname)
}

func main() {
	port := os.Getenv("PORT")
	log.Printf("listening on localhost:%s...", port)
	http.HandleFunc("/hello", delayHandler)
	http.HandleFunc("/inspect", inspectHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
