package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Fecha struct {
	Año int
	Mes string
	Día int
}
type Users struct {
	Id            int
	Nombre        string
	Apellido      string
	Email         string
	Edad          int
	Altura        float64
	Activo        bool
	FechaCreacion Fecha
}

func main() {

	jsonData, _ := os.ReadFile("Users.json")
	var u Users

	if err := json.Unmarshal(jsonData, &u); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product:", u)
	fmt.Println("Product: ", u.Email)

}
