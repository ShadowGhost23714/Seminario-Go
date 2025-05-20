package main

import (
	"errors"
	"fmt"
)

// Pila con Lista enlazada

type node struct {
	val  int
	next *node
}

type Stack struct {
	top *node
}

func New() Stack {
	return Stack{}
}

func (s Stack) IsEmpty() bool {
	return s.top == nil
}

func (s Stack) Len() int {
	cant := 0
	nodo := s.top
	for nodo != nil {
		cant++
		nodo = nodo.next
	}
	return cant
}

func (s Stack) ToString() string {
	aux := "["
	nodo := s.top
	for nodo != nil {
		aux += fmt.Sprintf("%d", nodo.val)
		if nodo.next != nil {
			aux += ", "
		}
		nodo = nodo.next
	}
	return aux + "]"
}

func (s Stack) FrontElement() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("La pila esta vacia")
	}
	return s.top.val, nil
}

func (s *Stack) Push(element int) {
	nue := &node{val: element, next: s.top}
	s.top = nue
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	x := s.top.val
	s.top = s.top.next
	return x
}

func (s Stack) Iterate(f func(int) int) {
	nodo := s.top
	for nodo != nil {
		f(nodo.val)
		nodo = nodo.next
	}
}

func main() {
	pila := New()
	pila.Push(10)
	pila.Push(20)
	pila.Push(30)
	pila.Push(40)
	fmt.Println("Pila:", pila.ToString())
	fmt.Println("Pop:", pila.Pop())
	fmt.Println("Pop:", pila.Pop())
	fmt.Println("Pila:", pila.ToString())
	val, err := pila.FrontElement()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Tope de la pila:", val)
	}
}
