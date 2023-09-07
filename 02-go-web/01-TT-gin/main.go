package main

// Debemos crear un repositorio en github.com para poder subir nuestros avances. Este repositorio es el que vamos a utilizar para llevar lo que realicemos durante las distintas prácticas de Go Web.
// Primero debemos clonar el repositorio creado, luego iniciar nuestro proyecto de go con con el comando go mod init.
// El siguiente paso será crear un archivo main.go donde deberán cargar en una slice, desde un archivo JSON, los datos de productos. Esta slice se debe cargar cada vez que se inicie la API para realizar las distintas consultas.
// El archivo para trabajar es el siguiente:
// https://drive.google.com/file/d/1oZ71o1BCml2EGhAQ31wvtv-RGZzTQjaW/view

// Vamos a levantar un servidor en el puerto 8080. Para probar nuestros endpoints haremos uso de postman.
// Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
// Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
// Crear una ruta /products/:id que nos devuelva un producto por su id.
// Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.

import (
	"01-TT-gin/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	service.Init()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	productsRouter := router.Group("/products")
	productsRouter.GET("/", getAllProducts)
	productsRouter.GET("/:id", getProductById)
	productsRouter.GET("/search", getProductsByMinPrice)

	router.Run(":8080")
}

func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAllProducts())
	return
}

func getProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, err := service.GetProductById(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}

	return
}

func getProductsByMinPrice(c *gin.Context) {
	param := c.Query("priceGt")

	if param == "" {
		c.JSON(http.StatusOK, service.GetAllProducts())
		return
	}

	minPrice, err := strconv.ParseFloat(param, 64)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, service.GetProductsByMinPrice(minPrice))
	return
}
