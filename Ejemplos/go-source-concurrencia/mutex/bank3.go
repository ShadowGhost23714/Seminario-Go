package main

import (
	"fmt"
	bs "mutex/bankSem"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			bs.Deposit(100)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(bs.Balance())
}
