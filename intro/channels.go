package main

import "fmt"

// START OMIT
func Worker(in <-chan int, out chan<- string) {
	for value := range in {
		fmt.Println("Worker:", value)
		out <- fmt.Sprintf("%v*%v = %v", value, value, value*value)
	}
	close(out)
}

func main() {
	in := make(chan int)
	out := make(chan string)
	go Worker(in, out)

	go func() {
		in <- 1
		in <- 2
		in <- 15
		close(in)
	}()

	for msg := range out {
		fmt.Println("Anwser:", msg)
	}
}

// END OMIT
