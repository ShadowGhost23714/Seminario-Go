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
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
			ch1 <- 1
		}
		ch1 <- 0
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
			ch2 <- 2
		}
		ch2 <- 0
	}()

	fin := 0
	for fin < 2 {
		select {
		case val := <-ch1:
			fmt.Println("Received from ch1:", val)
			if val == 0 {
				fin++
			}
		case val := <-ch2:
			fmt.Println("\t\t\t\t\t\tReceived from ch2:", val)
			if val == 0 {
				fin++
			}
		default:
			// do something
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
		}
	}
	fmt.Println("Fin")
}
