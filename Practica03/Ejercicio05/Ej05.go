/*
5. Realice un programa que tenga 2 productores y 2 consumidores. Cada
productor y consumidor debe ser una Goroutine. Cada productor
producirá 3 números y cada consumidor consumirá 3 números. Los
productores deben esperar un tiempo aleatorio entre 0 y 1 segundo
para producir un número aleatorio entre 0 y 100. Los consumidores
deben consumirlos inmediatamente e imprimirlos por pantalla indicando
cual es el consumidor que lo consumió.

Objetivo: WaitGroups
*/

package main

import (
	"math/rand"
	"sync"
	"time"
)

func productores(numeros chan int) {

	for i := 0; i < 3; i++ {
		num := rand.Intn(100)
		num2 := rand.Intn(100000000)

		time.Sleep(time.Duration(num2))
		numeros <- num
	}
}
func consumidor(numeros chan int, id int, group *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		num2 := rand.Intn(1000000000)

		time.Sleep(time.Duration(num2))
		println("recibi ", <-numeros, ", quien soy ", id)
	}
	group.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	numeros := make(chan int)
	go productores(numeros)
	go productores(numeros)
	go consumidor(numeros, 1, &wg)
	go consumidor(numeros, 2, &wg)
	wg.Wait()

}
