package main

import "fmt"

// START OMIT
func Worker(in <-chan int, out chan<- string) { // HL
	for value := range in {
		fmt.Println("Worker:", value)
		out <- fmt.Sprintf("%v*%v = %v", value, value, value*value) // HL
	}
	close(out) // HL
}

func main() {
	in := make(chan int)     // HL
	out := make(chan string) // HL
	go Worker(in, out)

	go func() {
		in <- 1  // HL
		in <- 2  // HL
		in <- 15 // HL
		close(in)
	}()

	for msg := range out {
		fmt.Println("Anwser:", msg)
	}
}

// END OMIT
