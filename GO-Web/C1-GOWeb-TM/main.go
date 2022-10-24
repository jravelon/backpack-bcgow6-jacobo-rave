package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type fecha struct {
	Año int
	Mes string
	Día int
}
type users struct {
	Id            int
	Nombre        string
	Apellido      string
	Email         string
	Edad          int
	Altura        float64
	Activo        bool
	FechaCreacion fecha
}

func Salute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hola Jacobo"})
}

func GetAll(ctx *gin.Context) {
	jsonData, _ := os.ReadFile("Users.json")
	var u users
	if err := json.Unmarshal(jsonData, &u); err != nil {
		log.Fatal(err)
	}
	UsersEj := []users{u}
	ctx.JSON(200, UsersEj)
}

func main() {
	router := gin.Default()

	router.GET("/", Salute)
	router.GET("/products", GetAll)
	router.Run()
}
