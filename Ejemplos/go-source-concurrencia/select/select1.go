package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case val := <-ch1:
			fmt.Println("Received from ch1:", val)
		case val := <-ch2:
			fmt.Println("Received from ch2:", val)
		}
	}
}
