/*
6. Escribir un programa que implemente dos funciones sobre slices. La primera
recibe dos slices de enteros y retorna un tercer slice del tamaño del mínimo
entre los dos sumando elemento a elemento. La segunda recibe un slice de
enteros y calcula el promedio de sus elementos. Por ejemplo, las definiciones
de las funciones pueden ser las siguientes:

func Sum(a , b []int) []int

func Avg(a []int) int

a. Re-implemente la función Avg para que retorne un float.

Sub-objetivo: Slices, funciones y parámetros.
*/

package main

import "fmt"

const (
	df1 = 3
	df2 = 4
)

func Inicializar(s []int) {
	var x int
	for i := 0; i < len(s); i++ {
		fmt.Print("Ingrese un nro: ")
		fmt.Scanln(&x)
		s[i] = x
	}
}

func Suma(a, b []int) []int {
	var x int
	if len(a) < len(b) {
		x = len(a)
	} else {
		x = len(b)
	}
	s := make([]int, x)
	for i := 0; i < x; i++ {
		s[i] = a[i] + b[i]
	}
	return s
}

func Avg(a []int) int {
	suma := 0
	for i := 0; i < len(a); i++ {
		suma += a[i]
	}
	return suma / len(a)
}

func AvgF(a []int) float64 {
	sum := 0

	for i := 0; i < len(a); i++ {

		sum += a[i]

	}

	return float64(sum / len(a))
}

func main() {
	a := make([]int, 3)
	b := make([]int, 4)
	fmt.Println("Ingrese ", df1, " números para el slice A")
	Inicializar(a)
	fmt.Println("Ingrese ", df2, " números para el slice B")
	Inicializar(b)
	fmt.Println("Suma de slices A y B = ", Suma(a, b))
	fmt.Println("Promedio de slice A = ", Avg(a))
	fmt.Println("Promedio en float de slice A = ", AvgF(a))
}
