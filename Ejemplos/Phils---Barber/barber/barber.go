package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const n = 5

var (
	temp int = 0
	mu   sync.Mutex
)

func prnt(cantTabs int, txt1, txt2 string, sillas <-chan string) {
	mu.Lock()
	temp++
	fmt.Printf("%v%v%v %v %v\n" /* time.Now() */, temp, strings.Repeat("\t", cantTabs), txt1, txt2, len(sillas))
	mu.Unlock()
}

func cortar(nombre string, listo chan<- bool, sillas <-chan string) {
	prnt(0, "- Cortando a", nombre, sillas)
	time.Sleep(time.Duration(100 + rand.Intn(100)))
	listo <- true
	prnt(0, "- Fin del corte a", nombre, sillas)
}

func barbero(sillas <-chan string, listo chan<- bool, despertar <-chan string) {
	for {
		// Durmiendo
		prnt(0, "- Durmiendo ...", "", sillas)
		nombre := <-despertar
		prnt(0, "- Despertado por", nombre, sillas)
		cortar(nombre, listo, sillas)
		// Me fijo si hay algún cliente esperando
		// Si hay, lo atiendo. Si no, duermo
		for len(sillas) > 0 {
			nombre := <-sillas
			prnt(6, nombre, "pasa a silla de corte", sillas)
			cortar(nombre, listo, sillas)
		}
	}
}

func cliente(nombre string, sillas chan string, listo <-chan bool, despertar chan<- string) {
	prnt(6, "Llega", nombre, sillas)
	select {
	case despertar <- nombre:
		prnt(6, nombre, "despierta al barbero", sillas)
	default:
		select {
		case sillas <- nombre:
			prnt(6, nombre, "se sienta en sala de espera", sillas)
			<-listo
			prnt(6, nombre, "se va con el pelo corto", sillas)
		default:
			prnt(6, nombre, "se va porque no hay lugar en sala de espera", sillas)
		}
	}
}

func main() {
	sillas := make(chan string, n)
	despertar := make(chan string)
	listo := make(chan bool)

	clientes := []string{"Alberto", "Braian", "Carlos", "Dionisio", "Esteban", "Fernando", "Gonzalo", "Hector", "Ignacio", "Javier", "Kevin", "Lionel", "Marcelo", "Nestor", "Osvaldo", "Pablo", "Quentin", "Raul", "Sergio", "Tobias", "Ubaldo", "Venencio", "Walter", "Xavi", "Yael", "Zacarias"}

	var wg sync.WaitGroup

	go barbero(sillas, listo, despertar)

	wg.Add(len(clientes))
	for _, nombre := range clientes {
		go func(nom string) {
			time.Sleep(time.Duration(1000 + rand.Intn(1000)))
			cliente(nom, sillas, listo, despertar)
			wg.Done()
		}(nombre)
	}
	wg.Wait()
}
