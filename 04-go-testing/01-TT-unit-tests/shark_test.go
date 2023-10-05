package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ej 1
// Testear los siguientes casos negativos:
// El tiburón está cansado.
// El tiburón ya comió y está lleno.
// La presa es más rápida que el tiburón.
// Validar no solo que la función devuelve un error, sino también que el error sea el correspondiente.

// Ej 2
// Testear el caso feliz, donde el tiburón caza a la presa. Se debe validar que una vez que cazó a la presa, el tiburón se llenó y se cansó.

// Ej 3
// ¿Qué sucede si la presa es nula? En un caso ideal se debería devolver un error indicando el problema. Sin embargo, el programa posee un error de diseño. Realizar un test unitario para este caso, donde se debe validar que la función devuelve un error.
// ¿Qué sucede cuando corremos el test? Corregir el test para que pase, realizando las validaciones necesarias.
// Finalmente, volver al test anterior y corregir el código para que el test pase, devolviendo un error sin panic.

func TestSharkHuntsSuccessfully(t *testing.T) {
	t.Run("TestSharkHuntsSuccessfully", func(t *testing.T) {
		// Arrange
		prey := &Prey{name: "Fish", speed: 8}
		shark := &Shark{hungry: true, tired: false, speed: 10}
		expectedHunger := false

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.Nil(t, err, "Shark should be able to hunt")
		assert.Equal(t, expectedHunger, shark.hungry, "Shark shouldn't be hungry")
	})
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	t.Run("TestSharkCannotHuntBecauseIsTired", func(t *testing.T) {
		// Arrange
		prey := &Prey{name: "Fish", speed: 8}
		shark := &Shark{hungry: true, tired: true, speed: 10}
		expectedErrorStr := "cannot hunt, i am really tired"
		expectedHunger := true

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.NotNil(t, err, "Shark should not be able to hunt")
		assert.Equal(t, expectedErrorStr, err.Error(), "Hunt error reason is not the expected one")
		assert.Equal(t, expectedHunger, shark.hungry, "Shark should be hungry")
	})
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	t.Run("TestSharkCannotHuntBecaisIsNotHungry", func(t *testing.T) {
		// Arrange
		prey := &Prey{name: "Fish", speed: 8}
		shark := &Shark{hungry: false, tired: false, speed: 10}
		expectedErrorStr := "cannot hunt, i am not hungry"
		expectedHunger := false

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.NotNil(t, err, "Shark should not be able to hunt")
		assert.Equal(t, expectedErrorStr, err.Error(), "Hunt error reason is not the expected one")
		assert.Equal(t, expectedHunger, shark.hungry, "Shark shouldn't be hungry")
	})
}

func TestSharkCannotReachThePrey(t *testing.T) {
	t.Run("TestSharkCannotReachThePrey", func(t *testing.T) {
		// Arrange
		prey := &Prey{name: "Fish", speed: 12}
		shark := &Shark{hungry: true, tired: false, speed: 10}
		expectedErrorStr := "could not catch it"
		expectedHunger := true

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.NotNil(t, err, "Shark should not be able to hunt")
		assert.Equal(t, expectedErrorStr, err.Error(), "Hunt error reason is not the expected one")
		assert.Equal(t, expectedHunger, shark.hungry, "Shark should be hungry")
	})
}

func TestSharkHuntNilPrey(t *testing.T) {
	t.Run("TestSharkHuntNilPrey", func(t *testing.T) {
		// Arrange
		var prey *Prey
		shark := &Shark{hungry: true, tired: false, speed: 10}
		expectedErrorStr := "invalid prey"
		expectedHunger := true

		// Act
		err := shark.Hunt(prey)

		// Assert
		assert.NotNil(t, err, "Shark should not be able to hunt")
		assert.Equal(t, expectedErrorStr, err.Error(), "Hunt error reason is not the expected one")
		assert.Equal(t, expectedHunger, shark.hungry, "Shark should be hungry")
	})
}
