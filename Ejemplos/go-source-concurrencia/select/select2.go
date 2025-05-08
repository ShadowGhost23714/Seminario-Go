package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			ch2 <- i
		}
		close(ch2)
	}()

	var val int
	ok1 := true
	ok2 := true
	for ok1 && ok2 {
		select {
		case val, ok1 = <-ch1:
			if ok1 {
				fmt.Println("Received from ch1:", val)
			}
		case val, ok2 = <-ch2:
			if ok2 {
				fmt.Println("\t\t\t\t\tReceived from ch2:", val)
			}
		}
	}
	if !ok2 {
		for val = range ch1 {
			fmt.Println("Received from ch1:", val)
		}
	}
	if !ok1 {
		for val = range ch2 {
			fmt.Println("\t\t\t\t\tReceived from ch2:", val)
		}
	}
}
