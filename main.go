package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"proxy/throttle"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	time.Sleep(20 * time.Millisecond)
}

func main() {
	c := flag.Uint64("c", 1, "Concurrency")
	flag.Parse()
	log.Printf("Running... with concurrency: %v\n", *c)
	log.Fatal(http.Serve(throttle.NewListener("0:8080", *c),
		http.HandlerFunc(HelloHandler)))
}
