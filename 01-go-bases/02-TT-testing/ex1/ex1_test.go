package ex1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
// Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
// Calcular el impuesto en caso de que el empleado gane por encima de $150.000.

func TestGetSalaryTaxes(t *testing.T) {
	t.Run("Tax is equal to 0.0%", func(t *testing.T) {
		// Arrange
		salary := 49999
		expectedTax := 0.0

		// Act
		result := getSalaryTaxes(salary)

		// Assert
		assert.Equal(t, int(expectedTax), int(result), "Tax is not equal to 0.0")
	})

	t.Run("Tax is equal to 0.17%", func(t *testing.T) {
		// Arrange
		salary := 50001
		expectedTax := 50001 * 0.17

		// Act
		result := getSalaryTaxes(salary)

		// Assert
		assert.Equal(t, int(expectedTax), int(result), "Tax is not equal to 0.17")
	})

	t.Run("Tax is equal to 0.27%", func(t *testing.T) {
		// Arrange
		salary := 150001
		expectedTax := 150001 * 0.27

		// Act
		result := getSalaryTaxes(salary)

		// Assert
		assert.Equal(t, int(expectedTax), int(result), "Tax is not equal to 0.27")
	})
}
