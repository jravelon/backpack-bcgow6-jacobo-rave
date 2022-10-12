package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 200000
	if salary < 150000 {
		err := errors.New("El salario ingresado no alcanza el minimo imponible")
		fmt.Println(err)
	}
}
