package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
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

func main() {
	wg := sync.WaitGroup{} // HL
	for i := 0; i < 10; i++ {
		wg.Add(1)        // HL
		go func(i int) { // HL
			body, err := Fetch("http://localhost:7000")
			if err != nil {
				fmt.Printf("%v. Err: %v\n", i, err)
			} else {
				fmt.Printf("%v. Body: %v\n", i, body)
			}
			defer wg.Done() // HL
		}(i) // HL
	}
	wg.Wait() // HL
}

// END OMIT
