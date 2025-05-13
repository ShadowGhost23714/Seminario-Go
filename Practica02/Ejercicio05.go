/*
5. Escribir un programa que defina el tipo vector de flotantes de tamaño fijo,
constante, con las operaciones:

func Initialize(v Vector ,f float64)

func Sum(v1 , v2 Vector) Vector

func SumInPlace(v1 , v2 Vector)

SumInPlace, a diferencia de la anterior, guarda el resultado de la suma en el
primer vector. Investigar formas que existen para encapsular y separar el código.

Sub-objetivo: Arreglos, funciones y parámetros por referencia.
*/

package main

import "fmt"

const df = 3

type Vector [df]float64

func Initialize(v *Vector) {
	var x float64
	for i := 0; i < df; i++ {
		fmt.Print("Ingrese un nro: ")
		fmt.Scanln(&x)
		v[i] = x
	}
}

func Sum(v1, v2 Vector) Vector {
	var result Vector
	for i := 0; i < df; i++ {
		result[i] = v1[i] + v2[i]
	}
	return result
}

func SumInPlace(v1 *Vector, v2 Vector) {
	for i := 0; i < df; i++ {
		v1[i] += v2[i]
	}
}

func (v Vector) String() string {
	return fmt.Sprintf("%v", [df]float64(v))

}

func main() {
	var v1, v2 Vector
	fmt.Println("Ingrese ", df, " nros para el vector 1")
	Initialize(&v1)
	fmt.Println("Ingrese ", df, " nros para el vector 2")
	Initialize(&v2)
	fmt.Println()
	fmt.Println("Vector 1 = ", v1)
	fmt.Println("Vector 2 =", v2)
	fmt.Println()
	fmt.Println("Suma de vectores 1 y 2 = ", Sum(v1, v2))
	SumInPlace(&v1, v2)
	fmt.Println("SumaInPlace = ", v1)
	fmt.Println()
	fmt.Println("Vector 1 (Nuevo) = ", v1)
	fmt.Println("Vector 2 =", v2)
}
