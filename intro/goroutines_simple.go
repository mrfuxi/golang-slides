// Example from: https://tour.golang.org/concurrency/1 OMIT
package main

import (
	"fmt"
	"time"
)

func echo(from string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(from, ":", i)
	}
}

func main() {
	go echo("123")
	echo("ABC")
}
