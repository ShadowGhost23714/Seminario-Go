package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func InvertirMayusMins(palabra string) string {
	letras := []rune(palabra)
	for i := 0; i < len(letras); i++ {
		if unicode.IsUpper(letras[i]) {
			letras[i] = unicode.ToLower(letras[i])
		} else {
			letras[i] = unicode.ToUpper(letras[i])
		}
	}
	return string(letras)
}

func main() {
	var palabra string
	fmt.Print("Ingrese una palabra: ")
	fmt.Scanln(&palabra)

	fmt.Print("Ingrese una frase: ")

	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimRight(frase, "\n")

	fraseMin := strings.ToLower(frase)
	palabraMin := strings.ToLower(palabra)

	var resultado strings.Builder

	i := 0
	for i < len(frase) {
		pos := strings.Index(fraseMin[i:], palabraMin) // guardo la posición de la palabra en la frase relativa al slice
		if pos == -1 {                                 // no hay mas coincidencias, caso i = 0 se copia la frase
			resultado.WriteString(frase[i:])
			i = len(frase) // para no usar un break en el for
		} else {
			pos = pos + i
			resultado.WriteString(frase[i:pos]) // copio la parte de la frase antes de la palabra
			original := frase[pos : pos+len(palabra)]
			resultado.WriteString(InvertirMayusMins(original)) // invierto la palabra
			i = pos + len(palabra)                             // la posicion donde ahora se va a hacer fraseMin[i:] es despues de la palabra encontrada
		}
	}

	fmt.Println("Frase modificada:", resultado.String())
}
