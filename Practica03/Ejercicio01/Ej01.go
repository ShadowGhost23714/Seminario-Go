/*
1. Realice un programa que haga dos usos con distintos tipos del
siguiente tipo genérico:

type Map[K comparable, V any] map[K]V

Objetivo: tipos genéricos
*/

package main

import "fmt"

// Definición del tipo genérico Map
type MapaG[K comparable, V any] map[K]V

func main() {

	m1 := make(MapaG[string, int])
	m1["a"] = 1
	m1["b"] = 2
	fmt.Println("Mapa string->int:", m1)

	type Persona struct {
		Nombre string
		Edad   int
	}
	m2 := make(MapaG[int, Persona])
	m2[1] = Persona{"Theo", 30}
	m2[2] = Persona{"Pepe", 25}
	fmt.Println(m2)

	m3 := make(MapaG[bool, []float64])
	m3[true] = []float64{1.1, 2.2, 3.3}
	m3[false] = []float64{-1.1, -2.2, -3.3}
	fmt.Println(m3)
}
