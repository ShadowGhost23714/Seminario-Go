package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Producer(out chan<- int) {
	timeProducer := rand.Intn(250)
	totalProduce := 10 //rand.Intn(5) + 1
	fmt.Println("Begin Producer: ", timeProducer, totalProduce)
	for i := 0; i < totalProduce; i++ {
		time.Sleep(time.Millisecond * time.Duration(timeProducer))
		product := rand.Intn(1000)
		fmt.Printf("P: %v para enviar\n", product)
		out <- product
		fmt.Printf("P: %v enviado\n", product)
	}
	fmt.Println("End Producer")
}

func Consumer(in <-chan int) {
	timeConsumer := rand.Intn(1000)
	fmt.Println("\t\t\t\t\tBegin Consumer: ", timeConsumer)
	for i := range in {
		time.Sleep(time.Millisecond * time.Duration(timeConsumer))
		fmt.Printf("\t\t\t\t\tC: %v\n", i)
	}
	fmt.Println("\t\t\t\t\tEnd Consumer")
}

func main() {
	ch := make(chan int, 5)

	var wgC sync.WaitGroup
	wgC.Add(1)

	// crear Productor
	fmt.Println("Crear Productor")
	go func() {
		Producer(ch)
		close(ch)
	}()

	// crear Consumidor
	fmt.Println("\t\t\t\t\tCrear Consumidor")
	go func() {
		Consumer(ch)
		wgC.Done()
	}()

	wgC.Wait()
}
