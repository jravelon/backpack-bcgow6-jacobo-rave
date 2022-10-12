package main

import "fmt"

type Error struct {
}

func (e *Error) Error() string {
	return "El salario ingresado no alcanza el minimo imponible"
}

func main() {
	var salary int = 200000
	if salary < 150000 {
		fmt.Println(&Error{})
	}
}
