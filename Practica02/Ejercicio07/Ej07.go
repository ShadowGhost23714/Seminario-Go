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
	salir := false
	letras := 0
	numeros := 0
	especiales := 0
	mayus := 0
	minus := 0
	m := make(map[string]int)
	reader := bufio.NewReader(os.Stdin) // Crea un lector para leer texto desde la entrada estándar (teclado)
	var ultimo rune
	for !salir {
		fmt.Print("Ingrese un caracter: ")
		letra, _ := reader.ReadString('\n')    // Lee una línea completa hasta que se presiona ENTER
		letra = strings.TrimRight(letra, "\n") // Borra del string el salto de linea
		char := rune(letra[0])                 // Toma el primer carácter ingresado y lo convierte a rune (carácter Unicode)
		switch {
		case unicode.IsLetter(char):
			{
				letras++
				if unicode.IsUpper(char) {
					mayus++
				} else {
					minus++
				}
				if ultimo == 'C' && char == 'R' {
					salir = true
				}
			}
		case unicode.IsDigit(char):
			m[string(char)]++ // Puede ser también m[letra]++
			numeros++
		default:
			especiales++
		}
		ultimo = char
	}
	fmt.Println("Letras = ", letras)
	fmt.Println("Números = ", numeros)
	fmt.Println("Caracteres especiales = ", especiales)
	fmt.Println("Mayúsculas = ", mayus)
	fmt.Println("Minúsculas = ", minus)
	fmt.Print(m)
}

/*
func Antiguo() { // Versión vieja del jercicio viejo
	pri := false
	salir := false
	letras := 0
	numeros := 0
	especiales := 0
	reader := bufio.NewReader(os.Stdin) // Crea un lector para leer texto desde la entrada estándar (teclado)
	for !salir {
		fmt.Print("Ingrese un caracter: ")  // Lee una línea completa hasta que se presiona ENTER
		input, _ := reader.ReadString('\n') // Elimina espacios y saltos de línea del inicio y fin
		input = strings.TrimSpace(input)    // Toma el primer carácter ingresado y lo convierte a rune (carácter Unicode)
		char := rune(input[0])
		switch {
		case unicode.IsLetter(char):
			{
				letras++
				if char == 'C' {
					pri = true
				} else if pri && char == 'R' {
					salir = true
				} else {
					pri = false
				}
			}
		case unicode.IsDigit(char):
			numeros++
		default:
			especiales++
		}
	}
	fmt.Println("Letras = ", letras)
	fmt.Println("Números = ", numeros)
	fmt.Println("Caracteres especiales = ", especiales)
}
*/
