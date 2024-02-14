package positioner_test

import (
	"exercise-01/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinearDistance(t *testing.T) {
	t.Run("All negative coordinates", func(t *testing.T) {
		from := &positioner.Position{X: -1, Y: -1, Z: -1}
		to := &positioner.Position{X: -1, Y: -1, Z: -1}

		impl := positioner.NewPositionerDefault()

		// act
		linearDistance := impl.GetLinearDistance(from, to)
		expectedValue := 0.0
		require.Equal(t, linearDistance, expectedValue)

	})

	t.Run("All positive coordinates", func(t *testing.T) {
		from := &positioner.Position{X: 1, Y: 2, Z: 3}
		to := &positioner.Position{X: 6, Y: 5, Z: 4}

		impl := positioner.NewPositionerDefault()

		// act
		linearDistance := impl.GetLinearDistance(from, to)
		expectedValue := 5.916079783099616
		require.Equal(t, linearDistance, expectedValue)

	})

	t.Run("Radicand is a perfect square", func(t *testing.T) {
		from := &positioner.Position{X: 0, Y: 0, Z: 3}
		to := &positioner.Position{X: 0, Y: 0, Z: 6}

		impl := positioner.NewPositionerDefault()

		// act
		linearDistance := impl.GetLinearDistance(from, to)
		expectedValue := 3.0
		require.Equal(t, linearDistance, expectedValue)

	})

}
