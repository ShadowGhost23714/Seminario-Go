package main

import (
	"fmt"
	bm "mutex/bankMonitor"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			bm.Deposit(100)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(bm.Balance())
}
