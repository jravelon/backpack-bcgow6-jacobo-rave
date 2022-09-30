package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...float64) (float64, error) {
	var promedio float64 = 0
	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("Nota negativa")
		}
		promedio += value
	}
	promedio = promedio / float64(len(notas))

	return promedio, nil
}

func main() {
	fmt.Print(promedio(1, 2, 3, 4, 5))
}
