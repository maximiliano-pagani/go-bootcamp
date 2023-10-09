package shark

import (
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Continuando con el ejercicio anterior, seguiremos testeando el método Hunt del tiburón,
// aunque esta vez modificado. Para comenzar, prey y shark pasaron a ser interfaces que son
// implementadas por 2 structs: tuna y whiteShark respectivamente. Ahora ya no nos interesa
// si el tiburón tiene hambre o está cansado, si no que nos enfocaremos solo en la diferencia
// de velocidad. Para poder calcular si el tiburón es capaz de cazar a su presa, se utiliza
// un simulador que hace cálculos muy triviales basándose en la distancia entre el tiburón y
// su presa, y la velocidad de ambos, utilizando una variable que define el tiempo máximo de
// persecución. En el caso particular del atún, la velocidad se obtiene, aleatoriamente, entre 0
// y su velocidad máxima. Para calcular la distancia, además, provee una función que la calcula
// basándose en la posición inicial del tiburón (aleatoria).

// Ejercicio 1
// Si se fijan bien en el código, ya existe un test para el struct whiteShark con el que trabajaremos.
// Sin embargo, el test no está bien realizado. ¿Por qué?
// Borrar el test antes de continuar con el resto de los ejercicios.

// _______ El componente de aleatoriedad que altera el resultado de la ejecución imposibilita la
// _______ repetividad de los tests, como también del control de un escenario de ejecución particular.

// Ejercicio 2
// Crear stubs de prey para poder realizar los tests. Se deben poder cubrir todos los casos del
// método GetSpeed.

// Ejercicio 3
// Crear un mock para el simulator. El mock debe implementar simular la implementación del método
// CanCatch y un spy del método GetLinearDistance.

// Ejercicio 4
// Realizar test unitarios del método Hunt del tiburón blanco, cubriendo todos los casos posibles,
// usando los stubs y mocks creados anteriormente:
// 1. El tiburón logra cazar el atún al ser más veloz y al estar en una distancia corta. Hacer un
// assert de que el método GetLinearDistance fue llamado.
// 2. El tiburón no logra cazar el atún al ser más lento.
// 3. El tiburón no logra cazar el atún por estar a una distancia muy larga, a pesar de ser más veloz.

func TestSharkHuntsSuccessfully(t *testing.T) {
	t.Run("TestSharkHuntsSuccessfully", func(t *testing.T) {
		// Arrange
		prey := &prey.StubPrey{Speed: 4.0}
		simulator := &simulator.MockCatchSimulator{
			ValueCanCatch: true,
		}
		shark := CreateWhiteShark(simulator)
		expectedGetLinearDistanceCallsCount := 1

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.Nil(t, err, "Shark should be able to hunt successfully")
		assert.Equal(t, expectedGetLinearDistanceCallsCount, len(simulator.CallsGetLinearDistance), "GetLinearDistance should be called once")
	})
}

func TestSharkCannotHuntBecauseIsSlow(t *testing.T) {
	t.Run("TestSharkCannotHuntBecauseIsSlow", func(t *testing.T) {
		// Arrange
		prey := &prey.StubPrey{Speed: 200.0}
		simulator := &simulator.MockCatchSimulator{
			ValueCanCatch: false,
		}
		shark := CreateWhiteShark(simulator)
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
		prey := &prey.StubPrey{Speed: 4.0}
		simulator := &simulator.MockCatchSimulator{
			ValueCanCatch:          false,
			ValueGetLinearDistance: 1000.0,
		}
		shark := CreateWhiteShark(simulator)
		expectedErrorMsg := "could not hunt the prey"

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.NotNil(t, err, "Shark should fail to hunt")
		assert.Equal(t, expectedErrorMsg, err.Error())
	})
}
