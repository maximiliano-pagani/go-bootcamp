package main

// Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
// 1. El endpoint deberá ser de método GET
// 2. La respuesta de “pong” deberá ser enviada como texto, NO como JSON

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":8080")
}
