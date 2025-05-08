package main

import (
	"fmt"
	"unicode"
)

func main() {
	letra := 'a' // rune
	mayuscula := unicode.ToUpper(letra)
	minuscula := unicode.ToLower(letra)

	fmt.Printf("Original: %c\n", letra)
	fmt.Printf("Mayúscula: %c\n", mayuscula)
	fmt.Printf("Minúscula: %c\n", minuscula)
}
