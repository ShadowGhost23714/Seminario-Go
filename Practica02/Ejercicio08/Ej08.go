/*
8. Escriba un programa que implemente la función:

func Convert ( v int , b int ) string

a. La función debe convertir el entero v, en un string en su representación
en base b. El string será el valor de retorno. Por ejemplo si se invoca:

s = Convert ( 23 , 2 )

En s se almacenaría el valor “10111”.

Operación:

25 / 8 = 3, residuo 1
3 / 8 = 0, residuo 3
El número en base 8 es 31.

La base debe ser mayor que 1 y menor que 37 dado que irı́a de base-2
hasta base-36, que usarı́a los dígitos:

“0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ”

b. Re-implemente la función considerando que v puede ser negativo, y se
empleará el mismo símbolo (-) para su representación en base-b.
*/

package main

import (
	"fmt"
	"strconv"
)

const digitos = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Convert4(v int, b int) string { // antigua solución
	cociente := v / b
	residuo := v % b
	if cociente == 0 {
		return strconv.Itoa(residuo)
	}
	return Convert3(cociente, b) + strconv.Itoa(residuo)
}

func Convert3(v int, b int) string { // antigua solución
	if b < 2 || b > 36 {
		fmt.Println("Base no válida")
	}
	if v == 0 {
		return ""
	}
	return Convert2(v/b, b) + string(digitos[v%b])
}

func Convert2(v int, b int) string { // Punto A
	if b < 2 || b > 36 {
		fmt.Println("Base no válida")
	}
	var resultado string
	for v > 0 {
		resultado = string(digitos[v%b]) + resultado
		v = v / b
	}
	return resultado
}

func Convert(v int, b int) string { // Punto B
	if b < 2 || b > 36 {
		fmt.Println("Base no válida")
	}
	var resultado string
	negativo := false
	if v < 0 {
		v = -v
		negativo = true
	}
	for v > 0 {
		resultado = string(digitos[v%b]) + resultado
		v = v / b
	}
	if negativo {
		resultado = "-" + resultado
	}
	return resultado
}

func main() {
	fmt.Println("Número 25 y base 8 = ", Convert(25, 8))
	fmt.Println("Número -25 y base 8 = ", Convert(-25, 8))
	fmt.Println("Número 23 y base 2 = ", Convert(23, 2))
	fmt.Println("Número -23 y base 2 = ", Convert(-23, 2))
}
