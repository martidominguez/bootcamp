package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// arrenge
		whiteShark := hunt.NewWhiteShark(true, false, 1)
		tuna := hunt.NewTuna("tuna", 0.5)

		// act
		err := whiteShark.Hunt(tuna)

		// assert
		require.NoError(t, err)
		require.False(t, whiteShark.Hungry)
		require.True(t, whiteShark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// arrenge
		whiteSharkNotHungry := hunt.NewWhiteShark(false, true, 0)
		tuna := hunt.NewTuna("tuna", 0)

		// act
		err := whiteSharkNotHungry.Hunt(tuna)

		// assert
		require.Error(t, err)
		expectedError := hunt.ErrSharkIsNotHungry
		require.Equal(t, expectedError, err)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// arrenge
		whiteSharkTired := hunt.NewWhiteShark(true, true, 0)
		tuna := hunt.NewTuna("tuna", 0)

		// act
		err := whiteSharkTired.Hunt(tuna)

		// assert
		require.Error(t, err)
		expectedError := hunt.ErrSharkIsTired
		require.Equal(t, expectedError, err)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// arrenge
		whiteSharkSlower := hunt.NewWhiteShark(true, false, 0)
		tuna := hunt.NewTuna("tuna", 1)

		// act
		err := whiteSharkSlower.Hunt(tuna)

		// assert
		require.Error(t, err)
		expectedError := hunt.ErrSharkIsSlower
		require.Equal(t, expectedError, err)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// arrenge
		whiteShark := hunt.NewWhiteShark(true, false, 1)
		var tuna *hunt.Tuna

		// act
		err := whiteShark.Hunt(tuna)

		// assert
		require.Error(t, err)
		expectedError := hunt.ErrTunaIsNil
		require.Equal(t, expectedError, err)
	})
}
