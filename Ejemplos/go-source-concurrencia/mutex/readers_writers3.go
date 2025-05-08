package main

import (
	"fmt"
	"math/rand"
	bs "mutex/bankSem3"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(65)
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)))
			bs.Deposit(100)
			wg.Done()
		}()
	}
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)))
			if bs.Withdraw(100) {
				fmt.Println("Se pudo retirar")
			} else {
				fmt.Println("No se pudo retirar")
			}
			wg.Done()
		}()
	}
	for i := 0; i < 50; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(10)))
			fmt.Println("Balance:", bs.Balance())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(bs.Balance())
}
