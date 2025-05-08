package main

import (
	"errors"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	cantidad = 10
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func leerValor(valor float64) (float64, error) {

	if valor > 90 || valor < -20 {
		return 1, errors.New("temperatura invalida")
	} else {
		return valor, nil
	}

}

func main() {
	var temperaturas [cantidad]float64
	var alta, normal, baja int
	var err error
	var valor float64
	tempMax := -20.0
	tempMin := 90.0
	grupos := map[string]float64{
		"bajaM":   0,
		"normalM": 0,
		"altaM":   0,
	}

	fmt.Println()
	fmt.Println("Ingrese las", cantidad, "temperaturas: ")

	for i := 0; i < cantidad; i++ {
		fmt.Scan(&temperaturas[i])
		valor, err = leerValor(temperaturas[i])
		if err == nil {
			switch {
			case valor > 37.5:
				alta++
				grupos["altaM"] += valor
			case 37.5 > valor && 36.0 < valor:
				normal++
				grupos["normalM"] += valor
			case valor < 36:
				baja++
				grupos["bajaM"] += valor
			}
			if tempMax < valor {
				tempMax = valor
			}
			if tempMin > valor {
				tempMin = valor
			}
		}
	}

	fmt.Println()
	fmt.Println("Temperaturas ingresadas convertidas a Fahrenheit")
	for i := 0; i < cantidad; i++ {
		c := Celsius(temperaturas[i])
		f := CToF(c)
		valor, err = leerValor(temperaturas[i])
		if err == nil {
			temperaturas[i] = float64(f)
			fmt.Printf("La temperatura es %.2f\n", temperaturas[i])
		}
	}

	fmt.Println()
	fmt.Println("Grupos de temperaturas :", grupos)
	fmt.Println()
	fmt.Printf("El promedio entre la mayor temperatura y la menor temperatura es de %.2f\n", ((tempMax + tempMin) / 2))
	fmt.Println()
}
