package ex3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Calcular el salario de la categoría “A”.
// Calcular el salario de la categoría “B”.
// Calcular el salario de la categoría “C”.

func TestGetCategorySalary(t *testing.T) {
	t.Run("Get category A salary", func(t *testing.T) {
		// Arrange
		minutes := 60
		cat := catA
		expectedSalary := 4500.0

		// Act
		result := getCategorySalary(minutes, cat)

		// Assert
		assert.Equal(t, int(expectedSalary), int(result), "Salary of category A should be 4500")
	})

	t.Run("Get category B salary", func(t *testing.T) {
		// Arrange
		minutes := 60
		cat := catB
		expectedSalary := 1800.0

		// Act
		result := getCategorySalary(minutes, cat)

		// Assert
		assert.Equal(t, int(expectedSalary), int(result), "Salary of category B should be 1800")
	})

	t.Run("Get category C salary", func(t *testing.T) {
		// Arrange
		minutes := 60
		cat := catC
		expectedSalary := 1000.0

		// Act
		result := getCategorySalary(minutes, cat)

		// Assert
		assert.Equal(t, int(expectedSalary), int(result), "Salary of category C should be 1000")
	})
}
