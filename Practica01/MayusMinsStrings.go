package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Hola Mundo"
	fmt.Println("Todo en mayúsculas:", strings.ToUpper(texto))
	fmt.Println("Todo en minúsculas:", strings.ToLower(texto))
}
