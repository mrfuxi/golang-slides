package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
