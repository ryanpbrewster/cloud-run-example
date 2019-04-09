package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const ADDRESS = ":80"

var OPEN_CONNECTIONS int32 = 0

func delayHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%d connections", atomic.AddInt32(&OPEN_CONNECTIONS, 1))
	defer func() {
		log.Printf("%d connections", atomic.AddInt32(&OPEN_CONNECTIONS, -1))
	}()

	values, ok := r.URL.Query()["delay"]
	if ok {
		if dur, err := time.ParseDuration(values[0]); err == nil {
			time.Sleep(dur)
		}
	}
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	log.Printf("listening on %s...", ADDRESS)
	http.HandleFunc("/hello", delayHandler)
	log.Fatal(http.ListenAndServe(ADDRESS, nil))
}
