package hunt_test

import (
	hunt "exercise-00"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {

		hunter := hunt.NewWhiteShark(true, false, 10)
		tuna := hunt.NewTuna("Tuna", 5)
		err := hunter.Hunt(tuna)
		require.NoError(t, err)
		require.False(t, hunter.Hungry)
		require.True(t, hunter.Tired)

	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		hunter := hunt.NewWhiteShark(false, false, 10)
		tuna := hunt.NewTuna("Tuna", 5)
		err := hunter.Hunt(tuna)
		require.Error(t, err)
		require.Equal(t, hunt.ErrSharkIsNotHungry, err)

	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		hunter := hunt.NewWhiteShark(true, true, 10)
		tuna := hunt.NewTuna("Tuna", 5)
		err := hunter.Hunt(tuna)
		require.Error(t, err)
		require.Equal(t, hunt.ErrSharkIsTired, err)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		hunter := hunt.NewWhiteShark(true, false, 5)
		tuna := hunt.NewTuna("Tuna", 10)
		err := hunter.Hunt(tuna)
		require.Error(t, err)
		require.Equal(t, hunt.ErrSharkIsSlower, err)

	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		hunter := hunt.NewWhiteShark(true, false, 10)
		var tuna *hunt.Tuna
		err := hunter.Hunt(tuna)
		require.Error(t, hunt.ErrTunaIsNil, err)
		require.Equal(t, hunt.ErrTunaIsNil, err)
	})
}

func NewWhiteShark(true, false bool, i int) {
	panic("unimplemented")
}
