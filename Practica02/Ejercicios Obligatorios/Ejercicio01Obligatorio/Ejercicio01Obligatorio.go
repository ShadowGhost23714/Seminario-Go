/*
1. Usando la estructura de datos definida en el ejercicio 9 resolver el siguiente
problema. Se dispone de una lista con la información de los ingresantes a la
Facultad del año anterior. De cada ingresante se conoce: apellido, nombre,
ciudad de origen, fecha de nacimiento (día, mes, año), si presentó el título del
colegio secundario y el código de la carrera en la que se inscribe (APU, LI, LS).
Con esta información de los ingresantes se pide que recorra el listado una
vez para:

a) Informar el nombre y apellido de los ingresantes cuya ciudad origen es
“Bariloche”.

b) Calcular e informar el año en que más ingresantes nacieron.

c) Informar la carrera con la mayor cantidad de inscriptos.

d) Eliminar de la lista aquellos ingresantes que no presentaron el título.*/

package main

import (
	"fmt"
)

type data struct {
	dia  int
	mes  int
	anio int
}

type ingresante struct {
	ape        string
	nom        string
	ciudad     string
	fecha      data
	secundario bool
	cod        string
}

type node struct {
	val  ingresante
	next *node
}

type List struct {
	first, last *node
}

func New() List {
	return List{}
}

func (self *List) PushFront(elem ingresante) {
	nodo := &node{val: elem, next: self.first}
	self.first = nodo
	if self.last == nil {
		self.last = nodo
	}
}

func (self *List) PushBack(elem ingresante) {
	nodo := &node{val: elem}
	if self.first == nil {
		self.first = nodo
		self.last = nodo
	} else {
		self.last.next = nodo
		self.last = nodo
	}
}

func (self *List) FrontElement() ingresante {
	if self.first == nil {
		panic("Lista vacia")
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
		str += l.val.nom + " - "
		l = l.next
	}
	return str + "nil }"
}

func (self *List) Remove() ingresante {
	x := self.first.val
	self.first = self.first.next
	return x
}

func Iterate(self List, f func(ingresante) int) {
	l := self.first
	for l != nil {
		f(l.val)
		l = l.next
	}
}

func analizarLista(self *List) (int, string) {
	l := self.first
	anios := map[int]int{}
	carreras := map[string]int{"LS": 0, "LI": 0, "APU": 0}
	for l != nil {
		if l.val.ciudad == "Bariloche" {
			fmt.Println(l.val.nom, "nació en Bariloche")
		}
		anios[l.val.fecha.anio]++
		carreras[l.val.cod]++
		if l.val.secundario == true { // no funciona correctamente
			fmt.Println(l.val)
			self.Remove()
		}
		l = l.next
	}
	var maxAnio int
	var maxValor int
	for anio, valor := range anios {
		if valor > maxValor {
			maxValor = valor
			maxAnio = anio
		}
	}
	var maxCarrera string
	maxValor = 0
	for carrera, valor := range carreras {
		if valor > maxValor {
			maxValor = valor
			maxCarrera = carrera
		}
	}
	return maxAnio, maxCarrera
}

func main() {
	lista := New()
	lista.PushFront(ingresante{"Argento", "Pepe", "Buenos Aires", data{14, 2, 1955}, true, "LS"})
	lista.PushFront(ingresante{"Argento", "Moni", "Claypole", data{7, 11, 1974}, false, "LI"})
	lista.PushFront(ingresante{"Argento", "Coqui", "Bariloche", data{29, 1, 1981}, true, "LS"})
	lista.PushFront(ingresante{"Argento", "Paola", "Buenos Aires", data{18, 5, 1987}, false, "LS"})
	lista.PushFront(ingresante{"Fuseneco", "Maria Elena", "Buenos Aires", data{1, 12, 1974}, true, "LI"})
	lista.PushFront(ingresante{"Fuseneco", "Dardo", "Buenos Aires", data{18, 6, 1968}, false, "APU"})
	fmt.Println(lista.ToString())
	anio, carrera := analizarLista(&lista)
	fmt.Println("La carrera con mas inscriptos es ", carrera)
	fmt.Println("El año con mas inscriptos es ", anio)
	fmt.Println(lista.ToString())
}
