package main

import (
	"fmt"
	"net/http"
)

func init() { // HL
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
