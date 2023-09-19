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
	"05-TM-middleware-docs/cmd/server/handler"
	"05-TM-middleware-docs/cmd/server/middleware"
	"05-TM-middleware-docs/cmd/server/router"
	"05-TM-middleware-docs/docs"
	"05-TM-middleware-docs/internal/product"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp Products API
// @version 1.0
// @description This API is built incrementally over many bootcamp exercises
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @licence.name Apache 2.0
// @licence.url https://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load("./.env")

	if err != nil {
		panic(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())

	host := os.Getenv("HOST")
	docs.SwaggerInfo.Host = host
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	token := os.Getenv("TOKEN")
	r.Use(middleware.TokenAuthMiddleware(token))

	jsonDbPath := os.Getenv("JSON_DB_PATH")
	repository := product.NewProductRepositoryJson(jsonDbPath)
	service := product.NewProductServiceDefault(repository)
	handler := handler.NewProductHandlerDefault(service)

	routerProduct := router.NewProductRouter(r, handler)
	routerProduct.MapRoutes()

	err = r.Run(host)

	if err != nil {
		panic(err)
	}
}
