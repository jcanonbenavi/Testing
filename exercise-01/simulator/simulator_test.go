package simulator_test

import (
	"exercise-01/positioner"
	"exercise-01/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCatch(t *testing.T) {
	t.Run("Can catch: the hunter can catch", func(t *testing.T) {
		ps := positioner.NewPositionerStub()
		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			return 10.0
		}
		canCatch := simulator.NewCatchSimulatorDefault(10.0, ps)
		hunter := &simulator.Subject{Position: &positioner.Position{X: 100, Y: 0, Z: 0}, Speed: 10}
		prey := &simulator.Subject{Position: &positioner.Position{X: 0, Y: 0, Z: 0}, Speed: 5}
		output := canCatch.CanCatch(hunter, prey)
		expectedValue := true
		require.Equal(t, output, expectedValue)
	})

	t.Run("Can't catch: the hunter can't catch", func(t *testing.T) {
		ps := positioner.NewPositionerStub()
		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			return 10.0
		}
		canCatch := simulator.NewCatchSimulatorDefault(10.0, ps)
		hunter := &simulator.Subject{Position: &positioner.Position{X: 100, Y: 0, Z: 0}, Speed: 5}
		prey := &simulator.Subject{Position: &positioner.Position{X: 0, Y: 0, Z: 0}, Speed: 15}
		output := canCatch.CanCatch(hunter, prey)
		expectedValue := false
		require.Equal(t, output, expectedValue)
	})

	t.Run("Can't catch: the hunter don't have time", func(t *testing.T) {
		ps := positioner.NewPositionerStub()
		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			return 10.0
		}
		canCatch := simulator.NewCatchSimulatorDefault(10.0, ps)
		hunter := &simulator.Subject{Position: &positioner.Position{X: 0, Y: 0, Z: 0}, Speed: 5}
		prey := &simulator.Subject{Position: &positioner.Position{X: 100, Y: 0, Z: 0}, Speed: 10}
		output := canCatch.CanCatch(hunter, prey)
		expectedValue := false
		require.Equal(t, output, expectedValue)
	})

}
