/*
11. Implementar el tipo de datos Ingresante y la funcionalidad solicitada.
a. Definir el tipo de datos Ingresante del cual se tiene la siguiente
información:

● Apellido.
● Nombre.
● Ciudad de origen.
● Fecha de nacimiento (día, mes, año).
● Si presentó el título del colegio secundario
● Código de la carrera en la que se inscribe (APU, LI, LS).
● Definir la función que implementa la interfaz Stringer para usar con
fmt.Println.

func (i Ingresante) String() string

b. Definir las funciones de comparación entre ingresantes por edad y por
orden alfabético/lexicográfico.

c. Cargar varios datos en un slice de ingresantes y ordenar primero por
edad y luego por apellido y nombre. Investigar el package ‘‘sort’’.

Sub-objetivo: Uso de arreglos, slices, structs, funciones de comparación.
Interfaz Stringer. Pensar cómo encapsular código, orientar al alumno a
pensar en packages. Métodos para mejorar la interfaz.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	apu = iota + 1
	li
	ls
)

type data struct {
	dia  int
	mes  int
	anio int
}

type Ingresante struct {
	Apellido   string
	Nombre     string
	Ciudad     string
	Fecha      data
	Secundario bool
	CodCarrera int
}

func (f data) String() string {
	return strconv.Itoa(f.dia) + "/" + strconv.Itoa(f.mes) + "/" + strconv.Itoa(f.anio)
}

func (i Ingresante) String() string {
	secu := "no presentado"
	if i.Secundario {
		secu = "presentado"
	}
	cod := ""
	switch i.CodCarrera {
	case apu:
		cod = "APU"
	case li:
		cod = "LI"
	default:
		cod = "LS"
	}
	return i.Nombre + " " + i.Apellido + " | " + i.Ciudad + " | " + i.Fecha.String() + " | Secundario " + secu + " | " + cod
}

func mayorEdad(i1 Ingresante, i2 Ingresante) bool {
	if i1.Fecha.anio != i2.Fecha.anio {
		return i1.Fecha.anio < i2.Fecha.anio
	}
	if i1.Fecha.mes != i2.Fecha.mes {
		return i1.Fecha.mes < i2.Fecha.mes
	}
	return i1.Fecha.dia < i2.Fecha.dia
}

func ordenAlfa(i1 Ingresante, i2 Ingresante) bool {
	if i1.Apellido != i2.Apellido {
		return i1.Apellido < i2.Apellido
	}
	return i1.Nombre < i2.Nombre
}

func main() {
	alumno := &Ingresante{"Argento", "Pepe", "Buenos Aires", data{14, 2, 1955}, true, 3}
	alumno2 := &Ingresante{"Argento", "Moni", "Claypole", data{7, 11, 1974}, false, 1}
	fmt.Println()
	fmt.Println(alumno)
	fmt.Println(alumno2)

	fmt.Println()
	fmt.Println("¿Quién nacio antes?")
	if mayorEdad(*alumno, *alumno2) {
		fmt.Println(alumno)
	} else {
		fmt.Println(alumno2)
	}

	fmt.Println()
	fmt.Println("Alfabéticamente ¿Quién va primero?")
	if ordenAlfa(*alumno, *alumno2) {
		fmt.Println(alumno)
	} else {
		fmt.Println(alumno2)
	}

	alumno3 := &Ingresante{"Argento", "Coqui", "Buenos Aires", data{29, 1, 1981}, true, 3}
	alumno4 := &Ingresante{"Argento", "Paola", "Buenos Aires", data{18, 5, 1987}, true, 3}
	alumno5 := &Ingresante{"Fuseneco", "María Elena", "Buenos Aires", data{1, 12, 1974}, true, 3}
	alumno6 := &Ingresante{"Fuseneco", "Dardo", "Buenos Aires", data{18, 6, 1968}, true, 3}
	ingresantes := []Ingresante{*alumno, *alumno2, *alumno3, *alumno4, *alumno5, *alumno6}

	// Ordenar primero por edad y luego por orden alfabético
	sort.Slice(ingresantes, func(i, j int) bool { // devuelve true si slice[i] debe ir antes que slice[j]
		// Primero: comparar por edad
		if mayorEdad(ingresantes[i], ingresantes[j]) {
			return true
		}
		if mayorEdad(ingresantes[j], ingresantes[i]) {
			return false
		}
		// Si tienen la misma edad: comparar por apellido y nombre
		return ordenAlfa(ingresantes[i], ingresantes[j])
	})

	fmt.Println()
	fmt.Println("Ingresantes ordenados:")
	for _, ing := range ingresantes {
		fmt.Println(ing)
	}
}
