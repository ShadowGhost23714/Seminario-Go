/*
Realice un programa que reciba una palabra como argumento y lee de
la entrada una frase. Luego, el programa debe imprimir la frase que
leyó con cada una de las ocurrencias de la palabra con las mayúsculas
y minúsculas invertidas. Por ejemplo, si la frase es:
“Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO”
y la palabra es “PEQUEÑO” entonces el programa imprimirá:
“Parece PEQueÑO, pero no es tan PEQUEñO el pequeño”
Tenga en cuenta que la palabra a buscar puede ser ingresada con
mayúsculas y minúsculas mezcladas.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func MayusMins(palabra string, msj string) string {
	r := []rune(palabra)
	r2 := []rune(msj)
	for i := 0; i < len(r2); i++ {
		if unicode.IsUpper(r2[i]) {
			r[i] = unicode.ToUpper(r[i])
		}
	}
	return string(r)
}

func main() {
	var msj string
	fmt.Scan("Ingrese una palabra: ")
	fmt.Scan(msj)
	fmt.Print("Ingrese ahora una frase: ")
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)

	palabras := strings.Fields(frase)

	for i, palabra := range palabras {
		if palabra == msj {
			palabras[i] = MayusMins(palabra, msj)
		}
	}

	resultado := strings.Join(palabras, " ")
	fmt.Println("Frase modificada:", resultado)
}
