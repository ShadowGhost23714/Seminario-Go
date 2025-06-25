/* 1. Realice un programa que acepte un número entero positivo N como
entrada desde la línea de comandos y calcule todos los números
primos menores o iguales a N.

a) Realice el programa con una única goroutine que muestre en
pantalla la lista de números primos encontrados.

b) Realice el programa utilizando más de una goroutine para dividir
el trabajo de comprobación de primos entre múltiples goroutines
en paralelo

i) Cada goroutine debe recibir un rango de números a
comprobar y devolver una lista de los números primos
encontrados

ii) El programa principal debe recibir el número N y dividir el
trabajo en goroutines, asignando a cada una un rango de
números a comprobar

iii) Una vez que todas las goroutines hayan finalizado, el
programa principal debe recopilar los resultados y mostrar
en pantalla la lista de números primos encontrados.

c) Realice la ejecución con N igual a 1.000, 100.000 y 1.000.000
tanto del programa a) como del b). Para cada caso calcule el
speed-up con la siguiente fórmula:

S(p) = T(1) / T(p)

donde, S(p) es el speed-up, T(1) es el tiempo que tarda la
ejecución con una única goroutine y T(p) es el tiempo de
ejecución con p goroutines. */

package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	wg    sync.WaitGroup
	canal chan int
)

/*
func numerosPrimos(n int) { // Punto A
	defer wg.Done()
	if n > 2 {
		fmt.Println("Numeros primos antes o iguales: ")
		for i := 2; i <= n; i++ {
			esPrimo := true
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					esPrimo = false
				}
			}
			if esPrimo {
				fmt.Println(i)"
			}
		}
	}
}
*/

func numerosPrimos(n int) { // Punto B
	if n > 2 {
		fmt.Print("Numeros primos antes o iguales: ")
		for i := 2; i <= n; i++ {
			wg.Add(1)
			go numerosPrimosConcurrente(i)
		}
		wg.Wait()
		close(canal)
		for num := range canal {
			fmt.Println(num)
		}
	}
}

func numerosPrimosConcurrente(n int) {
	defer wg.Done()
	esPrimo := true
	for j := 2; j*j <= n; j++ {
		if n%j == 0 {
			esPrimo = false
			break
		}
	}
	if esPrimo {
		canal <- n
	}
}

func main() {
	// Verificar que se paso al menos un argumento (ademas del nombre del programa)
	if len(os.Args) < 2 {
		fmt.Println("Debe ingresar un número como argumento.")
		return
	}
	// Tomar el primer argumento
	arg := os.Args[1]
	// Convertirlo a entero
	n, err := strconv.Atoi(arg)
	if err != nil || n <= 0 {
		fmt.Println("Debe ingresar un número entero positivo.")
		return
	}
	// Usar el numero
	fmt.Printf("El número ingresado es: %d\n", n)
	canal = make(chan int, n)
	numerosPrimos(n)
}
