/*
3. Ejecute el siguiente programa:

package main

import "fmt"

func main() {
	fmt.Println("Inicia Goroutine del main")
	go hello()
	fmt.Println("Termina Goroutine del main")
}

func hello() {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
	}
	fmt.Println("Termina Goroutine de hello")
}

a) ¿Cuántas veces se imprime Hello world?
b) ¿Cuántas Goroutines tiene el programa?
c) ¿Cómo cambiaría el programa (con la misma cantidad de
Goroutines) para que imprima 3 veces Hello world?
i) Hágalo usando time.Sleep
ii) Hágalo usando Channel Synchronization
*/

package main

import (
	"fmt"
)

func hello(done chan bool) {

	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
	}
	fmt.Println("Termina Goroutine de hello")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	fmt.Println("Inicia Goroutine del main")
	go hello(done)
	<-done
	fmt.Println("Termina Goroutine del main")
}
