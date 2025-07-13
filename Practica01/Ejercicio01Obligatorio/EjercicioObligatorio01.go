/*
Realice las modificaciones necesarias al ejercicio anterior para que en
lugar de reemplazar la palabra “jueves” por “martes” ahora se
reemplace “miércoles” por “automóvil”. Piense qué impacto tuvieron
esas modificaciones en el programa que había realizado.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func MayusMins(palabra string, patron string) string {
	r := []rune(palabra)
	p := []rune(patron)
	resultado := make([]rune, len(r))

	for i := 0; i < len(r); i++ {
		if i < len(p) && unicode.IsUpper(p[i]) {
			resultado[i] = unicode.ToUpper(r[i])
		} else {
			resultado[i] = unicode.ToLower(r[i])
		}
	}
	return string(resultado)
}

func sinTildes(s string) string {
	reemplazos := map[rune]rune{
		'á': 'a', 'é': 'e', 'í': 'i', 'ó': 'o', 'ú': 'u',
		'Á': 'A', 'É': 'E', 'Í': 'I', 'Ó': 'O', 'Ú': 'U',
	}
	runes := []rune(s)
	for i, r := range runes {
		if nuevo, ok := reemplazos[r]; ok {
			runes[i] = nuevo
		}
	}
	return string(runes)
}

func main() {
	fmt.Print("Ingrese una frase: ")
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)

	palabras := strings.Fields(frase)

	for i := 0; i < len(palabras); i++ {
		if strings.EqualFold(sinTildes(palabras[i]), "miercoles") {
			palabras[i] = MayusMins("automóvil", palabras[i])
		}
	}
	resultado := strings.Join(palabras, " ")
	fmt.Println("Palabra modificada:", resultado)
}
