/*
2. Implementar la función factorial de dos formas, una iterativa y otra
recursiva. Escribir un programa y compilar de forma que utilice una u
otra y la evalúe de 0 a 9.
*/

package main

import (
	"fmt"
)

func FactorialIteracion(x int) int {
	if x < 2 {
		return 1
	}
	res := x
	for i := 1; i < x; i++ {
		res = res * i
	}
	return res
}

func FactorialRecursivo(x int) int {
	if x < 2 {
		return 1
	}
	res := x * FactorialRecursivo(x-1)
	return res
}

func main() {
	var num int
	var opc int
	fmt.Print("Ingrese un número del 0 al 9 para saber su factorial: ")
	fmt.Scanln(&num)
	fmt.Println("1. Iteración")
	fmt.Println("2. Recursión")
	fmt.Print("Ingresar opción ")
	fmt.Scanln(&opc)
	fmt.Print("Resultado = ")
	if opc == 1 {
		fmt.Println(FactorialIteracion(num))
	} else {
		fmt.Println(FactorialRecursivo(num))
	}
}
