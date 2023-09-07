package controller

import (
	"02-TM-post/internal/repository"
	"02-TM-post/internal/service"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestNewProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type DataNewProduct struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseNewProduct struct {
	Code int             `json:"code"`
	Msg  string          `json:"message"`
	Data *DataNewProduct `json:"data"`
}

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, service.GetAllProducts())
		return
	}
}

func GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func GetProductsByMinPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func NewProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody := &RequestNewProduct{}
		err := c.ShouldBindJSON(reqBody)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = validateRequestNewProduct(reqBody)

		if err != nil {
			code := http.StatusBadRequest
			respBody := &ResponseNewProduct{Code: code, Msg: err.Error()}
			c.AbortWithStatusJSON(code, respBody)
			return
		}

		addedProduct, err := service.NewProduct(
			&repository.Product{
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
			respBody := &ResponseNewProduct{Code: code, Msg: err.Error()}
			c.AbortWithStatusJSON(code, respBody)
			return
		}

		code := http.StatusCreated
		respBody := &ResponseNewProduct{Code: code, Msg: "Created",
			Data: &DataNewProduct{
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

func validateRequestNewProduct(product *RequestNewProduct) error {
	switch {
	case product.Name == "":
		return errors.New("Invalid Name")
	case product.Code == "":
		return errors.New("Invalid Code")
	case product.Quantity == 0:
		return errors.New("Invalid Quantity")
	case product.Expiration == "" || !isValidExpirationDate(product.Expiration):
		return errors.New("Invalid Expiration Date")
	case product.Price == 0:
		return errors.New("Invalid Price")
	}
	return nil
}

func isValidExpirationDate(date string) bool {
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
