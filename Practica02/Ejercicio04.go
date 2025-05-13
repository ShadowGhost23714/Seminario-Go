/*
4. Se debe leer tres sucesiones, de N (constante), números enteros cada una: x1 .
. . xn, y1 . . . yn, z1 . . . zn, almacenarlas en sus respectivos arreglos y calcular
luego:

R = (∑(i=1 to n) 1/xi − ∏(i=1 to n) zi³) × maxmin(yi)

Para la productoria, la sumatoria y el máximo-mínimo usar funciones. La función
maxmin retorna el máximo y el mínimo de la serie y luego ambos son
multiplicados por el resto de la ecuación.

Sub-objetivo: Resaltar el tipado fuerte de Go y usar casting. Operaciones con
Integer y Float. Arreglos. Funciones que retornan más de un valor.
*/

package main

import (
	"fmt"
	"math"
)

const (
	N = 3
)

func Sumatoria(v [N]int) float64 {
	suma := 0.0
	for i := 0; i < N; i++ {
		suma += Funcion(v[i])
	}
	fmt.Println(suma)
	return suma
}

func Funcion(x int) float64 {
	valor := float64(x)
	result := 0.0
	result += 1 / valor
	return result
}

func Productoria(v [N]int) float64 {
	mul := Funcion2(v[0])
	for i := 1; i < N; i++ {
		mul = mul * Funcion2(v[i])
	}
	fmt.Println(mul)
	return mul
}

func Funcion2(x int) float64 {
	valor := float64(x)
	return valor * valor * valor
}

func Maxmin(v [N]int) (float64, float64) {
	min := math.MaxInt
	max := math.MinInt

	for i := 0; i < N; i++ {
		if v[i] > max {
			max = v[i]
		}
		if v[i] < min {
			min = v[i]
		}
	}
	fmt.Println(min, max)
	return float64(min), float64(max)
}

func Ecuacion(v1 [N]int, v2 [N]int, v3 [N]int) float64 {
	resultado1 := Sumatoria(v1)
	resultado2 := Productoria(v2)
	min, max := Maxmin(v3)
	return (resultado1 - resultado2) * (min * max)
}

func main() {
	var v [N]int
	fmt.Println("Ingrese tres nro")
	for i := 0; i < N; i++ {
		fmt.Scan(&v[i])
	}
	var v2 [N]int
	fmt.Println("Ingrese tres nro")
	for i := 0; i < N; i++ {
		fmt.Scan(&v2[i])
	}
	var v3 [N]int
	fmt.Println("Ingrese tres nro")
	for i := 0; i < N; i++ {
		fmt.Scan(&v3[i])
	}
	fmt.Println("Resultado = ", Ecuacion(v, v2, v3))
}
