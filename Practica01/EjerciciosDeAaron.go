package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func eje9() {
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

func ejeObli1() {
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

func invertirPalabra(palabra string) string {
	r := []rune(palabra)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ejeObli2() {
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

func main() {
	var opcion int
	fmt.Print("ingrese un numero de la practica a resolver: ")
	fmt.Scan(&opcion)
	switch opcion {
	case 9:
		eje9()
	case 1:
		ejeObli1()
	case 2:
		ejeObli2()
	}
}
