package main

import "fmt"

type student struct {
	name     string
	lastName string
	dni      string
	date     string
}

func (s student) details() {

	fmt.Printf("Nombre: %v\n", s.name)
	fmt.Printf("Apellido: %v\n", s.lastName)
	fmt.Printf("DNI: %v\n", s.dni)
	fmt.Printf("Fecha: %v\n", s.date)
}

func main() {
	s := student{"Jacobo", "rave", "10", "13 marzo"}
	s.details()
}