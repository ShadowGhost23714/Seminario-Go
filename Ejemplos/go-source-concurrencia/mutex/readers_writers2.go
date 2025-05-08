package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		mu      sync.RWMutex // guards balance
		balance int
	)
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(500)))
			mu.Lock()
			defer mu.Unlock()
			fmt.Println("Writer", id, "balance pre:", balance)
			balance++
			//time.Sleep(time.Duration(rand.Intn(100)))
			fmt.Println("Writer", id, "balance post:", balance)
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(500)))
			for j := 0; j < 10; j++ {
				mu.RLock()
				fmt.Println("R", id, "balance:", balance)
				time.Sleep(time.Duration(rand.Intn(100)))
				mu.RUnlock()
			}
		}(i)
	}
	wg.Wait()
}
