package main

import "fmt"

func main() {
	edad := 23
	esEmpleado := true
	antiguedad := 2
	sueldoInteres := 100000

	if edad > 22 && esEmpleado == true && antiguedad > 1 {
		fmt.Println("Puede acceder a prestamo")
		if sueldoInteres > 100000 {
			fmt.Print("No se cobra interes")
		} else {
			fmt.Println("Se le cobrará interes")
		}
	} else {
		fmt.Println("No puede acceder a un crédito")
	}
}
