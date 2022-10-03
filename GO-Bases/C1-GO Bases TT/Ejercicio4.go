package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	//imprimir edad benjamin
	employeeName := "Benjamin"
	fmt.Println("Nombre empleado: ", employeeName, "\nEdad: ", employees["Benjamin"])

	//numero de empleados mayores a 21
	cont := 0
	for _, element := range employees {
		if element > 21 {
			cont++
		}
	}
	fmt.Print(cont, " empleados mayores a 21 años")

	//agregar a Federico
	employees["Federico"] = 25

	//Eliminar a pedro
	delete(employees, "Pedro")

}
