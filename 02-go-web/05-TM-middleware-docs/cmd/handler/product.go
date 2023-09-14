package handler

import (
	"04-TT-functional-testing/internal/domain"
	"04-TT-functional-testing/internal/product"
	response "04-TT-functional-testing/pkg/web"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductHandlerDefault struct {
	service product.ProductService
	token   string
}

func NewProductHandlerDefault(service product.ProductService, token string) *ProductHandlerDefault {
	handler := &ProductHandlerDefault{service: service, token: token}
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

type ResponseProduct struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (h *ProductHandlerDefault) GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("TOKEN") != h.token {
			response.Failure(c, http.StatusUnauthorized, errors.New("token inválido"))
			return
		}

		results, err := h.service.GetAllProducts()

		if err != nil {
			response.Failure(c, http.StatusInternalServerError, nil)
			return
		}

		response.Success(c, http.StatusOK, results)
		return
	}
}

func (h *ProductHandlerDefault) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("TOKEN") != h.token {
			response.Failure(c, http.StatusUnauthorized, errors.New("token inválido"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		result, err := h.service.GetProductById(id)

		if err != nil {
			response.Failure(c, http.StatusNotFound, nil)
			return
		}

		response.Success(c, http.StatusOK, &result)
		return
	}
}

func (h *ProductHandlerDefault) GetProductsByMinPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("TOKEN") != h.token {
			response.Failure(c, http.StatusUnauthorized, errors.New("token inválido"))
			return
		}

		param := c.Query("priceGt")

		if param == "" {
			if results, err := h.service.GetAllProducts(); err != nil {
				response.Failure(c, http.StatusInternalServerError, nil)
			} else {
				response.Success(c, http.StatusOK, results)
			}

			return
		}

		minPrice, err := strconv.ParseFloat(param, 64)

		if err != nil {
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		results, err := h.service.GetProductsByMinPrice(minPrice)

		if err != nil {
			response.Failure(c, http.StatusInternalServerError, nil)
			return
		}

		response.Success(c, http.StatusOK, results)
		return
	}
}

func (h *ProductHandlerDefault) NewProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("TOKEN") != h.token {
			response.Failure(c, http.StatusUnauthorized, errors.New("token inválido"))
			return
		}

		reqBody := &RequestProduct{}
		err := c.ShouldBindJSON(reqBody)

		if err != nil {
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		err = h.validateRequestProduct(reqBody)

		if err != nil {
			response.Failure(c, http.StatusBadRequest, err)
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
			response.Failure(c, http.StatusBadRequest, err)
			return
		}

		respBody := ResponseProduct{
			Id:          addedProduct.Id,
			Name:        addedProduct.Name,
			Quantity:    addedProduct.Quantity,
			Code:        addedProduct.Code,
			IsPublished: addedProduct.IsPublished,
			Expiration:  addedProduct.Expiration,
			Price:       addedProduct.Price,
		}

		response.Success(c, http.StatusCreated, &respBody)
		return
	}
}

func (h *ProductHandlerDefault) ReplaceProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("TOKEN") != h.token {
			response.Failure(c, http.StatusUnauthorized, errors.New("token inválido"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		reqBody := &RequestProduct{}
		err = c.ShouldBindJSON(reqBody)

		if err != nil {
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		err = h.validateRequestProduct(reqBody)

		if err != nil {
			response.Failure(c, http.StatusBadRequest, err)
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
			response.Failure(c, http.StatusNotFound, nil)
			return
		}

		respBody := ResponseProduct{
			Id:          replacedProduct.Id,
			Name:        replacedProduct.Name,
			Quantity:    replacedProduct.Quantity,
			Code:        replacedProduct.Code,
			IsPublished: replacedProduct.IsPublished,
			Expiration:  replacedProduct.Expiration,
			Price:       replacedProduct.Price,
		}

		response.Success(c, http.StatusOK, &respBody)
		return
	}
}

func (h *ProductHandlerDefault) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("TOKEN") != h.token {
			response.Failure(c, http.StatusUnauthorized, errors.New("token inválido"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		originalProduct, err := h.service.GetProductById(id)

		if err != nil {
			response.Failure(c, http.StatusNotFound, nil)
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
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		err = h.validateRequestProduct(reqBody)

		if err != nil {
			response.Failure(c, http.StatusBadRequest, err)
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
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		respBody := ResponseProduct{
			Id:          updatedProduct.Id,
			Name:        updatedProduct.Name,
			Quantity:    updatedProduct.Quantity,
			Code:        updatedProduct.Code,
			IsPublished: updatedProduct.IsPublished,
			Expiration:  updatedProduct.Expiration,
			Price:       updatedProduct.Price,
		}

		response.Success(c, http.StatusOK, &respBody)
		return
	}
}

func (h *ProductHandlerDefault) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("TOKEN") != h.token {
			response.Failure(c, http.StatusUnauthorized, errors.New("token inválido"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			response.Failure(c, http.StatusBadRequest, nil)
			return
		}

		err = h.service.DeleteProduct(id)

		if err != nil {
			response.Failure(c, http.StatusNotFound, nil)
			return
		}

		response.Success(c, http.StatusNoContent, nil)
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
