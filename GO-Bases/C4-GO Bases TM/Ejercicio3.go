package main

import (
	"fmt"
)

func main() {
	var salary int = 20000
	if salary < 150000 {
		err := fmt.Errorf("el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
		fmt.Println(err)
	}
}
