package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Jacobo",
		})
	})
	router.Run()
}
