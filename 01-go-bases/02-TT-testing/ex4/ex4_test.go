package ex4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Realizar test para calcular el mínimo de calificaciones.
// Realizar test para calcular el máximo de calificaciones.
// Realizar test para calcular el promedio de calificaciones.

func TestMinOp(t *testing.T) {
	t.Run("Min operation returns minimum value", func(t *testing.T) {
		// Arrange
		minFunc, _ := operation(minimum)
		grades := []int{2, 4, 5, 6, 8}
		expectedMin := 2.0

		// Act
		result := minFunc(grades...)

		// Assert
		assert.Equal(t, int(expectedMin), int(result), "Min value should be 2")
	})
}

func TestAvgOp(t *testing.T) {
	t.Run("Avg operation returns average value", func(t *testing.T) {
		// Arrange
		averageFunc, _ := operation(average)
		grades := []int{2, 4, 5, 6, 8}
		expectedAvg := 5.0

		// Act
		result := averageFunc(grades...)

		// Assert
		assert.Equal(t, int(expectedAvg), int(result), "Avg value should be 5")
	})
}

func TestMaxOp(t *testing.T) {
	t.Run("Max operation returns maximum value", func(t *testing.T) {
		// Arrange
		maxFunc, _ := operation(maximum)
		grades := []int{2, 4, 5, 6, 8}
		expectedMax := 8.0

		// Act
		result := maxFunc(grades...)

		// Assert
		assert.Equal(t, int(expectedMax), int(result), "Max value should be 8")
	})
}
