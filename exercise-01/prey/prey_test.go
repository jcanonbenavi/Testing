package prey_test

import (
	"exercise-01/positioner"
	"exercise-01/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuna_CreateTuna(t *testing.T) {
	t.Run("Success: Create tuna with randon parameters", func(t *testing.T) {
		// arrange
		tuna := prey.CreateTuna()
		output := tuna.GetSpeed()
		// act
		require.NotEmpty(t, output)

	})

}

func TestTuna_GetSpeed(t *testing.T) {
	t.Run("Speed is greater than 0.0", func(t *testing.T) {
		// arrange
		tuna := prey.NewTuna(10.0, nil)
		output := tuna.GetSpeed()
		// act

		expectedValue := 10.0
		require.Equal(t, output, expectedValue)

	})
	t.Run("Speed is 0.0", func(t *testing.T) {
		// arrange
		tuna := prey.NewTuna(0.0, nil)
		output := tuna.GetSpeed()
		// act

		expectedValue := 0.0
		require.Equal(t, output, expectedValue)

	})

	t.Run("Get Speed: negative parameters", func(t *testing.T) {
		// arrange
		velocity := -10.0
		position := &positioner.Position{X: 0, Y: 0, Z: 0}
		tuna := prey.NewTuna(velocity, position)
		tuna.GetSpeed()
		// act

		expectedValue := -10.0
		require.Equal(t, tuna.GetSpeed(), expectedValue)

	})

}

func TestTuna_GetPosition(t *testing.T) {
	t.Run("Get position: positive parameters", func(t *testing.T) {
		// arrange
		position := &positioner.Position{X: 10, Y: 10, Z: 10}
		tuna := prey.NewTuna(0.0, position)
		output := tuna.GetPosition()
		// act

		expectedValue := &positioner.Position{X: 10, Y: 10, Z: 10}
		require.Equal(t, output, expectedValue)

	})
	t.Run("Get position is nil", func(t *testing.T) {
		// arrange
		impl := prey.NewTuna(0.0, nil)

		// act
		output := impl.GetPosition()

		// assert
		require.Nil(t, output)
	})

}
