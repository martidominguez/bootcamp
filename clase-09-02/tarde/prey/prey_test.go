package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSpeed(t *testing.T) {
	t.Run("get speed - using default config", func(t *testing.T) {
		// arrange
		newTuna := prey.CreateTuna()

		// act
		speed := newTuna.GetSpeed()

		// assert
		max := 267.0
		min := 15.0
		require.GreaterOrEqual(t, speed, min)
		require.LessOrEqual(t, speed, max)
	})
	t.Run("get speed - custom speed", func(t *testing.T) {
		// arrange
		newTuna := prey.NewTuna(100, positioner.Position{})

		// act
		speed := newTuna.GetSpeed()

		// assert
		expectedSpeed := 100.0
		require.Equal(t, expectedSpeed, speed)
	})
}

func TestGetPosition(t *testing.T) {
	t.Run("get position - using default config", func(t *testing.T) {
		// arrange
		newTuna := prey.CreateTuna()

		// act
		position := newTuna.GetPosition()

		// assert
		min := 0.0
		max := 500.0
		require.GreaterOrEqual(t, position.X, min)
		require.LessOrEqual(t, position.X, max)
		require.GreaterOrEqual(t, position.Y, min)
		require.LessOrEqual(t, position.Y, max)
		require.GreaterOrEqual(t, position.Z, min)
		require.LessOrEqual(t, position.Z, max)
	})

	t.Run("get position - custom position", func(t *testing.T) {
		// arrange
		newTuna := prey.NewTuna(0, positioner.Position{
			X: 1,
			Y: 2,
			Z: 3,
		})

		// act
		position := newTuna.GetPosition()

		// assert
		expectedPosition := &positioner.Position{
			X: 1,
			Y: 2,
			Z: 3,
		}
		require.Equal(t, expectedPosition, position)
	})
}
