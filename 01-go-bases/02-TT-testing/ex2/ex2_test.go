package ex2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Calcular el promedio de las notas de los alumnos.

func TestGetAverage(t *testing.T) {
	t.Run("Average of grades", func(t *testing.T) {
		// Arrange
		grades := []float32{1.0, 4.0, 5.0, 6.0, 9.0}
		expectedAverage := 5.0

		// Act
		result := getAverage(grades...)

		// Assert
		assert.Equal(t, int(expectedAverage), int(result), "Average is not correct, should be 5.0")
	})
}
