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

func worker(id int, file0, file1 *os.File, wg *sync.WaitGroup) {
	c0 := true
	c1 := true
	c2 := true
	c3 := true
	for c0 || c1 || c2 || c3 {
		if len(prio0Chan) > 0 {
			tarea := <-prio0Chan
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
			c0 = false
			for {
				muTrabajos.Lock()
				if trabajosEnCero == 0 {
					muTrabajos.Unlock()
					break
				}
				muTrabajos.Unlock()
				time.Sleep(200 * time.Millisecond) // OPCIONAL : para no sobrecargar la CPU, se espera un tiempo antes de verificar si no hay tareas en 0
			}
			if len(prio1Chan) > 0 {
				tarea := <-prio1Chan
				str := invertirDigitos(tarea.num)
				fmt.Println("prioridad 1, worker id, valor de la tarea, resultado: ", id, " ", tarea.num, " ", str, " ")
				mut1.Lock()
				file1.WriteString(fmt.Sprintf("(1, %s)\n", str))
				mut1.Unlock()
			} else {
				c1 = false
				if len(prio2Chan) > 0 {
					tarea := <-prio2Chan
					res := tarea.num * 10
					fmt.Println("prioridad 2, worker id, valor de la tarea, resultado: ", id, " ", tarea.num, " ", res)
				} else {
					c2 = false
					if len(prio3Chan) > 0 {
						tarea := <-prio3Chan
						muAcum.Lock()
						acumulado += tarea.num
						fmt.Println("prioridad 3, worker id, valor de la tarea, resultado: ", id, " ", tarea.num, " ", acumulado)
						muAcum.Unlock()
					} else {
						c3 = false
						time.Sleep(1 * time.Second)

					}

				}

			}

		}

	}
	fmt.Println("HOLA1")
	wg.Done()
}

func scheduler(done chan int) {
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

	done := make(chan int) // los canales podrían ser globales

	var wg sync.WaitGroup

	for i := range 4 {
		wg.Add(1)
		go worker(i, file0, file1, &wg)
	}

	go scheduler(done)

	time.Sleep(cuantoTiempo * time.Second)

	close(done)

	wg1.Wait()

	close(prio0Chan)
	close(prio1Chan)
	close(prio2Chan)
	close(prio3Chan)

	wg.Wait()

	fmt.Println("termine")

	/*

		Este programa trata las tareas de la siguiente manera:

		genera una tarea aleatoria, la envia, si es prioridad 0 se le da trato exclusivo (solo se pueden procesar a la vez otras tareas de prioridad 0 mientras se este procesando tareas de prioridad 0)
		si la tarea no es de prioridad 0 y no hay tareas de prioridad 0 siendo tratadas, se aprovecha al maximo la concurrencia y se empiezan a procesar tareas hasta que se genere un 0,
		un worker la trate y se vuelva a producir una interrupcion en el flujo de procesamiento de tareas

		al ser generadas de forma instantanea y aleatoria, no es posible tratar las tareas de prioridad 1 2 y 3 con exclusividad una por encima de la otra



	*/

}
