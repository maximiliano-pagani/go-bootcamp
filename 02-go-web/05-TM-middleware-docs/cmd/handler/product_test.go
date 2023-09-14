package handler

import (
	"04-TT-functional-testing/internal/domain"
	"04-TT-functional-testing/internal/product"
	"04-TT-functional-testing/mock"
	response "04-TT-functional-testing/pkg/web"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var productsTestSample = []domain.Product{
	{
		Id:          1,
		Name:        "Oil - Margarine",
		Quantity:    439,
		Code:        "S82254D",
		IsPublished: true,
		Expiration:  "15/12/2021",
		Price:       71.42,
	},
	{
		Id:          2,
		Name:        "Pineapple - Canned, Rings",
		Quantity:    345,
		Code:        "M4637",
		IsPublished: true,
		Expiration:  "09/08/2021",
		Price:       352.79,
	},
	{
		Id:          3,
		Name:        "Wine - Red Oakridge Merlot",
		Quantity:    367,
		Code:        "T65812",
		IsPublished: false,
		Expiration:  "24/05/2021",
		Price:       179.23,
	},
	{
		Id:          4,
		Name:        "Cookie - Oatmeal",
		Quantity:    130,
		Code:        "M7157",
		IsPublished: false,
		Expiration:  "28/01/2022",
		Price:       275.47,
	},
	{
		Id:          5,
		Name:        "Flavouring Vanilla Artificial",
		Quantity:    336,
		Code:        "S60152S",
		IsPublished: true,
		Expiration:  "10/02/2022",
		Price:       839.02,
	},
}

func createTestServer(dbMock mock.ProductDBMock) *gin.Engine {
	token := "123456"

	repository := mock.NewProductRepositoryMock(dbMock)
	service := product.NewProductServiceDefault(repository)
	handler := NewProductHandlerDefault(service, token)

	testRouter := gin.Default()

	productsRouter := testRouter.Group("/products")

	productsRouter.GET("/", handler.GetAllProducts())
	productsRouter.POST("/", handler.NewProduct())
	productsRouter.GET("/:id", handler.GetProductById())
	productsRouter.PUT("/:id", handler.ReplaceProduct())
	productsRouter.PATCH("/:id", handler.UpdateProduct())
	productsRouter.DELETE("/:id", handler.DeleteProduct())
	productsRouter.GET("/search", handler.GetProductsByMinPrice())

	return testRouter
}

func createTestRequest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetAllProducts_OK(t *testing.T) {
	// Arrange
	expectedResponse, err := json.Marshal(
		response.SuccessResponse{
			Data: productsTestSample,
		},
	)
	assert.Nil(t, err)

	dataSample := make([]domain.Product, len(productsTestSample))
	copy(dataSample, productsTestSample)

	r := createTestServer(
		mock.ProductDBMock{
			Products: dataSample,
			LastId:   5,
			Error:    nil,
		},
	)

	req, respRec := createTestRequest(http.MethodGet, "/products/", "")

	// Act
	r.ServeHTTP(respRec, req)

	// Assert
	assert.Equal(t, http.StatusOK, respRec.Code)
	assert.Equal(t, expectedResponse, respRec.Body.Bytes())
}

func Test_GetProductById_OK(t *testing.T) {
	// Arrange
	expectedResponse, err := json.Marshal(
		response.SuccessResponse{
			Data: domain.Product{
				Id:          4,
				Name:        "Cookie - Oatmeal",
				Quantity:    130,
				Code:        "M7157",
				IsPublished: false,
				Expiration:  "28/01/2022",
				Price:       275.47,
			},
		},
	)
	assert.Nil(t, err)

	dataSample := make([]domain.Product, len(productsTestSample))
	copy(dataSample, productsTestSample)

	r := createTestServer(
		mock.ProductDBMock{
			Products: dataSample,
			LastId:   5,
			Error:    nil,
		},
	)
	req, respRec := createTestRequest(http.MethodGet, "/products/4", "")

	// Act
	r.ServeHTTP(respRec, req)

	// Assert
	assert.Equal(t, http.StatusOK, respRec.Code)
	assert.Equal(t, expectedResponse, respRec.Body.Bytes())
}

func Test_GetProductById_Error(t *testing.T) {
	t.Run("GetProductById_400_BadRequest", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status: http.StatusBadRequest,
				Code:   http.StatusText(http.StatusBadRequest),
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		r := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    nil,
			},
		)
		req, respRec := createTestRequest(http.MethodGet, "/products/a2d9m6", "")

		// Act
		r.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("GetProductById_404_NotFound", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status: http.StatusNotFound,
				Code:   http.StatusText(http.StatusNotFound),
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		r := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    errors.New(""),
			},
		)
		req, respRec := createTestRequest(http.MethodGet, "/products/2841", "")

		// Act
		r.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("GetProductById_401_Unauthorized", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Code:    http.StatusText(http.StatusUnauthorized),
				Message: "token inválido",
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		r := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    nil,
			},
		)

		req, respRec := createTestRequest(http.MethodGet, "/products/4", "")
		req.Header.Del("Token")

		// Act
		r.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})
}

