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

// En esta ocasión vamos a añadir un producto al slice cargado en memoria. Dentro de la ruta /products añadimos el método POST, al cual vamos a enviar en el cuerpo de la request el nuevo producto. El mismo tiene ciertas restricciones, conozcámoslas:
// 1. No es necesario pasar el Id, al momento de añadirlo se debe inferir del estado de la lista de productos, verificando que no se repitan ya que debe ser un campo único.
// 2. Ningún dato puede estar vacío, exceptuando is_published (vacío indica un valor false).
// 3. El campo code_value debe ser único para cada producto.
// 4. Los tipos de datos deben coincidir con los definidos en el planteo del problema.
// 5. La fecha de vencimiento debe tener el formato: XX/XX/XXXX, además debemos verificar que día, mes y año sean valores válidos.
// Recordá: si una consulta está mal formulada por parte del cliente, el status code cae en los 4XX. 

import (
	"02-TM-post/cmd/server/controller"
	"02-TM-post/internal/service"
	"net/http"


	"github.com/gin-gonic/gin"
)

func main() {
	service.Init()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	productsRouter := router.Group("/products")

	productsRouter.GET("/", controller.GetAllProducts())
	productsRouter.POST("/", controller.NewProduct())
	productsRouter.GET("/:id", controller.GetProductById())
	productsRouter.GET("/search", controller.GetProductsByMinPrice())

	router.Run(":8080")
}

