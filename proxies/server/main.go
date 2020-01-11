package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v - %v\n", r.Proto, r.Method)
		log.Printf("Headers: %v\n", r.Header)
		fmt.Fprintln(w, "Welcome to my website!")
	})
	http.ListenAndServe(":4080", nil)
}
