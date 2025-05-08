package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func(out chan<- int) {
		for x := 0; x < 10; x++ {
			out <- x
		}
		close(out)
	}(naturals)

	// Squarer
	go func(in <-chan int, out chan<- int) {
		for x := range in {
			out <- x * x
		}
		close(out)
	}(naturals, squares)

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
