package main

import "fmt"

const (
	categA = "A"
	categB = "B"
	categC = "C"
)

func calcTipo(tipo string) float64 {
	switch tipo {
	case categA:
		return 3000.0
	case categB:
		return 1500.0
	case categC:
		return 1000.0
	}
	return 0
}

func calcAumento(tipo string) float64 {
	switch tipo {
	case categA:
		return 0.5
	case categB:
		return 0.2
	case categC:
		return 0
	}
	return 0
}

func calcSalario(minutos int, tipo string) float64 {
	horas := float64(minutos) / 60
	neto := horas * calcTipo(tipo)
	salarioFinal := neto + neto*calcAumento(tipo)
	return salarioFinal
}

func main() {
	fmt.Println(calcSalario(60, categC))
}
