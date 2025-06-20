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
)

const (
	productoConstante = 10
	df                = 100000
)

type tarea struct {
	Numero    int
	Prioridad int
}

var (
	totalPrioridad3 int
	mu              sync.Mutex
	contador1       int
	contador2       int
	wg              sync.WaitGroup
)

func GenerarTarea() tarea {
	return tarea{Numero: rand.Intn(1000), Prioridad: rand.Intn(4)}
}

func cargarSliceTareas(sliceTareas []tarea) []tarea {
	for range df {
		tarea := GenerarTarea()
		sliceTareas = append(sliceTareas, tarea)
	}
	return sliceTareas
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
	contador2++
	fmt.Fprintf(archivo0, "(%d), (0 , %d)\n", contador2, suma)
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
	fmt.Fprintf(archivo1, "(%d) , (1 , %d)\n", contador1, resultado)
	mu.Unlock()
}

func resolver(tarea tarea, archivo1 *os.File) {
	switch tarea.Prioridad {
	case 1:
		fmt.Println("Resolviendo tarea de prioridad 1")
		invertirDigitos(tarea.Numero, archivo1)
	case 2:
		fmt.Println("Resolviendo tarea de prioridad 2")
		fmt.Print(Multiplicar(tarea.Numero))
	case 3:
		fmt.Println("Resolviendo tarea de prioridad 3")
		totalPrioridad3 += tarea.Numero
		fmt.Println("total hasta ahora: ", totalPrioridad3)
	}

}

func worker(tc chan tarea, archivo1 *os.File) {
	tareaAuxiliar := tarea{Prioridad: 0}
	for tareaAuxiliar.Prioridad != -1 {
		tareaAuxiliar = <-tc
		resolver(tareaAuxiliar, archivo1)
	}
	wg.Done()
}

func scheduler(sliceTareas []tarea) {
	totalPrioridad3 = 0
	contador1 = 0
	contador2 = 0

	archivo0, _ := os.Create("prioridad0.txt")
	archivo1, _ := os.Create("prioridad1.txt")

	defer archivo1.Close()
	defer archivo0.Close()

	total1 := 0
	total2 := 0
	total3 := 0
	total4 := 0
	total5 := 0
	total0 := 0

	w1 := make(chan tarea)
	w2 := make(chan tarea)
	w3 := make(chan tarea)
	w4 := make(chan tarea)

	wg.Add(4)
	go worker(w1, archivo1)
	go worker(w2, archivo1)
	go worker(w3, archivo1)
	go worker(w4, archivo1)

	for i := range df {
		if sliceTareas[i].Prioridad == 0 {
			fmt.Println("Resolviendo tarea de prioridad 0")
			total0++
			sumarDigitos(sliceTareas[i].Numero, archivo0)
		} else {
			select {
			case w1 <- sliceTareas[i]:
				fmt.Println("consumido por worker 1")
				total1++
			case w2 <- sliceTareas[i]:
				fmt.Println("consumido por worker 2")
				total2++
			case w3 <- sliceTareas[i]:
				fmt.Println("consumido por worker 3")
				total3++
			case w4 <- sliceTareas[i]:
				fmt.Println("consumido por worker 4")
				total4++
			default:
				total5++
				fmt.Println("consumido por scheduler")
				resolver(sliceTareas[i], archivo1)
			}
		}
	}

	w1 <- tarea{Prioridad: -1}
	w2 <- tarea{Prioridad: -1}
	w3 <- tarea{Prioridad: -1}
	w4 <- tarea{Prioridad: -1}

	fmt.Print(
		"Total de tareas procesadas por worker (1): ", total1, " , porcentaje ", (total1*100)/df, "%\n",
		"Total de tareas procesadas por worker (2): ", total2, " , porcentaje ", (total2*100)/df, "%\n",
		"Total de tareas procesadas por worker (3): ", total3, " , porcentaje ", (total3*100)/df, "%\n",
		"Total de tareas procesadas por worker (4): ", total4, " , porcentaje ", (total4*100)/df, "%\n",
		"Total de tareas procesadas por scheduler que no son (0):  ", total5, " , porcentaje ", (total5*100)/df, "%\n",
		"Total de tareas procesadas por scheduler que son (0):     ", total0, " , porcentaje ", (total0*100)/df, "%\n")

	if total1+total2+total3+total4+total5+total0 != df {
		fmt.Println("No se procesaron todas las tares")
	}

	wg.Done()
}

func ordenarPorPrioridad(sliceTareas []tarea) []tarea {
	for i := range df {
		for j := i + 1; j < df; j++ {
			if sliceTareas[i].Prioridad > sliceTareas[j].Prioridad {
				sliceTareas[i], sliceTareas[j] = sliceTareas[j], sliceTareas[i]
			}
		}
	}
	return sliceTareas
}

func main() {
	sliceTareas := make([]tarea, 0)
	sliceTareas = cargarSliceTareas(sliceTareas)
	sliceTareas = ordenarPorPrioridad(sliceTareas)
	wg.Add(1)
	go scheduler(sliceTareas)
	wg.Wait()
}
