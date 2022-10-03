package main

import (
	"github.com/gin-gonic/gin"
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

func GetAll(ctx *gin.Context) {
	UsersEj := []Users{
		{
			Id:       1212,
			Nombre:   "J",
			Apellido: "R",
			Email:    "j.r@mercadolibre.com.co",
			Edad:     19,
			Altura:   1.85,
			Activo:   true,
			FechaCreacion: Fecha{
				Año: 2022,
				Mes: "Octubre",
				Día: 3,
			},
		},
		{
			Id:       1212,
			Nombre:   "J",
			Apellido: "R",
			Email:    "j.r@mercadolibre.com.co",
			Edad:     19,
			Altura:   1.85,
			Activo:   true,
			FechaCreacion: Fecha{
				Año: 2022,
				Mes: "Octubre",
				Día: 3,
			},
		},
	}
	ctx.JSON(200, UsersEj)
}

func main() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Jacobo",
		})
	})

	router.GET("/products", GetAll)
	router.Run()
}
