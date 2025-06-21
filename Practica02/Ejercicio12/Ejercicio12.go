/*
12. Un banco dispone de un listado en donde almacena la información de
aquellos clientes que vienen a pagar impuestos. De cada cliente conoce: DNI,
Nombre, Apellido, código de impuesto a pagar (A, B, C o D) y el monto a pagar.
Se pide:

a. Realizar la atención de los clientes hasta que se recauden al menos
10.000 pesos o hasta que se terminen los clientes.

b. Al finalizar la atención informar el código de impuesto que más veces se
pagó por los clientes que fueron atendidos.

c. Al finalizar la atención informar, en caso de que hayan quedado, la
cantidad de clientes sin atender. */

package main

import "fmt"

type cliente struct {
	dni   int
	nom   string
	ape   string
	cod   string
	monto float64
}

func iniciarSlice() []cliente {
	var clientes []cliente
	clientes = append(clientes, cliente{101, "Pepe", "Argento", "A", 2000})
	clientes = append(clientes, cliente{102, "Moni", "Argento", "B", 2000})
	clientes = append(clientes, cliente{103, "Coqui", "Argento", "C", 2000})
	clientes = append(clientes, cliente{104, "Paola", "Argento", "D", 2000})
	clientes = append(clientes, cliente{105, "María Elena", "Fuseneco", "A", 3000})
	clientes = append(clientes, cliente{106, "Dardo", "Fuseneco", "A", 2000})
	return clientes
}

func atencionCliente(clientes []cliente) (string, int) {
	var montoTotal float64
	contadores := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0}
	cantAtendidos := 0
	for i := 0; montoTotal < 10000 && i < len(clientes); i++ {
		montoTotal += (clientes[i].monto)
		contadores[clientes[i].cod]++
		fmt.Println("Se atendio a ", clientes[i].nom, " ", clientes[i].ape)
		cantAtendidos++
	}
	var maxCod string
	maxValor := -1
	for cod, valor := range contadores {
		if valor > maxValor {
			maxCod = cod
			maxValor = valor
		}
	}
	return maxCod, len(clientes) - cantAtendidos
}

func main() {
	clientes := iniciarSlice()
	max, restantes := atencionCliente(clientes)
	fmt.Println("El impuesto que mas se pago fue ", max)
	if restantes == 0 {
		fmt.Println("No quedaron clientes sin atender")
	} else {
		fmt.Println("Clientes restantes sin antender: ", restantes)
	}
}

/*
	alternativa del primer For de la funcion atencionCliente

	for _, c := range clientes {
		if montoTotal >= 10000 {
			break
		}
		montoTotal += c.monto
		contadores[c.cod]++
		fmt.Printf("Se atendió a %s %s\n", c.nom, c.ape)
		cantAtendidos++
	}

*/
