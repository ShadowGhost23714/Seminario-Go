/*
6. Realice un programa que utilice select para recibir valores desde tres
canales diferentes. Cada canal debe recibir una secuencia de números
enteros. El programa debe recibir un valor de cada canal y mostrar el
valor recibido. Debes usar select para determinar cuál canal tiene un
valor disponible y recibir ese valor. El programa debe continuar hasta
haber recibido todos los valores enviados a cada canal. Al final debe
mostrar el total de valores recibidos desde cada canal.

Objetivo: select
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const lim = 20

var wg sync.WaitGroup

func recibir(c1 chan int, c2 chan int, c3 chan int) {
	defer wg.Done()
	cant1 := 0
	cant2 := 0
	cant3 := 0
	for i := 0; i < lim; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received canal 1 ", msg1)
			cant1++
		case msg2 := <-c2:
			fmt.Println("received canal 2 ", msg2)
			cant2++
		case msg3 := <-c3:
			fmt.Println("received canal 3 ", msg3)
			cant3++
		}
	}
	fmt.Println("recibidos en canal 1 =", cant1)
	fmt.Println("recibidos en canal 2 =", cant2)
	fmt.Println("recibidos en canal 3 =", cant3)
}

func main() {
	wg.Add(1)
	canal1 := make(chan int)
	canal2 := make(chan int)
	canal3 := make(chan int)
	go recibir(canal1, canal2, canal3)
	for i := 0; i < lim; i++ {
		sel := rand.Intn(3)
		sel++
		num := rand.Intn(100)
		switch {
		case sel == 1:
			canal1 <- num
		case sel == 2:
			canal2 <- num
		case sel == 3:
			canal3 <- num
		}

	}
	//close(canal1)
	//close(canal2)
	//close(canal3)
	wg.Wait()
}
