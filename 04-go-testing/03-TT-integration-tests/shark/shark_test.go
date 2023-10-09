package shark

import (
	"integrationtests/pkg/storage"
	"integrationtests/prey"
	"integrationtests/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Continuando con el ejercicio anterior, seguiremos testeando el método Hunt del tiburón,
// aunque nuevamente nuestro sistema fue modificado. Aunque la aleatoriedad de nuestro sistema
// anterior nos permitió hacer uso de los tests dobles, no deseamos que nuestro sistema sea
// tan impredecible. Es por esto que, ahora, los datos relacionados con la velocidad y posición,
// tanto de la presa como del tiburón, se guardan en un archivo de configuración. Para esto
// creamos un paquete que nos permite manejar este archivo y obtener los valores que necesitamos.

// Ejercicio 1
// Realizar un mock del paquete recientemente creado storage.

// Ejercicio 2
// Realizar tests de integración que cubran los siguientes casos, para esto se debe usar únicamente
// el mock del paquete storage:
// 1. El tiburón logra cazar el atún al ser más veloz y al estar en una distancia corta. Hacer un
// assert de que el método GetLinearDistance fue llamado.
// 2. El tiburón no logra cazar el atún al ser más lento.
// 3. El tiburón no logra cazar el atún por estar a una distancia muy larga, a pesar de ser más veloz.

func TestSharkHuntsSuccessfully(t *testing.T) {
	t.Run("TestSharkHuntsSuccessfully", func(t *testing.T) {
		// Arrange
		stubConfig := map[string]float64{
			"tuna_speed":        6.0,
			"white_shark_speed": 8.0,
			"white_shark_x":     2.0,
			"white_shark_y":     4.0,
		}
		mockStorage := storage.NewMockStorage(stubConfig)
		prey := prey.CreateTuna(mockStorage)
		simulator := simulator.NewCatchSimulator(5.0)
		shark := CreateWhiteShark(simulator, mockStorage)

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.Nil(t, err, "Shark should be able to hunt successfully")
	})
}

func TestSharkCannotHuntBecauseIsSlow(t *testing.T) {
	t.Run("TestSharkCannotHuntBecauseIsSlow", func(t *testing.T) {
		// Arrange
		stubConfig := map[string]float64{
			"tuna_speed":        6.0,
			"white_shark_speed": 5.0,
			"white_shark_x":     2.0,
			"white_shark_y":     4.0,
		}
		mockStorage := storage.NewMockStorage(stubConfig)
		prey := prey.CreateTuna(mockStorage)
		simulator := simulator.NewCatchSimulator(5.0)
		shark := CreateWhiteShark(simulator, mockStorage)
		expectedErrorMsg := "could not hunt the prey"

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.NotNil(t, err, "Shark should fail to hunt")
		assert.Equal(t, expectedErrorMsg, err.Error())
	})
}

func TestSharkCannotHuntBecaisIsNotNear(t *testing.T) {
	t.Run("TestSharkCannotHuntBecaisIsNotNear", func(t *testing.T) {
		// Arrange
		stubConfig := map[string]float64{
			"tuna_speed":        6.0,
			"white_shark_speed": 8.0,
			"white_shark_x":     17.0,
			"white_shark_y":     16.0,
		}
		mockStorage := storage.NewMockStorage(stubConfig)
		prey := prey.CreateTuna(mockStorage)
		simulator := simulator.NewCatchSimulator(5.0)
		shark := CreateWhiteShark(simulator, mockStorage)
		expectedErrorMsg := "could not hunt the prey"

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.NotNil(t, err, "Shark should fail to hunt")
		assert.Equal(t, expectedErrorMsg, err.Error())
	})
}
