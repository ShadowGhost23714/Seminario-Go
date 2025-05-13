/*
7. Se debe leer una secuencia de caracteres que finaliza con CR e informar la
cantidad de letras, números y caracteres especiales leídos.

a. Modificar el programa anterior para que cuente de forma separada
mayúsculas de minúsculas.

b. Modificar para que, además, cuente de forma separada las ocurrencias
de cada dígito decimal. Utilice la estructura de datos Map.
Sub-objetivo: Operaciones sobre caracteres (runas) y estructuras de
control.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	pri := false
	reader := bufio.NewReader(os.Stdin)
	letras := 0
	numeros := 0
	especiales := 0
	for {
		fmt.Print("Ingrese una letra: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		char := rune(input[0])
		switch {
		case unicode.IsLetter(char):
			letras++
			if char == 'C' {
				pri = true
			} else if pri && char == 'R' {
				break
			} else {
				pri = false
			}
		case unicode.IsDigit(char):
			numeros++
		default:
			especiales++
		}
	}
}