func Test_NewProduct_OK(t *testing.T) {
	// Arrange
	expectedResponse, err := json.Marshal(
		response.SuccessResponse{
			Data: domain.Product{
				Id:          6,
				Name:        "New Test Product",
				Quantity:    100,
				Code:        "AAABBBCCC",
				IsPublished: false,
				Expiration:  "11/09/2023",
				Price:       1050.50,
			},
		},
	)
	assert.Nil(t, err)

	dataSample := make([]domain.Product, len(productsTestSample))
	copy(dataSample, productsTestSample)

	router := createTestServer(
		mock.ProductDBMock{
			Products: dataSample,
			LastId:   5,
			Error:    nil,
		},
	)

	req, respRec := createTestRequest(
		http.MethodPost,
		"/products/",
		`{
			"name": "New Test Product",
			"quantity": 100,
			"code_value": "AAABBBCCC",
			"is_published": false,
			"expiration": "11/09/2023",
			"price": 1050.50
		}`,
	)

	// Act
	router.ServeHTTP(respRec, req)

	// Assert
	assert.Equal(t, http.StatusCreated, respRec.Code)
	assert.Equal(t, expectedResponse, respRec.Body.Bytes())
}

func Test_DeleteProduct_OK(t *testing.T) {
	// Arrange
	dataSample := make([]domain.Product, len(productsTestSample))
	copy(dataSample, productsTestSample)

	router := createTestServer(
		mock.ProductDBMock{
			Products: dataSample,
			LastId:   5,
			Error:    nil,
		},
	)
	req, respRec := createTestRequest(http.MethodDelete, "/products/3", "")

	// Act
	router.ServeHTTP(respRec, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, respRec.Code)
}

func Test_DeleteProduct_Error(t *testing.T) {
	t.Run("DeleteProduct_400_BadRequest", func(t *testing.T) {
		// Arrange
		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    nil,
			},
		)
		req, respRec := createTestRequest(http.MethodDelete, "/products/a2d9m6", "")

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, respRec.Code)
	})

	t.Run("DeleteProduct_401_Unauthorized", func(t *testing.T) {
		// Arrange
		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    nil,
			},
		)

		req, respRec := createTestRequest(http.MethodDelete, "/products/4", "")
		req.Header.Del("Token")

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, respRec.Code)
	})

	t.Run("DeleteProduct_404_NotFound", func(t *testing.T) {
		// Arrange
		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    errors.New(""),
			},
		)
		req, respRec := createTestRequest(http.MethodDelete, "/products/2841", "")

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, respRec.Code)
	})
}

func Test_ReplaceProduct_Error(t *testing.T) {
	t.Run("ReplaceProduct_400_BadRequest", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status: http.StatusBadRequest,
				Code:   http.StatusText(http.StatusBadRequest),
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    nil,
			},
		)

		req, respRec := createTestRequest(
			http.MethodPut,
			"/products/3",
			`{
				"name": "Update Test Product",
				"quantity": "BADQUANTITY",
				"code_value": "we98fefn9f",
				"is_published": false,
				"expiration": "11/09/2023",
				"price": 1050.50
			}`,
		)

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("ReplaceProduct_404_NotFound", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status: http.StatusNotFound,
				Code:   http.StatusText(http.StatusNotFound),
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    errors.New(""),
			},
		)

		req, respRec := createTestRequest(
			http.MethodPut,
			"/products/4934",
			`{
				"name": "Update Test Product",
				"quantity": 123,
				"code_value": "we98fefn9f",
				"is_published": false,
				"expiration": "11/09/2023",
				"price": 1050.50
			}`,
		)

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("ReplaceProduct_401_Unauthorized", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Code:    http.StatusText(http.StatusUnauthorized),
				Message: "token inválido",
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    nil,
			},
		)

		req, respRec := createTestRequest(
			http.MethodPut,
			"/products/3",
			`{
				"name": "Update Test Product",
				"quantity": 123,
				"code_value": "we98fefn9f",
				"is_published": false,
				"expiration": "11/09/2023",
				"price": 1050.50
			}`,
		)
		req.Header.Del("Token")

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})
}

func Test_UpdateProduct_Error(t *testing.T) {
	t.Run("UpdateProduct_400_BadRequest", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status: http.StatusBadRequest,
				Code:   http.StatusText(http.StatusBadRequest),
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    nil,
			},
		)

		req, respRec := createTestRequest(
			http.MethodPatch,
			"/products/3",
			`{
				"price": "$$$1349.0"
			}`,
		)

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("UpdateProduct_404_NotFound", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status: http.StatusNotFound,
				Code:   http.StatusText(http.StatusNotFound),
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    errors.New(""),
			},
		)

		req, respRec := createTestRequest(
			http.MethodPatch,
			"/products/4934",
			`{
				"price": 1349.0
			}`,
		)

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("UpdateProduct_401_Unauthorized", func(t *testing.T) {
		// Arrange
		expectedResponse, err := json.Marshal(
			response.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Code:    http.StatusText(http.StatusUnauthorized),
				Message: "token inválido",
			},
		)
		assert.Nil(t, err)

		dataSample := make([]domain.Product, len(productsTestSample))
		copy(dataSample, productsTestSample)

		router := createTestServer(
			mock.ProductDBMock{
				Products: dataSample,
				LastId:   5,
				Error:    nil,
			},
		)

		req, respRec := createTestRequest(
			http.MethodPatch,
			"/products/3",
			`{
				"price": 1349.0
			}`,
		)
		req.Header.Del("Token")

		// Act
		router.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})
}
