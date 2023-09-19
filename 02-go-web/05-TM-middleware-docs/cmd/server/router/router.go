package router

import (
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetAllProducts() gin.HandlerFunc
	GetProductById() gin.HandlerFunc
	GetProductsByMinPrice() gin.HandlerFunc
	NewProduct() gin.HandlerFunc
	ReplaceProduct() gin.HandlerFunc
	UpdateProduct() gin.HandlerFunc
	DeleteProduct() gin.HandlerFunc
}

type ProductRouter struct {
	r       *gin.Engine
	handler ProductHandler
}

func NewProductRouter(r *gin.Engine, handler ProductHandler) *ProductRouter {
	return &ProductRouter{
		r:       r,
		handler: handler,
	}
}

func (rp *ProductRouter) MapRoutes() {
	productsRouter := rp.r.Group("/products")
	{
		productsRouter.GET("/", rp.handler.GetAllProducts())
		productsRouter.POST("/", rp.handler.NewProduct())
		productsRouter.GET("/:id", rp.handler.GetProductById())
		productsRouter.PUT("/:id", rp.handler.ReplaceProduct())
		productsRouter.PATCH("/:id", rp.handler.UpdateProduct())
		productsRouter.DELETE("/:id", rp.handler.DeleteProduct())
		productsRouter.GET("/search", rp.handler.GetProductsByMinPrice())
	}
}
