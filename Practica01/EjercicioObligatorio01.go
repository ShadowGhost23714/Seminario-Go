/*
Realice las modificaciones necesarias al ejercicio anterior para que en
lugar de reemplazar la palabra “jueves” por “martes” ahora se
reemplace “miércoles” por “automóvil”. Piense qué impacto tuvieron
esas modificaciones en el programa que había realizado.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var fraseBase string
	var fraseCopia string = "automóvil"
	fmt.Print("Ingrese el día miércoles: ")
	fmt.Scan(&fraseBase)
	r1 := []rune(fraseBase)
	r2 := []rune(fraseCopia)

	for i := 0; i < len(r1); i++ {
		if unicode.IsUpper(r1[i]) {
			r2[i] = unicode.ToUpper(r2[i])
		}
	}

	fmt.Print("Palabra modificada: ", string(r2))
}
