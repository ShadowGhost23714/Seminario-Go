package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		var val int
		for val != 100 {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
			val := <-ch1
			fmt.Println("Received from ch1:", val)
		}
		wg.Done()
	}()

	go func() {
		var val int
		for val != 100 {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
			val := <-ch2
			fmt.Println("\t\t\t\t\t\tReceived from ch2:", val)
		}
		wg.Done()
	}()

	for i := 0; i < 20; i++ {
		select {
		case ch1 <- i:
		case ch2 <- i:
		default:
			// do something
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		}
	}
	ch1 <- 100
	ch2 <- 100
	wg.Wait()
}
