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

	for c := 0; c < 3; c++ {
		wg.Add(2)
		go func(id int) {
			for {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
				val, ok := <-ch1
				if !ok {
					break
				}
				fmt.Println("Received from ch1", id, "val:", val)
			}
			wg.Done()
		}(10 + c)

		go func(id int) {
			for {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
				val, ok := <-ch2
				if !ok {
					break
				}
				fmt.Println("\t\t\t\t\t\tReceived from ch2:", id, "val:", val)
			}
			wg.Done()
		}(20 + c)
	}

	for i := 0; i < 20; i++ {
		select {
		case ch1 <- i:
		case ch2 <- i:
		default:
			// do something
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		}
	}
	close(ch1)
	close(ch2)
	wg.Wait()
}
