package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ejercicio 1
// Siguiendo con la metodología TDD (testear los siguientes casos antes del código) en el paquete tests :
// 1. Se configura la presa para que sea más veloz que el tiburón, y al simular la caza logra escaparse.
// 2. Se configura el tiburón más rápido que la presa, pero se encuentra demasiado lejos y no logra cazarla.
// 3. El tiburón y la presa se configuran de modo que el tiburón logra cazarla luego de 24 segundos (tener en
// cuenta el algoritmo que usa el simulador).
// 4. Testear, para todos los endpoints de configuración, casos donde los tipos de los campos no son esperados.
// Tener en cuenta los structs creados en el archivo handlers.go. Es importante que utilicemos los métodos
// creados en el archivo utils.go del paquete tests.

// Ejercicio 2
// Completar los métodos del handler hasta que los tests pasen.

func TestSharkHuntsSuccessfullyAfter24Secs(t *testing.T) {
	server := createServer()

	t.Run("TestSharkConfigureSpeed", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(
			http.MethodPut,
			"/v1/shark",
			`{ "x_position": 120.0, "y_position": 0.0, "speed": 15.0 }`,
		)
		expectedCode := http.StatusOK
		expectedResponse, err := json.Marshal(
			struct {
				Success bool `json:"success"`
			}{Success: true},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("TestPreyConfigureSpeed", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(
			http.MethodPut,
			"/v1/prey",
			`{ "speed": 10.0 }`,
		)
		expectedCode := http.StatusOK
		expectedResponse, err := json.Marshal(
			struct {
				Success bool `json:"success"`
			}{Success: true},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("TestSharkHuntsSuccessfullyAfter24Secs", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(http.MethodPost, "/v1/simulate", "")
		expectedCode := http.StatusOK
		expectedResponse, err := json.Marshal(
			struct {
				Success bool    `json:"success"`
				Message string  `json:"message"`
				Time    float64 `json:"time"`
			}{
				Success: true,
				Message: "hunted successfully",
				Time:    24.0,
			},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})
}

func TestSharkCannotHuntBecauseIsSlow(t *testing.T) {
	server := createServer()

	t.Run("TestSharkConfigureSpeed", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(
			http.MethodPut,
			"/v1/shark",
			`{ "x_position": 20.0, "y_position": 20.0, "speed": 10.0 }`,
		)
		expectedCode := http.StatusOK
		expectedResponse, err := json.Marshal(
			struct {
				Success bool `json:"success"`
			}{Success: true},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("TestPreyConfigureSpeed", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(
			http.MethodPut,
			"/v1/prey",
			`{ "speed": 15.0 }`,
		)
		expectedCode := http.StatusOK
		expectedResponse, err := json.Marshal(
			struct {
				Success bool `json:"success"`
			}{Success: true},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("TestSharkCannotHuntBecauseIsSlow", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(http.MethodPost, "/v1/simulate", "")
		expectedCode := http.StatusConflict
		expectedResponse, err := json.Marshal(
			struct {
				Success bool    `json:"success"`
				Message string  `json:"message"`
				Time    float64 `json:"time"`
			}{
				Success: false,
				Message: "could not catch it",
				Time:    0,
			},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})
}

func TestSharkCannotHuntBecauseIsTooFar(t *testing.T) {
	server := createServer()

	t.Run("TestSharkConfigureSpeed", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(
			http.MethodPut,
			"/v1/shark",
			`{ "x_position": 1000.0, "y_position": 1000.0, "speed": 15.0 }`,
		)
		expectedCode := http.StatusOK
		expectedResponse, err := json.Marshal(
			struct {
				Success bool `json:"success"`
			}{Success: true},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("TestPreyConfigureSpeed", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(
			http.MethodPut,
			"/v1/prey",
			`{ "speed": 10.0 }`,
		)
		expectedCode := http.StatusOK
		expectedResponse, err := json.Marshal(
			struct {
				Success bool `json:"success"`
			}{Success: true},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("TestSharkCannotHuntBecauseIsTooFar", func(t *testing.T) {
		// Arrange
		req, respRec := createRequestTest(http.MethodPost, "/v1/simulate", "")
		expectedCode := http.StatusConflict
		expectedResponse, err := json.Marshal(
			struct {
				Success bool    `json:"success"`
				Message string  `json:"message"`
				Time    float64 `json:"time"`
			}{
				Success: false,
				Message: "could not catch it",
				Time:    0,
			},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})
}

func TestConfigureBadRequest(t *testing.T) {
	t.Run("TestSharkConfigureSpeedBadRequest", func(t *testing.T) {
		// Arrange
		server := createServer()
		req, respRec := createRequestTest(
			http.MethodPut,
			"/v1/shark",
			`{ "x_position": "1000.0", "y_position": 1000.0, "speed": 15.0 }`,
		)
		expectedCode := http.StatusUnprocessableEntity
		expectedResponse, err := json.Marshal(
			struct {
				Success bool `json:"success"`
			}{Success: false},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})

	t.Run("TestPreyConfigureSpeedBadRequest", func(t *testing.T) {
		// Arrange
		server := createServer()
		req, respRec := createRequestTest(
			http.MethodPut,
			"/v1/prey",
			`{ speed: 10..0 }`,
		)
		expectedCode := http.StatusUnprocessableEntity
		expectedResponse, err := json.Marshal(
			struct {
				Success bool `json:"success"`
			}{Success: false},
		)
		assert.Nil(t, err)

		// Act
		server.ServeHTTP(respRec, req)

		// Assert
		assert.Equal(t, expectedCode, respRec.Code)
		assert.Equal(t, expectedResponse, respRec.Body.Bytes())
	})
}
