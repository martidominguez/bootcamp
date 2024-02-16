package positioner_test

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinearDistance(t *testing.T) {
	t.Run("get linear distance - coordinates are negative", func(t *testing.T) {
		// arrange
		newPositioner := positioner.NewPositionerDefault()
		from := &positioner.Position{
			X: -1,
			Y: -1,
			Z: -1,
		}
		to := &positioner.Position{
			X: -2,
			Y: -2,
			Z: -2,
		}

		// act
		linearDistance := newPositioner.GetLinearDistance(from, to)

		// assert
		expectedDistance := 1.7320508075688772
		require.Equal(t, expectedDistance, linearDistance)
	})
	t.Run("get linear distance - coordinates are positive", func(t *testing.T) {
		// arrange
		newPositioner := positioner.NewPositionerDefault()
		from := &positioner.Position{
			X: 1,
			Y: 1,
			Z: 1,
		}
		to := &positioner.Position{
			X: 2,
			Y: 2,
			Z: 2,
		}

		// act
		linearDistance := newPositioner.GetLinearDistance(from, to)

		// assert
		expectedDistance := 1.7320508075688772
		require.Equal(t, expectedDistance, linearDistance)
	})
	t.Run("get linear distance - coordinates return a linear distance without decimals", func(t *testing.T) {
		// arrange
		newPositioner := positioner.NewPositionerDefault()
		from := &positioner.Position{
			X: 1,
			Y: 0,
			Z: 0,
		}
		to := &positioner.Position{
			X: 2,
			Y: 0,
			Z: 0,
		}

		// act
		linearDistance := newPositioner.GetLinearDistance(from, to)

		// assert
		expectedDistance := 1.0
		require.Equal(t, expectedDistance, linearDistance)
	})
}
