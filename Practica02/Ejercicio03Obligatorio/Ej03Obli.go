/*
Implemente una serie de funciones para manejar slices de enteros que
estadísticamente tienen muchas rachas de números repetidos. Utilice una
estructura (que con el objetivo de ahorrar memoria) tenga en cada elemento el
número entero y la cantidad de ocurrencias. Implemente:

func New(s slice) OptimumSlice
func IsEmpty(o OptimumSlice) bool
func Len(o OptimumSlice) int
func FrontElement(o OptimumSlice) int
func LastElement(o OptimumSlice) int
func Insert(o OptimumSlice, element int, position int) int
func SliceArray(o OptimumSlice) []int

Por ejemplo, si se invoca Insert con o =

[3(5), 1(7), 23(6), 3(8), 7(1), 5(3)]

que sería la representación del arreglo:

{3,3,3,3,3,1,1,1,1,1,1,1,23,23,23,23,23,23,3,3,3,3,3,3,3,3,7,5,5,5}

y donde X[Y], X es el elemento e Y es la cantidad de ocurrencias consecutivas
element = 9
position = 6
*/

package main

import (
	"fmt"
	"strconv"
)

type racha struct {
	num  int
	cant int
}

type OptimumSlice []racha

func New(s []int) OptimumSlice {
	if len(s) == 0 {
		return nil
	}
	var o OptimumSlice
	cant := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cant++
		} else {
			o = append(o, racha{s[i-1], cant})
			cant = 1
		}
	}
	o = append(o, racha{s[len(s)-1], cant})
	return o
}

func IsEmpty(o OptimumSlice) bool {
	return len(o) == 0
}

func Len(o OptimumSlice) int {
	cant := 0
	for _, r := range o {
		cant += r.cant
	}
	return cant
}

func FrontElement(o OptimumSlice) int {
	if !IsEmpty(o) {
		return o[0].num
	}
	return -1
}

func LastElement(o OptimumSlice) int {
	if !IsEmpty(o) {
		return o[len(o)-1].num
	}
	return -1
}

func (o OptimumSlice) ToString() string {
	str := "["
	for i := range o {
		str += strconv.Itoa(o[i].num) + "(" + strconv.Itoa(o[i].cant) + ")"
		if i+1 != len(o) {
			str += ", "
		}
	}
	return str + "]"
}

func main() {
	s := []int{3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 1, 1, 23, 23, 23, 23, 23, 23, 3, 3, 3, 3, 3, 3, 3, 3, 7, 5, 5, 5}
	o := New(s)
	fmt.Println(o.ToString())
	fmt.Println("Cantidad de nros en total =", Len(o))
	fmt.Println("Nro del frente =", FrontElement(o))
	fmt.Println("Nro de atras =", LastElement(o))
}

/*
// primera solucion (incorrecta)
func New(s []int) OptimumSlice {
	if len(s) < 0 {
		return nil
	}
	var o OptimumSlice
	var cant int
	var ultimo int
	for i := 0; i < len(s); i++ {
		ultimo = s[i]
		cant = 0
		for i < len(s) {
			if ultimo == s[i] {
				cant++
			}
			i++
		}
		o = append(o, racha{ultimo, cant})
	}
	return o
}
*/
