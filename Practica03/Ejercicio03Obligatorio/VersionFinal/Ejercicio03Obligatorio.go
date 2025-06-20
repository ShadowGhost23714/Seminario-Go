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

	ok bool

	wg1 sync.WaitGroup
)

func worker(file0, file1 *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	ok1 := true
	finalRound := false

	for {

		if !ok && !finalRound {
			finalRound = true
		} else if !ok && finalRound {
			break
		}

		ok1 = true
		for ok1 {
			select {
			case tarea := <-prio0Chan:
				{
					muTrabajos.Lock()
					trabajosEnCero++
					muTrabajos.Unlock()

					res := sumarDigitos(tarea.num)
					time.Sleep(5 * time.Second)
					fmt.Println("prioridad 0, valor de la tarea, resultado: ", tarea.num, " ", res)
					mut0.Lock()
					file0.WriteString(fmt.Sprintf("(0, %d)\n", res))
					mut0.Unlock()

					muTrabajos.Lock()
					trabajosEnCero--
					muTrabajos.Unlock()
				}
			default:
				for {
					muTrabajos.Lock()
					if trabajosEnCero == 0 {
						muTrabajos.Unlock()
						break
					}
					muTrabajos.Unlock()
					time.Sleep(200 * time.Millisecond)
				}

				select {
				case tarea := <-prio1Chan:
					str := invertirDigitos(tarea.num)
					fmt.Println("prioridad 1, valor de la tarea, resultado: ", tarea.num, " ", str, " ")
					mut1.Lock()
					file1.WriteString(fmt.Sprintf("(1, %s)\n", str))
					mut1.Unlock()
				default:
					select {
					case tarea := <-prio2Chan:
						res := tarea.num * 10
						fmt.Println("prioridad 2, valor de la tarea, resultado: ", tarea.num, " ", res)
					default:
						select {
						case tarea := <-prio3Chan:
							muAcum.Lock()
							acumulado += tarea.num
							fmt.Println("prioridad 3, valor de la tarea, resultado: ", tarea.num, " ", acumulado)
							muAcum.Unlock()
						default:
							ok1 = false
						}
					}
				}
			}
		}
	}
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
	ok = true

	file0, _ := os.Create("prioridad0.txt")
	defer file0.Close()

	file1, _ := os.Create("prioridad1.txt")
	defer file1.Close()

	done := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(file0, file1, &wg)
	}

	go scheduler(done)

	time.Sleep(cuantoTiempo * time.Second)

	close(done)
	wg1.Wait()
	ok = false

	wg.Wait()

	/*

			La variable finalRound permite que los workers sigan procesando tareas después de que el scheduler se detiene (luego de 30 segundos).Para garantizar que todas las tareas generadas sean procesadas.
		Por qué no alcanza con ok1 = false?
		Aunque los workers ponen ok1 = false cuando todas las colas están vacías, puede suceder que:

		condiciones

		1 - El scheduler este por cerrar,  Justo antes de cerrarse, envíe una nueva tarea,

		2 - El worker ya haya pasado la verificación de colas vacías (ultimo default)

		En ese caso, sin finalRound, esa tarea recién enviada no se procesaría, finalRound es el ultimo ciclo de vueltas que se hace despues de que Scheduler cierre, para tratar el caso especial descripto

		Se puede evitar finalRound?
		Solo si el objetivo es que tanto generación como procesamiento ocurran solo durante los 30 segundos. En ese caso, se debería forzar que todo termine al mismo tiempo,
		incluso si quedan tareas sin procesar, se cierran los canales sin importar si aun tienen datos dentro.

	*/

}
