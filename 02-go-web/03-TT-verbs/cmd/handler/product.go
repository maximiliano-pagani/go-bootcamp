package handler

import (
	"03-TT-verbs/internal/domain"
	"03-TT-verbs/internal/product"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductHandlerDefault struct {
	service product.ProductService
}

func NewProductHandlerDefault(service product.ProductService) *ProductHandlerDefault {
	handler := &ProductHandlerDefault{service: service}
	return handler
}

type RequestProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type DataProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseProduct struct {
	Code int                  `json:"code"`
	Msg  string               `json:"message"`
	Data *DataProductResponse `json:"data"`
}

func (h *ProductHandlerDefault) GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := h.service.GetAllProducts()

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, results)
		return
	}
}

func (h *ProductHandlerDefault) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		result, err := h.service.GetProductById(id)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, result)
		return
	}
}

func (h *ProductHandlerDefault) GetProductsByMinPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Query("priceGt")

		if param == "" {
			if results, err := h.service.GetAllProducts(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			} else {
				c.JSON(http.StatusOK, results)
			}

			return
		}

		minPrice, err := strconv.ParseFloat(param, 64)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		results, err := h.service.GetProductsByMinPrice(minPrice)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, results)
		return
	}
}

func (h *ProductHandlerDefault) NewProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody := &RequestProduct{}
		err := c.ShouldBindJSON(reqBody)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = h.validateRequestProduct(reqBody)

		if err != nil {
			code := http.StatusBadRequest
			respBody := &ResponseProduct{Code: code, Msg: err.Error()}
			c.AbortWithStatusJSON(code, respBody)
			return
		}

		addedProduct, err := h.service.NewProduct(
			&domain.Product{
				Name:        reqBody.Name,
				Quantity:    reqBody.Quantity,
				Code:        reqBody.Code,
				IsPublished: reqBody.IsPublished,
				Expiration:  reqBody.Expiration,
				Price:       reqBody.Price,
			},
		)

		if err != nil {
			code := http.StatusBadRequest
			respBody := &ResponseProduct{Code: code, Msg: err.Error()}
			c.AbortWithStatusJSON(code, respBody)
			return
		}

		code := http.StatusCreated
		respBody := &ResponseProduct{
			Code: code,
			Msg:  "Created",
			Data: &DataProductResponse{
				Id:          addedProduct.Id,
				Name:        addedProduct.Name,
				Quantity:    addedProduct.Quantity,
				Code:        addedProduct.Code,
				IsPublished: addedProduct.IsPublished,
				Expiration:  addedProduct.Expiration,
				Price:       addedProduct.Price,
			},
		}

		c.JSON(code, respBody)
		return
	}
}

func (h *ProductHandlerDefault) ReplaceProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		reqBody := &RequestProduct{}
		err = c.ShouldBindJSON(reqBody)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = h.validateRequestProduct(reqBody)

		if err != nil {
			code := http.StatusBadRequest
			respBody := &ResponseProduct{Code: code, Msg: err.Error()}
			c.AbortWithStatusJSON(code, respBody)
			return
		}

		replacedProduct, err := h.service.ReplaceProduct(
			&domain.Product{
				Id:          id,
				Name:        reqBody.Name,
				Quantity:    reqBody.Quantity,
				Code:        reqBody.Code,
				IsPublished: reqBody.IsPublished,
				Expiration:  reqBody.Expiration,
				Price:       reqBody.Price,
			},
		)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		code := http.StatusOK
		respBody := &ResponseProduct{
			Code: code,
			Msg:  "Replaced",
			Data: &DataProductResponse{
				Id:          replacedProduct.Id,
				Name:        replacedProduct.Name,
				Quantity:    replacedProduct.Quantity,
				Code:        replacedProduct.Code,
				IsPublished: replacedProduct.IsPublished,
				Expiration:  replacedProduct.Expiration,
				Price:       replacedProduct.Price,
			},
		}

		c.JSON(code, respBody)
		return
	}
}

func (h *ProductHandlerDefault) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		originalProduct, err := h.service.GetProductById(id)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		reqBody := &RequestProduct{
			Name:        originalProduct.Name,
			Quantity:    originalProduct.Quantity,
			Code:        originalProduct.Code,
			IsPublished: originalProduct.IsPublished,
			Expiration:  originalProduct.Expiration,
			Price:       originalProduct.Price,
		}

		err = c.ShouldBindJSON(reqBody)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = h.validateRequestProduct(reqBody)

		if err != nil {
			code := http.StatusBadRequest
			respBody := &ResponseProduct{Code: code, Msg: err.Error()}
			c.AbortWithStatusJSON(code, respBody)
			return
		}

		updatedProduct, err := h.service.UpdateProduct(
			&domain.Product{
				Id:          id,
				Name:        reqBody.Name,
				Quantity:    reqBody.Quantity,
				Code:        reqBody.Code,
				IsPublished: reqBody.IsPublished,
				Expiration:  reqBody.Expiration,
				Price:       reqBody.Price,
			},
		)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		code := http.StatusOK
		respBody := &ResponseProduct{
			Code: code,
			Msg:  "Updated",
			Data: &DataProductResponse{
				Id:          updatedProduct.Id,
				Name:        updatedProduct.Name,
				Quantity:    updatedProduct.Quantity,
				Code:        updatedProduct.Code,
				IsPublished: updatedProduct.IsPublished,
				Expiration:  updatedProduct.Expiration,
				Price:       updatedProduct.Price,
			},
		}

		c.JSON(code, respBody)
		return
	}
}

func (h *ProductHandlerDefault) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = h.service.DeleteProduct(id)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.Status(http.StatusNoContent)
		return
	}
}

func (h *ProductHandlerDefault) validateRequestProduct(product *RequestProduct) error {
	switch {
	case product.Name == "":
		return errors.New("Invalid Name")
	case product.Code == "":
		return errors.New("Invalid Code")
	case product.Quantity == 0:
		return errors.New("Invalid Quantity")
	case product.Expiration == "" || !h.isValidExpirationDate(product.Expiration):
		return errors.New("Invalid Expiration Date")
	case product.Price == 0:
		return errors.New("Invalid Price")
	}
	return nil
}

func (h *ProductHandlerDefault) isValidExpirationDate(date string) bool {
	subStrings := strings.Split(date, "/")

	if len(subStrings) != 3 {
		return false
	}

	day, dayErr := strconv.Atoi(subStrings[0])
	month, monthErr := strconv.Atoi(subStrings[1])
	year, yearErr := strconv.Atoi(subStrings[2])

	switch {
	case dayErr != nil || monthErr != nil || yearErr != nil:
		return false
	case day < 1 || day > 31:
		return false
	case month < 1 || month > 12:
		return false
	case year < 2020 || year > 2050:
		return false
	}

	return true
}
