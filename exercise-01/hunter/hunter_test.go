package hunter_test

import (
	"exercise-01/hunter"
	"exercise-01/positioner"
	"exercise-01/prey"
	"exercise-01/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHunt(t *testing.T) {
	t.Run("Hunt: the hunter can catch", func(t *testing.T) {
		newPrey := prey.NewPreyStub()
		newPrey.FuncGetPosition = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		newPrey.FuncGetSpeed = func() (speed float64) {
			return 5.0
		}

		sm := simulator.NewSimulatorMock()
		sm.FuncCanCatch = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return true
		}

		newHunter := hunter.NewWhiteShark(10.0,
			&positioner.Position{X: 100, Y: 0, Z: 0}, sm)

		err := newHunter.Hunt(newPrey)

		expErr := error(nil)
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.Equal(t, sm.Call.CanCatch, expMockCallCanCatch)

	})
	t.Run("Hunt: the hunter can't catch", func(t *testing.T) {
		newPrey := prey.NewPreyStub()
		newPrey.FuncGetPosition = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		newPrey.FuncGetSpeed = func() (speed float64) {
			return 10
		}

		sm := simulator.NewSimulatorMock()
		sm.FuncCanCatch = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return false
		}

		newHunter := hunter.NewWhiteShark(1,
			&positioner.Position{X: 1, Y: 1, Z: 1}, sm)
		err := newHunter.Hunt(newPrey)
		expErr := hunter.ErrCanNotHunt
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.Equal(t, sm.Call.CanCatch, expMockCallCanCatch)

	})
}
