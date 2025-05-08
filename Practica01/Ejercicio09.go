/*
Realice un programa que reciba una frase e imprima en pantalla
la misma frase reemplazando las ocurrencias de “jueves” por
“martes” respetando las letras minúsculas o mayúsculas de la
palabra original en su posición correspondiente. Por ejemplo, se
reemplaza “Jueves” por “Martes” o “jueveS” por “marteS”.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var diaBase string
	var diaCopia string = "martes"
	fmt.Println("Ingrese el día jueves.")
	fmt.Scan(&diaBase)
	r1 := []rune(diaBase)
	r2 := []rune(diaCopia)

	for i := 0; i < len(r1); i++ {
		if unicode.IsUpper(r1[i]) {
			r2[i] = unicode.ToUpper(r2[i])
		}
	}

	fmt.Println("Día modificado: ", string(r2))
}
