/*
Realice un programa que reciba una frase e imprima en pantalla la
misma frase con cada una de las palabras invertidas siempre que su
ubicación sea impar en la frase comenzando a contar las palabras
desde 1, por ejemplo, si la frase ingresada es:
Qué lindo día es hoy.
El programa imprimirá:
éuQ lindo aíd es yoh.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func invertirPalabra(palabra string) string {
	r := []rune(palabra)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Print("Ingrese una frase: ")
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)

	palabras := strings.Fields(frase)

	for i, palabra := range palabras {
		if i%2 == 0 {
			palabras[i] = invertirPalabra(palabra)
		}
	}

	resultado := strings.Join(palabras, " ")
	fmt.Println("Frase modificada:", resultado)
}
