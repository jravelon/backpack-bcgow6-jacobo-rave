package main

import "fmt"

func main() {
	nombre := "Jacobo"
	nombreSlc := []rune(nombre)
	fmt.Println("Tamaño nombre: ", len(nombre))
	for i := 0; i < len(nombreSlc); i++ {
		fmt.Println(string(nombreSlc[i]))
	}
}
