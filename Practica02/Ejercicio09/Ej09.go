/*
9. Usando memoria dinámica con punteros escribir una programa que implemente
y use una lista enlazada de enteros.

a. Definir el tipo, operaciones para agregar elementos adelante, atrás,
poder iterar, etc. Las operaciones pueden ser:

func New() List 							// crea una lista vacía
func IsEmpty(self List) bool 				// si la lista está vacía
func Len(self List) int 					// cantidad de elementos
func FrontElement(self List) int			// devuelve el primer elemento
func Next(self List) List 					// devuelve la sublista sin el primero
func ToString(self List) string 			// representación en texto
func PushFront(self List, elem int) 		// agregar al inicio
func PushBack(self List, elem int) 			// agregar al final
func Remove(self List) int 					// elimina y devuelve el primer elemento
func Iterate(self List, f func(int) int) 	// aplica una función a cada elemento

b. Generar el programa que utiliza las operaciones programadas.

c. Investigar y usar el paquete: ‘‘container/list’’. Ver las diferencias y
similitudes con su implementación. Pensar de qué forma se podría hacer
de tipos genéricos.

d. Ver de mejorar la interfaz de las funciones, por ejemplo usando métodos
y códigos de error.

Sub-objetivo: Uso de memoria dinámica, structs, funciones anónimas.
Estudiar biblioteca estándar ofrecida por el lenguaje. Pensar cómo
encapsular código, orientar al alumno a pensar en packages. Métodos
para mejorar la interfaz y ver la posibilidad de retornar más de un valor
con código de errores en los casos que sea necesario.
*/

package main

import (
	"container/list"
	"errors"
	"fmt"
)

type node struct {
	val  int
	next *node
}

type List struct {
	first, last *node
}

func New() List {
	return List{}
}

func IsEmpty(self List) bool {
	return self.first == nil
}

func Len(self List) int {
	cant := 0
	l := self.first
	for l != nil {
		cant++
		l = l.next
	}
	return cant
}

func FrontElement(self List) (int, error) {
	if self.first == nil {
		return -1, errors.New("La lista esta vacia")
	}
	return self.first.val, nil
}

func Next(self List) List {
	return List{self.first.next, self.last}
}

func ToString(self List) string {
	str := "{ "
	l := self.first
	for l != nil {
		str += fmt.Sprintf("%d - ", l.val)
		l = l.next
	}
	return str + "nil }"
}

func PushFront(self *List, elem int) {
	nodo := &node{val: elem, next: self.first}
	self.first = nodo
	if self.last == nil {
		self.last = nodo
	}
}

func PushBack(self *List, elem int) {
	nodo := &node{val: elem}
	if self.first == nil {
		self.first = nodo
		self.last = nodo
	} else {
		self.last.next = nodo
		self.last = nodo
	}
}

func Remove(self *List) (int, error) {
	if self.first == nil {
		return -1, errors.New("No se puede eliminar en una lista vacia")
	}
	x := self.first.val
	self.first = self.first.next
	return x, nil
}

func Iterate(self List, f func(int) int) {
	l := self.first
	for l != nil {
		f(l.val)
		l = l.next
	}
}

func (self *List) FrontElement() int {
	if self.first == nil {
		return -1
	}
	return self.first.val
}

func (self *List) Len() int {
	cant := 0
	l := self.first
	for l != nil {
		cant++
		l = l.next
	}
	return cant
}

func (self *List) IsEmpty() bool {
	return self.first == nil
}

func (self List) Next() List {
	if self.first == nil {
		return self
	}
	return List{self.first.next, self.last}
}

func (self List) ToString() string {
	str := "{ "
	l := self.first
	for l != nil {
		str += fmt.Sprintf("%d - ", l.val)
		l = l.next
	}
	return str + "nil }"
}

func (self *List) Remove() int {
	x := self.first.val
	self.first = self.first.next
	return x
}

func main() {
	lista := New()
	PushFront(&lista, 20)
	PushBack(&lista, 30)
	PushBack(&lista, 40)
	PushFront(&lista, 10)
	fmt.Println()
	fmt.Println("Lista 1 =", ToString(lista))
	fmt.Println("Esta vacia?", IsEmpty(lista))
	fmt.Println("Longitud =", Len(lista))
	val, err := FrontElement(lista)
	if err == nil {
		fmt.Println("Primer elemento =", val)
	} else {
		fmt.Println("Error:", err)
	}
	for i := 0; i < 5; i++ {
		val, err := Remove(&lista)
		if err == nil {
			fmt.Println("Se elimino =", val)
		} else {
			fmt.Println("Error:", err)
		}
	}

	lista2 := list.New()
	lista2.PushFront(10)
	lista2.PushBack(20)
	lista2.PushBack(30)
	lista2.PushFront(0)
	str := "{ "
	for nodo := lista2.Front(); nodo != nil; nodo = nodo.Next() {
		str += fmt.Sprintf("%d - ", nodo.Value)
	}
	str += "nil }"
	fmt.Println()
	fmt.Println("Lista 2 =", str)
	fmt.Println("Longitud =", lista2.Len())

	lista3 := New()
	PushFront(&lista3, 5)
	fmt.Println()
	fmt.Println("Lista 3 =", lista3.ToString())
	fmt.Println("Se elimino =", lista3.Remove())
	fmt.Println("Esta vacia?", lista3.IsEmpty())
	fmt.Println("Longitud =", lista3.Len())
	fmt.Println("Primer elemento =", lista3.FrontElement())
}
