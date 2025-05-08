package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var philos = []string{"Mark", "\t\t\tRussell", "\t\t\t\t\tRocky", "\t\t\t\t\t\t\tHaris", "\t\t\t\t\t\t\t\t\tRoot"}

var forks = [5]sync.Mutex{}

type ints [5]int

var count = ints{}
var penalty = ints{}

var dining sync.WaitGroup

func philosopher(id int, forkL, forkR *sync.Mutex) {
	//name := philos[id]
	//fmt.Println(name, "seated")
	for i := 0; i < 100; i++ {
		//fmt.Println(name, "thinking")
		time.Sleep(time.Duration(500 * penalty[id]))

		forkL.Lock()
		forkR.Lock()

		//fmt.Println(name, "eating")
		time.Sleep(time.Duration(rand.Intn(100)))

		forkL.Unlock()
		forkR.Unlock()

		report(id)
	}
	//fmt.Println(name, "left the table")
	dining.Done()
}

func max(count ints, id int) int {
	result := 0
	for i, v := range count {
		if i != id && v > result {
			result = v
		}
	}
	return result
}

func report(id int) {
	// registro la cantidad de veces que come cada phil y calculo una penalidad
	var mu sync.RWMutex
	count[id]++
	mu.RLock()
	if float64(count[id]) > float64(max(count, id))*1.1 {
		penalty[id]++
	} else {
		if penalty[id] > 0 {
			penalty[id]--
		}
	}
	fmt.Println("Reporte:", id, count, penalty)
	mu.RUnlock()
}

func main() {
	dining.Add(5)
	for i := range philos {
		// Rompo la posibilidad de espera circular haciendo que algunos
		// tomen primero el tenedor de la derecha y luego el de la izquierda,
		// y otros a la inversa

		go philosopher(i, &forks[(i+i%2)%5], &forks[(i+1-i%2)%5])
		/* if i%2 == 0 {
			go philosopher(i, &forks[i], &forks[(i+1)%5])
		} else {
			go philosopher(i, &forks[(i+1)%5], &forks[i])
		} */
	}
	dining.Wait()
}
