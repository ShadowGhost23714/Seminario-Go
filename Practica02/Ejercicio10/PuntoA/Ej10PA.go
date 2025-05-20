/*
10. Implementar usando slices un programa que use una pila de enteros.
a. Definir el tipo, operaciones para agregar elementos adelante, atrás,
poder iterar, etc. Las operaciones pueden ser:

func New() Stack
func IsEmpty(s Stack) bool
func Len(s Stack) int
func ToString(s Stack) string
func FrontElement(s Stack) int
func Push(s Stack , element int)
func Pop(s Stack) int
func Iterate(s Stack , f func (int) int)

b. Re-implementar usando la lista enlazada.
Sub-objetivo: Uso de slices, structs, funciones anónimas. Estudiar usar
código de terceros. Pensar cómo encapsular código usando packages.
Ver de cómo reusar código. Pensar en definir con métodos.
*/

package main

import "fmt"

// Pila con Slice

type Stack []int

func New() Stack {
	return Stack{}
}

func (s Stack) IsEmpty() bool {
	return len(s) > 0
}

func Len(s Stack) int {
	return len(s)
}

func (s Stack) ToString() string {
	str := "["
	for i := 0; i < len(s); i++ {
		str += fmt.Sprintf("%d", s[i])
		if i != len(s)-1 {
			str += ", "
		}
	}
	return str + "]"
}

func (s Stack) FrontElement() int {
	if !s.IsEmpty() {
		return s[len(s)]
	}
	return -1
}

func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

func (s *Stack) Pop() int {
	if !s.IsEmpty() {
		return -1
	}
	ultimo := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ultimo
}

func (s Stack) Iterate(f func(int) int) {
	for i := range s {
		f(s[i])
	}
}

func main() {
	pila := New()
	pila.Push(10)
	pila.Push(20)
	pila.Push(30)
	fmt.Println("Pila:", pila.ToString())
	fmt.Println("Pop:", pila.Pop())
	fmt.Println("Pop:", pila.Pop())
	fmt.Println("Pop:", pila.Pop())
	pila.Push(20)
	fmt.Println("Pila:", pila.ToString())
}
