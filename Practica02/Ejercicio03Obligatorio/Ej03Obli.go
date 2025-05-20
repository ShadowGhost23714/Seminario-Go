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

type racha struct {
	num  int
	cant int
}

type OptimumSlice []racha

func New(s slice) OptimumSlice {
	if IsEmpty(s) {
		return nil
	}

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

func Insert(o OptimumSlice, element int, position int) int {

}

func main() {

}
