package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var philos = []string{"Mark", "\t\t\tRussell", "\t\t\t\t\tRocky", "\t\t\t\t\t\t\tHaris", "\t\t\t\t\t\t\t\t\tRoot"}

var forks = [5]sync.Mutex{}

var dining sync.WaitGroup

func philosopher(id int, forkL, forkR *sync.Mutex) {
	name := philos[id]
	fmt.Println(name, "seated")
	for i := 0; i < 50; i++ {
		fmt.Println(name, "thinking")
		time.Sleep(time.Duration(rand.Intn(100) * id))

		forkL.Lock()
		time.Sleep(time.Duration(rand.Intn(100))) // para aumentar las chances de ocurrencia de deadlock
		forkR.Lock()

		fmt.Println(name, "eating")
		time.Sleep(time.Duration(rand.Intn(100) * id))

		forkL.Unlock()
		forkR.Unlock()
	}
	fmt.Println(name, "left the table")
	dining.Done()
}

func main() {
	dining.Add(5)
	for i := range philos {
		go philosopher(i, &forks[i], &forks[(i+1)%5])
	}
	dining.Wait()
}
