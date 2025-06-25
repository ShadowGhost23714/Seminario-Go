/* 3) Desarrolla un programa que implemente un sistema de planificación
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
aleatorios. */

package main

import (
	"fmt"
	"math/rand"
)

type tarea struct {
	Numero    int
	Prioridad int
}

func GenerarTarea() tarea {
	return tarea{Numero: rand.Intn(1000), Prioridad: rand.Intn(4)}
}

func cargarSliceTareas(sliceTareas []tarea) []tarea {
	for i := 0; i < 100000; i++ {
		tarea := GenerarTarea()
		sliceTareas = append(sliceTareas, tarea)
	}
	return sliceTareas
}

func ordenarPorPrioridad(sliceTareas []tarea) []tarea {
	for i := 0; i < len(sliceTareas); i++ {
		for j := i + 1; j < len(sliceTareas); j++ {
			if sliceTareas[i].Prioridad > sliceTareas[j].Prioridad {
				sliceTareas[i], sliceTareas[j] = sliceTareas[j], sliceTareas[i]
			}
		}
	}
	return sliceTareas
}

func resolver(tarea tarea) {
	switch tarea.Prioridad {
	case 1:
		fmt.Println("Resolviendo tarea de prioridad 1")
	case 2:
		fmt.Println("Resolviendo tarea de prioridad 2")
	case 3:
		fmt.Println("Resolviendo tarea de prioridad 3")
	}
}

func worker(tc chan tarea) {
	tareaAuxiliar := <-tc
	resolver(tareaAuxiliar)
}

func scheduler(sliceTareas []tarea) {

	total1 := 0
	total2 := 0
	total3 := 0
	total4 := 0
	total5 := 0
	total0 := 0

	w1 := make(chan tarea, 1)
	w2 := make(chan tarea, 1)
	w3 := make(chan tarea, 1)
	w4 := make(chan tarea, 1)

	for i := 0; i < len(sliceTareas); i++ {
		if sliceTareas[i].Prioridad == 0 {
			fmt.Println("Resolviendo tarea de prioridad 0")
			total0++
		} else {
			select {
			case w1 <- sliceTareas[i]:
				fmt.Println("consumido por worker 1")
				total1++
				go worker(w1)
			case w2 <- sliceTareas[i]:
				fmt.Println("consumido por worker 2")
				total2++
				go worker(w2)
			case w3 <- sliceTareas[i]:
				fmt.Println("consumido por worker 3")
				total3++
				go worker(w3)
			case w4 <- sliceTareas[i]:
				fmt.Println("consumido por worker 4")
				total4++
				go worker(w4)
			default:
				total5++
				fmt.Println("consumido por scheduler")
				resolver(sliceTareas[i])
			}
		}
	}

	fmt.Print("Total de tareas procesadas por worker 1: ", total1, "\n",
		"Total de tareas procesadas por worker 2: ", total2, "\n",
		"Total de tareas procesadas por worker 3: ", total3, "\n",
		"Total de tareas procesadas por worker 4: ", total4, "\n",
		"Total de tareas procesadas por scheduler: ", total5, "\n",
		"Total de tareas procesadas por scheduler que son 0: ", total0, "\n")
}

func main() {
	sliceTareas := make([]tarea, 0)
	sliceTareas = cargarSliceTareas(sliceTareas)
	sliceTareas = ordenarPorPrioridad(sliceTareas)
	scheduler(sliceTareas)
}
