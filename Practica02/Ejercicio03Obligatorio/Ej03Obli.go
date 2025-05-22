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
element = 9 position = 6

el resultado sería:

[3(5), 1(1), 9(1) 1(6), 23(6), 3(8), 7(1), 5(3)]

Nota: no se permite realizar el Insert convirtiendo el OptimunSlice a un slice, insertar
y luego volver a convertirlo.
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

func FrontElement(o OptimumSlice) (int, bool) {
	if !IsEmpty(o) {
		return o[0].num, true
	}
	return -1, false
}

func LastElement(o OptimumSlice) (int, bool) {
	if !IsEmpty(o) {
		return o[len(o)-1].num, true
	}
	return -1, false
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

func Insert(o *OptimumSlice, element int, position int) int { // devulve un int para verificar
	if position < 0 {
		return -1
	}
	switch {
	case IsEmpty(*o):
		*o = append(*o, racha{element, 1})
	case position == 0:
		if (*o)[0].num == element {
			(*o)[0].cant++
			return 1

		} else {
			var copia OptimumSlice
			copia = append(copia, racha{element, 1})
			*o = append(copia, *o...)
			return 1
		}
	case position == Len(*o):
		if (*o)[len(*o)-1].num == element {
			(*o)[len(*o)-1].cant++
			return 1
		} else {
			*o = append(*o, racha{element, 1})
			return 1
		}
	default:
		posSlice := 0
		posArray := 0
		numeros := 0
		encontre := false
		for i := 0; i < len(*o) && !encontre; i++ {
			for numeros < (*o)[i].cant {
				if posArray == position {
					encontre = true
					break
				}
				posArray++
				numeros++
			}
			if !encontre {
				posSlice++
				numeros = 0
			}
		}
		if (*o)[posSlice].num == element {
			(*o)[posSlice].cant++
			return 1
		} else {
			division := (*o)[posSlice]
			division.cant = numeros
			(*o)[posSlice].cant -= numeros
			*o = append((*o)[:posSlice], append([]racha{division}, (*o)[posSlice:]...)...)
			*o = append((*o)[:posSlice+1], append([]racha{{element, 1}}, (*o)[posSlice+1:]...)...)
			return 1
		}
	}
	return 0
}

func SliceArray(o OptimumSlice) []int {
	s := []int
}

func main() {
	s := []int{3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 1, 1, 23, 23, 23, 23, 23, 23, 3, 3, 3, 3, 3, 3, 3, 3, 7, 5, 5, 5}
	o := New(s)
	fmt.Println(o.ToString())
	val := Insert(&o, 9, 30)
	if val == 1 {
		fmt.Println("Se inserto un nro")
	}
	fmt.Println(o.ToString())
	fmt.Println("Cantidad de nros en total =", Len(o))
	ok := false
	num := 0

	num, ok = FrontElement(o)
	if ok {
		fmt.Println("Nro del frente =", num)
	} else {
		fmt.Println("el slice esta vacio")
	}
	num, ok = LastElement(o)
	if ok {
		fmt.Println("Nro de atras =", num)
	} else {
		fmt.Println("el slice esta vacio")
	}
}
