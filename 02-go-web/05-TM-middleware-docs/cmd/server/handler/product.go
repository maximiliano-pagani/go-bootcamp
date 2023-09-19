package handler

import (
	"05-TM-middleware-docs/internal/domain"
	"05-TM-middleware-docs/internal/product"
	response "05-TM-middleware-docs/pkg/web"
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

type ResponseProduct struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// GetAllProducts godoc
// @Summary      List all products
// @Description  Searches and returns all products currently listed
// @Tags         products
// @Produce      json
// @Param        token header string true "Token"
// @Success      200 {object}  response.SuccessResponse
// @Failure      500 {object}  response.ErrorResponse
// @Router       /products [get]
func (h *ProductHandlerDefault) GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := h.service.GetAllProducts()

		if err != nil {
			response.Failure(c, http.StatusInternalServerError, nil)
			return
		}

		response.Success(c, http.StatusOK, results)
		return
	}
}

// GetProductById godoc
// @Summary      List a specific product
// @Description  Searches and returns a product that matches the id provided
// @Tags         products
// @Produce      json
// @Param        id path int true "Id"
// @Param        token header string true "Token"
// @Success      200 {object}  response.SuccessResponse
// @Failure      400 {object}  response.ErrorResponse
// @Failure      404 {object}  response.ErrorResponse
// @Router       /products/{id} [get]
func (h *ProductHandlerDefault) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
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

// GetProductsByMinPrice godoc
// @Summary      List all products with min price
// @Description  Searches and returns all products with price greater than min provided
// @Tags         products
// @Produce      json
// @Param        priceGt query float64 true "Price greater than"
// @Param        token header string true "Token"
// @Success      200 {object}  response.SuccessResponse
// @Failure      400 {object}  response.ErrorResponse
// @Failure      500 {object}  response.ErrorResponse
// @Router       /products/search [get]
func (h *ProductHandlerDefault) GetProductsByMinPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
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

// NewProduct godoc
// @Summary      New product
// @Description  Creates a new product with the values provided in request body
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        token header string true "Token"
// @Param        body body handler.RequestProduct true "Product"
// @Success      201 {object}  response.SuccessResponse
// @Failure      400 {object}  response.ErrorResponse
// @Router       /products [post]
func (h *ProductHandlerDefault) NewProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
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

// ReplaceProduct godoc
// @Summary      Replace product
// @Description  Completely replaces the product id specified with the values in the request body
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path int true "Id"
// @Param        token header string true "Token"
// @Param        body body handler.RequestProduct true "Product"
// @Success      200 {object}  response.SuccessResponse
// @Failure      400 {object}  response.ErrorResponse
// @Failure      404 {object}  response.ErrorResponse
// @Router       /products/{id} [put]
func (h *ProductHandlerDefault) ReplaceProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
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

// UpdateProduct godoc
// @Summary      Update product
// @Description  Partially updates the product id specified with the values in the request body
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path int true "Id"
// @Param        token header string true "Token"
// @Param        body body handler.RequestProduct true "Product"
// @Success      201 {object}  response.SuccessResponse
// @Failure      400 {object}  response.ErrorResponse
// @Failure      404 {object}  response.ErrorResponse
// @Router       /products/{id} [patch]
func (h *ProductHandlerDefault) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
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

// DeleteProduct godoc
// @Summary      Delete product
// @Description  Deletes a product with the specified id, if exists
// @Tags         products
// @Param        id path int true "Id"
// @Param        token header string true "Token"
// @Success      204 {object}  response.SuccessResponse
// @Failure      400 {object}  response.ErrorResponse
// @Failure      404 {object}  response.ErrorResponse
// @Router       /products/{id} [delete]
func (h *ProductHandlerDefault) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
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
