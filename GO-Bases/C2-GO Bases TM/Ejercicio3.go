package main

import (
	"fmt"
)

const (
	categoriaA = "A"
	categoriaB = "B"
	categoriaC = "C"
)

func conversor(minutos int) float64 {
	return float64(minutos) / 60
}

func salario(minutos int, categoria string) float64 {
	horas := conversor(minutos)
	if categoria == categoriaC {
		return horas * 1000
	} else if categoria == categoriaB {
		salary := horas * 1500
		return salary + (salary * 0.2)
	} else {
		salary := horas * 3000
		return salary + (salary * 0.5)
	}
}

func main() {
	fmt.Println(salario(60, categoriaB))
}
