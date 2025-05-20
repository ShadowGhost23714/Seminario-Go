/*
3. Declare el tipo de datos punto cardinal como un enumerativo. Realizar un
programa que lea un punto cardinal del cual viene el viento (N, S, E, O, NO,
SE, NE, SO) e imprima hacia cuál se dirige. Encapsule la funcionalidad en una
función.

a. Resolver usando un switch/case.
b. Resolver usando el orden en que fueron definidos, notar que el contrario
es hacia adelante o atrás por su índice par o impar.
c. Resolver con un Map que tiene como índice el tipo punto cardinal y cada
elemento es el punto cardinal contrario al índice.
d. ¿Cómo se declaran los tipos enumerativos definidos por el usuario en
otros lenguajes que conoce?

rta :
En Pascal se usa type con una lista de valores entre paréntesis:
type PuntoCardinal = (norte, sur, este, oeste);
En Java se usa la palabra clave enum:
public enum PuntoCardinal { NORTE, SUR, ESTE, OESTE; }

e. Definir la función que implementa la interfaz Stringer para usar con
fmt.Println sobre un punto Cardinal.
f. ¿Qué sucede con las funciones anteriores cuando reciben un valor fuera
de rango?

Sub-objetivo: Declarar tipos enumerativos definidos por el usuario. Usar
el operador iota. Uso de Maps. Realizar E/S de tipos enumerativos. Funciones
sobre tipos enumerativos.
*/

package main

import "fmt"

type PuntoCardinal int

var puntos = map[PuntoCardinal]string{
	norte:    "sur",
	sur:      "norte",
	este:     "oeste",
	oeste:    "este",
	noreste:  "suroeste",
	suroeste: "noreste",
	noroeste: "sureste",
	sureste:  "noroeste",
}

const (
	norte PuntoCardinal = iota
	sur
	este
	oeste
	noroeste
	sureste
	noreste
	suroeste
)

func ConseguirPunto(viento string) PuntoCardinal {
	switch viento {
	case "norte":
		return norte
	case "sur":
		return sur
	case "este":
		return este
	case "oeste":
		return oeste
	case "noroeste":
		return noroeste
	case "sureste":
		return sureste
	case "noreste":
		return noreste
	case "suroeste":
		return suroeste
	}
	return -1
}

func ConseguirStringOpuesto(x PuntoCardinal) string {
	switch x {
	case norte:
		return "sur"
	case sur:
		return "norte"
	case este:
		return "oeste"
	case oeste:
		return "este"
	case noroeste:
		return "sureste"
	case sureste:
		return "noroeste"
	case noreste:
		return "suroeste"
	case suroeste:
		return "noreste"
	}
	return "desconocido"
}

func ConseguirPorIndice(x PuntoCardinal) string {
	if x%2 == 0 {
		x++
	} else {
		x--
	}
	return ConseguirString(x)
}

func ConseguirString(x PuntoCardinal) string {
	switch x {
	case norte:
		return "norte"
	case sur:
		return "sur"
	case este:
		return "este"
	case oeste:
		return "oeste"
	case noroeste:
		return "noroeste"
	case sureste:
		return "sureste"
	case noreste:
		return "noreste"
	case suroeste:
		return "suroeste"
	}
	return "desconocido"
}

func (x PuntoCardinal) String() string {
	switch x {
	case norte:
		return "sur"
	case sur:
		return "norte"
	case este:
		return "oeste"
	case oeste:
		return "este"
	case noroeste:
		return "sureste"
	case sureste:
		return "noroeste"
	case noreste:
		return "suroeste"
	case suroeste:
		return "noreste"
	}
	return "desconocido"
}

func main() {
	var viento string
	var opc int
	fmt.Print("Ingrese el la punto cardinal del cual viene el viento: ")
	fmt.Scanln(&viento)
	fmt.Print("Ingrese la opción (1: switch, 2: índice, 3: map, 4: interfaz): ")
	fmt.Scanln(&opc)
	p := ConseguirPunto(viento)
	fmt.Print("El viento se dirije hacia el ")
	switch opc {
	case 1:
		fmt.Println(ConseguirStringOpuesto(p))
	case 2:
		fmt.Println(ConseguirPorIndice(p))
	case 3:
		fmt.Println(puntos[p])
	case 4:
		fmt.Println(p)
	default:
		fmt.Println("opción inválida")
	}
}
