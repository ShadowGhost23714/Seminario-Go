/* 3. Desarrolla un programa que implemente un sistema de planificación
(scheduler) utilizando 5 goroutines (el main y 4 más) . El programa
debe generar una serie de números enteros aleatorios, cada uno con
una prioridad aleatoria entre 0 y 3 (donde 0 es la prioridad más alta y 3
la más baja).

El scheduler debe seguir las siguientes reglas:

a) El scheduler debe procesar los datos en orden de prioridad,
comenzando por los de prioridad 0, luego 1, 2 y 3.

b) Mientras haya datos de prioridad 0, el scheduler debe
procesarlos exclusivamente.

c) Si no hay datos de prioridad 0 y hay goroutines disponibles, el
scheduler puede asignarles datos de menor prioridad para su
procesamiento.

d) Una vez que no haya datos de prioridad 0, el scheduler debe
pasar a procesar los datos de prioridad 1, y así sucesivamente,
utilizando las goroutines disponibles de forma dinámica.

e) El programa principal debe generar los datos numéricos
aleatorios con sus respectivas prioridades aleatorias y distribuir el
trabajo a las goroutines disponibles manteniendo el orden en el
que llegan los datos por prioridad.

f) Con los datos se debe:

i) Prioridad 0: Sumar los dígitos del número y almacenar el
resultado en un archivo llamado "prioridad0.txt" en formato
de par ordenado (0, resultado).

ii) Prioridad 1: Invertir los dígitos del número y almacenar el
resultado en un archivo llamado "prioridad1.txt" en formato
de par ordenado (1, resultado).

iii) Prioridad 2: Multiplicar el número por un valor constante
(por ejemplo, 10) e imprimir el resultado en la consola.

iv) Prioridad 3: Acumular los números y mostrar el valor
acumulado en la consola cada vez que se procesa un dato.

Tip: Puedes utilizar la librería math/rand para generar números
aleatorios.*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	productoConstante = 10
	df                = 25
)

type tarea struct {
	Numero    int
	Prioridad int
}

var (
	totalPrioridad3 int
	mu              sync.Mutex
	mu2             sync.Mutex
	mu3             sync.Mutex
	contador1       int
	contador2       int
	wg              sync.WaitGroup
	p0              chan tarea
	p1              chan tarea
	p2              chan tarea
	p3              chan tarea
)

func GenerarTarea() tarea {
	return tarea{Numero: rand.Intn(100), Prioridad: rand.Intn(4)}
}

func Multiplicar(numero int) int {
	return numero * productoConstante
}

func sumarDigitos(numero int, archivo0 *os.File) {
	suma := 0
	for numero > 0 {
		digito := numero % 10
		suma += digito
		numero /= 10
	}
	mu2.Lock()
	contador2++
	fmt.Fprintf(archivo0, "(%d), (0 , %d)\n", contador2, suma)
	mu2.Unlock()
}

func invertirDigitos(numero int, archivo1 *os.File) {
	var resultado string
	for numero > 0 {
		digito := numero % 10
		resultado += fmt.Sprint(digito)
		numero /= 10
	}
	mu.Lock()
	contador1++
	fmt.Fprintf(archivo1, "(%d) , (1 , %s)\n", contador1, resultado)
	mu.Unlock()
}

func worker(archivo1 *os.File, archivo0 *os.File) {
	defer wg.Done()
	c0 := true
	c1 := true
	c2 := true
	c3 := true
	for c0 || c1 || c2 || c3 {
		select {
		case tarea := <-p0:
			if len(p0) < 0 {
				fmt.Println("Resolviendo tarea de prioridad 0")
				sumarDigitos(tarea.Numero, archivo0)
				continue
			} else {
				c0 = false
			}
		default:
			select {
			case tarea := <-p1:
				if len(p1) < 0 {
					fmt.Println("Resolviendo tarea de prioridad 1")
					invertirDigitos(tarea.Numero, archivo1)
					continue
				} else {
					c1 = false
				}
			default:
				select {
				case tarea := <-p2:
					if len(p2) < 0 {
						fmt.Println("Resolviendo tarea de prioridad 2")
						fmt.Print(Multiplicar(tarea.Numero))
						continue
					} else {
						c2 = false
					}
				default:
					select {
					case tarea := <-p3:
						if len(p3) < 0 {
							mu3.Lock()
							fmt.Println("Resolviendo tarea de prioridad 3")
							totalPrioridad3 += tarea.Numero
							fmt.Println("total hasta ahora: ", totalPrioridad3)
							mu3.Unlock()
							continue
						} else {
							c3 = false
						}
					default:
						time.Sleep(1 * time.Second)
					}
				}
			}
		}
	}
	fmt.Println("HOLA")
}

func Distribuir(tarea tarea) {
	switch tarea.Prioridad {
	case 0:
		p0 <- tarea
	case 1:
		p1 <- tarea
	case 2:
		p2 <- tarea
	case 3:
		p3 <- tarea
	}
}

func scheduler() {
	totalPrioridad3 = 0
	contador1 = 0
	contador2 = 0

	archivo0, _ := os.Create("prioridad0.txt")
	archivo1, _ := os.Create("prioridad1.txt")
	defer archivo0.Close()
	defer archivo1.Close()

	p0 = make(chan tarea)
	p1 = make(chan tarea)
	p2 = make(chan tarea)
	p3 = make(chan tarea)

	wg.Add(4)
	for range 4 {
		go worker(archivo1, archivo0)
	}

	for range df {
		tarea := GenerarTarea()
		Distribuir(tarea)
	}

	close(p0)
	close(p1)
	close(p2)
	close(p3)
	wg.Wait()

	fmt.Println("hola")
}

func main() {
	scheduler()
}
