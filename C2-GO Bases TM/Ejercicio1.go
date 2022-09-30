package main

import "fmt"

func impuestoSalario(salario float64) float64 {
	var salarioConImpuesto float64
	if salario > 50000 && salario <= 150000 {
		salarioConImpuesto = salario - (salario * 0.17)
	} else if salario > 150000 {
		salarioConImpuesto = salario - (salario*0.17 + salario*0.1)
	} else {
		fmt.Println("No se aplica impuesto al salario")
	}
	return salarioConImpuesto
}

func main() {
	fmt.Println("El salario con el descuento de impuesto seria:", impuestoSalario(150000))
}
