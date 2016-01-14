package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

// START OMIT

func FetchWorker(in <-chan string, out chan<- string) {
	for url := range in {
		fmt.Println("Worker: fetch", url)
		body, err := Fetch(url)
		if err != nil {
			out <- err.Error()
		} else {
			out <- body
		}
	}
	close(out)
}

func main() {
	in := make(chan string, 10)
	out := make(chan string, 10)
	go FetchWorker(in, out)
	in <- "http://localhost:7000"
	in <- "http://localhost:7001"
	close(in)

	for msg := range out {
		fmt.Println("Main:", msg)
	}
}

// END OMIT
