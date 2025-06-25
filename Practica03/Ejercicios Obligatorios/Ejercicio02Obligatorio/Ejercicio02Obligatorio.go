/* 2. Realice un programa que simule la atención de clientes en las cajas de
un supermercado. La atención de cada cliente por parte del cajero se
debe simular con un timer entre 0 y 1 segundo.

a) Realice el programa haciendo esperar a los clientes en un única
cola global y luego enviándolo a la caja para su atención cuando
esta se encuentre disponible

b) Realice el programa haciendo esperar a los clientes en colas
individuales por caja asignando la caja para su atención con un
algoritmo round-robin

c) Realice el programa haciendo esperar a los clientes en colas
individuales por caja asignando la caja para su atención a aquella
que tenga la cola más corta

d) Imprima los tiempos de ejecución de cada uno de los programas
implementados en a), b) y c) */

package main

import (
	"fmt"
	"sync"
	"time"
)

const clientes = 20

var (
	//atencionChan chan int
	cola1 chan int
	cola2 chan int
	//wg        sync.WaitGroup
	wgCliente sync.WaitGroup
	wgCaja    sync.WaitGroup
)

/*// ----- punto A -----
func cliente(id int) {
	defer wg.Done()
	atencionChan <- id
}

func caja() {
	defer wg.Done()
	for range clientes {
		time.Sleep(1 * time.Second)
		id := <-atencionChan
		fmt.Println("Se atendio el cliente ", id)
	}
}

func super() {
	clieChan = make(chan int, clientes/2)
	for i := range clientes {
		wg.Add(1)
		go cliente(i + 1)
	}
	wg.Add(1)
	go caja()
	wg.Wait()
}
*/

/*// ----- Punto B -----
func cliente(id int, turno int) {
	defer wgCliente.Done()
	if turno%2 == 0 {
		cola1 <- id
	} else {
		cola2 <- id
	}
}

func caja(id int, cola chan int) {
	defer wgCaja.Done()
	for idCliente := range cola {
		time.Sleep(1 * time.Second)
		fmt.Println("El cajero", id, " atendio al cliente ", idCliente)
	}
}

func super() {
	cola1 = make(chan int, 5)
	cola2 = make(chan int, 5)
	wgCaja.Add(2)
	go caja(1, cola1)
	go caja(2, cola2)

	for i := range clientes {
		wgCliente.Add(1)
		go cliente(i+1, i)
	}

	wgCliente.Wait()
	close(cola1)
	close(cola2)
	wgCaja.Wait()
}
*/

// ----- Punto C -----
func cliente(id int) {
	defer wgCliente.Done()
	if len(cola1) < len(cola2) {
		cola1 <- id
	} else {
		cola2 <- id
	}
}

func caja(id int, cola chan int) {
	defer wgCaja.Done()
	for idCliente := range cola {
		time.Sleep(1 * time.Second)
		fmt.Println("El cajero", id, " atendio al cliente ", idCliente)
	}
}

func super() {
	cola1 = make(chan int, clientes/2)
	cola2 = make(chan int, clientes/2)
	wgCaja.Add(2)
	go caja(1, cola1)
	go caja(2, cola2)

	for i := range clientes {
		wgCliente.Add(1)
		go cliente(i + 1)
	}

	wgCliente.Wait()
	close(cola1)
	close(cola2)
	wgCaja.Wait()
}

func main() {
	super()
}
