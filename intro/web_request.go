package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// START OMIT
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

func main() {
	for i := 0; i < 10; i++ {
		body, err := Fetch("http://localhost:7000")
		if err != nil {
			fmt.Printf("%v. Err: %v\n", i, err)
		} else {
			fmt.Printf("%v. Body: %v\n", i, body)
		}
	}
}

// END OMIT
