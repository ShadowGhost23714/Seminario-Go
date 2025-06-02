/*
2. Definir e implementar las operaciones de lista enlazada de la práctica
anterior usando el siguiente tipo de datos, de una lista genérica:

	type List[T any] struct {
		head, tail *element[T]
	}

	type element[T any] struct {
		next *element[T]
		val T
	}

Objetivo: tipos genéricos
*/

package main

import "fmt"

type Node struct {
	elem any // interface{}
	next *Node
}

type List *Node

func New() List {
	return (List(nil))
}

func IsEmpty(self List) bool {
	return (self == nil)
}

func Len(self List) int {
	if IsEmpty(self) {
		return 0
	} else {
		return 1 + Len(Next(self))
	}
}

func FrontElement(self List) any {
	return self.elem
}

func FrontElementTest(self List) (any, bool) {
	if !IsEmpty(self) {
		return self.elem, true
	} else {
		return 0, false
	}
}

func Next(self List) List {
	return self.next
}

func NextTest(self List) (List, bool) {
	if !IsEmpty(self) {
		return self.next, true
	} else {
		return nil, false
	}
}

func PushFront(self *List, elem any) {
	aux := new(Node)
	aux.elem = elem
	aux.next = *self
	*self = aux
}

func PushBack(self *List, elem any) {
	if IsEmpty(*self) {
		aux := new(Node)
		aux.elem = elem
		aux.next = nil
		*self = aux
	} else {
		PushBack((*List)(&((*self).next)), elem)
	}
}

func ToString(l List) string {
	if IsEmpty(l) {
		return ""
	} else {
		s := "(" + fmt.Sprint(l.elem) + ")"
		return s + " -> " + ToString(l.next)
	}
}

//invalid receiver type List (pointer or interface type)
//func (l List) String() string {
func (n Node) String() string {
	return ToString(&n)
}

func Remove(self *List) any {
	elem := (*self).elem
	//aux  := (*self)
	*self = (List)(((*self).next))
	return elem
}

func main() {
	l := New()
	fmt.Println("Lista vacia:", l)
	fmt.Println("Está vacia?", IsEmpty(l))
	fmt.Println("Long:", Len(l))

	PushFront(&l, 1)
	PushFront(&l, 2)
	PushFront(&l, 3)
	fmt.Println("Lista despues de PushFront:", l)
	fmt.Println("Esta vacia?", IsEmpty(l))
	fmt.Println("Long:", Len(l))

	elem, ok := FrontElementTest(l)
	if ok {
		fmt.Println("FrontElement:", elem)
	} else {
		fmt.Println("no hay elemento frontal")
	}

	nextList, ok := NextTest(l)
	if ok {
		fmt.Println("Siguiente lista:", nextList)
	} else {
		fmt.Println("no hay siguiente")
	}

	PushBack(&l, 4)
	PushBack(&l, 5)
	fmt.Println("PushBack:", l)

	removedElem := Remove(&l)
	fmt.Println("Elemento eliminado:", removedElem)
	fmt.Println("Lista después de sacar:", l)

	fmt.Println("Representación de la lista:", ToString(l))
}
