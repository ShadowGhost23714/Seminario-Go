package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	cuantoTiempo = 30 // cantidad de tiempo que queres que se ejecute el programa
)

type Tarea struct {
	num       int
	prioridad int
}

var (
	prio0Chan = make(chan Tarea, 10)
	prio1Chan = make(chan Tarea, 10)
	prio2Chan = make(chan Tarea, 10)
	prio3Chan = make(chan Tarea, 10)

	mut1 sync.Mutex
	mut0 sync.Mutex

	acumulado      int
	muAcum         sync.Mutex
	trabajosEnCero int
	muTrabajos     sync.Mutex

	wg1 sync.WaitGroup
)

func worker(id int, tareaCh chan Tarea, file0, file1 *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	for tarea := range tareaCh {
		if tarea.prioridad == 0 {
			muTrabajos.Lock()
			trabajosEnCero++
			muTrabajos.Unlock()

			res := sumarDigitos(tarea.num)
			time.Sleep(5 * time.Second) // OPCIONAL : simular trabajo para que se note el corte en el flujo de procesamiento de datos a la hora que hay un 0 presente
			fmt.Println("prioridad 0, worker id, valor de la tarea, resultado: ", id, " ", tarea.num, " ", res)
			mut0.Lock()

			file0.WriteString(fmt.Sprintf("(0, %d)\n", res))
			mut0.Unlock()

			muTrabajos.Lock()
			trabajosEnCero--
			muTrabajos.Unlock()
		} else {
			for {
				muTrabajos.Lock()
				if trabajosEnCero == 0 {
					muTrabajos.Unlock()
					break
				}
				muTrabajos.Unlock()
				time.Sleep(200 * time.Millisecond) // OPCIONAL : para no sobrecargar la CPU, se espera un tiempo antes de verificar si no hay tareas en 0
			}

			switch tarea.prioridad {
			case 1:
				str := invertirDigitos(tarea.num)
				fmt.Println("prioridad 1, worker id, valor de la tarea, resultado: ", id, " ", tarea.num, " ", str, " ")
				mut1.Lock()

				file1.WriteString(fmt.Sprintf("(1, %s)\n", str))
				mut1.Unlock()
			case 2:
				res := tarea.num * 10
				fmt.Println("prioridad 2, worker id, valor de la tarea, resultado: ", id, " ", tarea.num, " ", res)
			case 3:
				muAcum.Lock()
				acumulado += tarea.num
				fmt.Println("prioridad 3, worker id, valor de la tarea, resultado: ", id, " ", tarea.num, " ", acumulado)
				muAcum.Unlock()
			}

		}
	}
}

func scheduler(tareaCh chan Tarea, done chan int) {
	wg1.Add(1)
	for {
		select {
		case <-done:
			wg1.Done()

			return
		default:
			num := rand.Intn(900) + 100
			prio := rand.Intn(4)
			t := Tarea{num: num, prioridad: prio}

			switch prio {
			case 0:

				prio0Chan <- t
			case 1:

				prio1Chan <- t
			case 2:
				prio2Chan <- t
			case 3:
				prio3Chan <- t
			}

			select {
			case tarea := <-prio0Chan:
				tareaCh <- tarea
			default:
				select {
				case tarea := <-prio1Chan:
					tareaCh <- tarea
				case tarea := <-prio2Chan:
					tareaCh <- tarea
				case tarea := <-prio3Chan:
					tareaCh <- tarea
				default:
					time.Sleep(50 * time.Millisecond)
				}
			}

			time.Sleep(50 * time.Millisecond)
		}
	}
}

func sumarDigitos(n int) int {
	suma := 0
	for n > 0 {
		suma += n % 10
		n /= 10
	}
	return suma
}

func invertirDigitos(numero int) string {
	var resultado string
	for numero > 0 {
		digito := numero % 10
		resultado += fmt.Sprint(digito)
		numero /= 10
	}
	return resultado
}

func main() {

	file0, _ := os.Create("prioridad0.txt")
	defer file0.Close()

	file1, _ := os.Create("prioridad1.txt")
	defer file1.Close()

	tareaCh := make(chan Tarea) //
	done := make(chan int)      // los canales podrían ser globales

	var wg sync.WaitGroup

	for i := range 4 {
		wg.Add(1)
		go worker(i, tareaCh, file0, file1, &wg)
	}

	go scheduler(tareaCh, done)

	time.Sleep(cuantoTiempo * time.Second)

	close(done)
	wg1.Wait()
	close(tareaCh)

	wg.Wait()

	//fmt.Println("termine")

	/*

		Este programa trata las tareas de la siguiente manera:

		genera una tarea aleatoria, la envia, si es prioridad 0 se le da trato exclusivo (solo se pueden procesar a la vez otras tareas de prioridad 0 mientras se este procesando tareas de prioridad 0)
		si la tarea no es de prioridad 0 y no hay tareas de prioridad 0 siendo tratadas, se aprovecha al maximo la concurrencia y se empiezan a procesar tareas hasta que se genere un 0,
		un worker la trate y se vuelva a producir una interrupcion en el flujo de procesamiento de tareas

		al ser generadas de forma instantanea y aleatoria, no es posible tratar las tareas de prioridad 1 2 y 3 con exclusividad una por encima de la otra



	*/

}
