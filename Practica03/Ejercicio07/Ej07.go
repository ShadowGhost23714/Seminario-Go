/*
7. Realice un programa que envíe datos a dos canales desde dos
goroutines y estos sean recibidos en el programa principal por un
select. Los datos se deben recibir en uno de los canales por el lapso de
5 segundos y por el otro en el lapso de 10 segundos.

Objetivo: timeouts */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const lim = 10

func Enviar5(c5 chan int) {
	for i := 0; i < lim; i++ {
		sel := rand.Intn(2)
		sel++
		num := rand.Intn(100)
		c5 <- num
	}

}

func Enviar10(c10 chan int) {
	for i := 0; i < lim; i++ {
		sel := rand.Intn(2)
		sel++
		num := rand.Intn(100)
		c10 <- num
	}
}

func main() {
	canal5 := make(chan int)
	canal10 := make(chan int)
	go Enviar5(canal5)
	go Enviar10(canal10)

	for i := 0; i < lim*2; i++ {
		select {
		case <-time.After(5 * time.Second):
			msg1 := <-canal5
			fmt.Println("received canal 1 ", msg1)

		case <-time.After(10 * time.Second):
			msg2 := <-canal10
			fmt.Println("received canal 2 ", msg2)
		}
	}
}
