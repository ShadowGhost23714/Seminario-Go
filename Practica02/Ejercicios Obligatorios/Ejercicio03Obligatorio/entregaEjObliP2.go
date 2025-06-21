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

func FrontElement(o OptimumSlice) int {
	if !IsEmpty(o) {
		return o[0].num
	}
	panic("El slice esta vacío")
}

func LastElement(o OptimumSlice) int {
	if !IsEmpty(o) {
		return o[len(o)-1].num
	}
	panic("El slice esta vacío")
}

func (o OptimumSlice) ToString() string { // Este metodo se usa para imprimir el contenido del Optimum Slice
	str := "["
	for i := range o {
		str += strconv.Itoa(o[i].num) + "(" + strconv.Itoa(o[i].cant) + ")"
		if i+1 != len(o) {
			str += ", "
		}
	}
	return str + "]"
}

func Insert(o *OptimumSlice, element int, position int) int {
	if position < 0 || position > Len(*o) {
		panic("posicion invalida")
	}

	switch {
	case IsEmpty(*o):
		*o = append(*o, racha{element, 1})
		return 1
	case position == 0:
		if (*o)[0].num == element {
			(*o)[0].cant++
			return 1
		} else {
			*o = append([]racha{{element, 1}}, *o...)
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
			if position < posArray+(*o)[i].cant {
				posSlice = i
				numeros = position - posArray
				encontre = true
			} else {
				posArray += (*o)[i].cant
			}
		}

		if (*o)[posSlice].num == element {
			(*o)[posSlice].cant++
			return 1
		} else {
			if numeros == 0 && posSlice-1 != -1 && (*o)[posSlice-1].num == element {
				(*o)[posSlice-1].cant++
				return 1
			}
			if numeros+1 == (*o)[posSlice].cant && posSlice+1 < len(*o) && (*o)[posSlice+1].num == element {
				(*o)[posSlice+1].cant++
				return 1
			}
			if numeros == 0 {
				*o = append((*o)[:posSlice], append([]racha{{element, 1}}, (*o)[posSlice:]...)...)
			} else if numeros == (*o)[posSlice].cant {
				*o = append((*o)[:posSlice+1], append([]racha{{element, 1}}, (*o)[posSlice+1:]...)...)
			} else {
				izq := racha{(*o)[posSlice].num, numeros}
				der := racha{(*o)[posSlice].num, (*o)[posSlice].cant - numeros}
				*o = append((*o)[:posSlice], append([]racha{izq, {element, 1}, der}, (*o)[posSlice+1:]...)...)
			}
			return 1
		}
	}
}

func SliceArray(o OptimumSlice) []int {
	len := Len(o)
	result := make([]int, len)
	aux := 0
	for _, actual := range o {
		for i := 0; i < actual.cant; i++ {
			result[aux] = actual.num
			aux++
		}
	}
	return result
}

func main() {
	s := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	o := New(s)
	fmt.Println()
	fmt.Println("Slice original")
	fmt.Println(s)
	fmt.Println()
	fmt.Println("Optimum Slice")
	fmt.Println(o.ToString())
	fmt.Println()
	Insert(&o, 3, 5)
	fmt.Println("Nuevo Optimum Slice (Se inserto 9 en la posicion 6)")
	fmt.Println(o.ToString())
	fmt.Println()
	fmt.Println("Slice Optimum descomprimido")
	fmt.Println(SliceArray(o))
	fmt.Println()
	fmt.Println("El elemento del adelante es =", FrontElement(o))
	fmt.Println("El elemento de atras es =", LastElement(o))
}

/*
En este programa el Insert contempla el caso en el que el valor ingresado es igual al valor que se encuentra en la posicion anterior o posterior, este se ve a sumar.

[3(5), 1(7), 23(6), 3(8), 7(1), 5(3)]

[3 3 3 3 3 (1) 1 1 1 1 1 1 23 23 23 23 23 23 3 3 3 3 3 3 3 3 7 5 5 5]
			^
Si se ingresa un 3 donde esta indicado entre parentesis se considera que es correcto acoplarlo con los numeros correspondientes, con el objetivo de ahorrar memoria.

El resultado seria el siguiente:

[3(6), 1(7), 23(6), 3(8), 7(1), 5(3)]

Este programa esta tambien esta pensado para que pueda realizar un agregar atras y un agregar adelante.
*/
