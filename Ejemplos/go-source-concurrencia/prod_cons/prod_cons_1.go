package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Producer(id int, out chan<- int) {
	timeProducer := rand.Intn(250)
	totalProduce := rand.Intn(5) + 1
	fmt.Println("Begin Producer: ", id, timeProducer, totalProduce)
	for i := 0; i < totalProduce; i++ {
		time.Sleep(time.Millisecond * time.Duration(timeProducer))
		product := rand.Intn(1000)
		fmt.Printf("P %v: %v para enviar\n", id, product)
		out <- product
		fmt.Printf("P %v: %v enviado\n", id, product)
	}
	fmt.Println("End Producer: ", id)
}

func Consumer(id int, in <-chan int) {
	timeConsumer := rand.Intn(1000)
	fmt.Println("\t\t\t\t\tBegin Consumer: ", id, timeConsumer)
	for i := range in {
		time.Sleep(time.Millisecond * time.Duration(timeConsumer))
		fmt.Printf("\t\t\t\t\tC %v: %v\n", id, i)
	}
	fmt.Println("\t\t\t\t\tEnd Consumer: ", id)
}

func main() {
	ch := make(chan int)
	cProd := 2
	cCons := 5

	var wgP, wgC sync.WaitGroup

	wgP.Add(cProd)
	wgC.Add(cCons)

	// crear Consumidores
	fmt.Println("\t\t\t\t\tCrear Consumidores")
	for c := 1; c <= cCons; c++ {
		go func(id int) {
			Consumer(id, ch)
			wgC.Done()
		}(c)
	}

	// crear Productores
	fmt.Println("Crear Productores")
	for p := 1; p <= cProd; p++ {
		go func(id int) {
			Producer(id, ch)
			wgP.Done()
		}(p)
	}

	wgP.Wait()
	close(ch)
	wgC.Wait()
}
