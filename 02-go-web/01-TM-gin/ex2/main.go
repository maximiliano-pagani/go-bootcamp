package main

// Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hola + nombre + apellido”

// 1. El endpoint deberá ser de método POST
// 2. Se deberá usar el package JSON para resolver el ejercicio
// 3. La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
// 4. La estructura deberá ser como esta:
// {
// 		“nombre”: “Andrea”,
// 		“apellido”: “Rivas”
// }

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"nombre"`
	Lastname string `json:"apellido"`
}

func main() {
	router := gin.Default()

	router.POST("/saludo", func(c *gin.Context) {
		var user User

		err := json.NewDecoder(c.Request.Body).Decode(&user)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		reply := fmt.Sprintf("Hola %s %s", user.Name, user.Lastname)
		c.String(http.StatusOK, reply)
	})

	router.Run(":8080")
}
