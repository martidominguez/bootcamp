package statistics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimum(t *testing.T) {
	operationName := minimum

	min := operation(operationName)

	minValue, err := min(2, 3, 3, 4, 10, 2, 4, -10)

	require.Equal(t, "", err)
	require.Equal(t, -10, minValue)
}

func TestAverage(t *testing.T) {
	operationName := average

	average := operation(operationName)

	averageValue, err := average(2, 3, 3, 4, 10, 2, 4)

	require.Equal(t, "", err)
	require.Equal(t, 4, averageValue)
}

func TestMaximum(t *testing.T) {
	operationName := maximum

	max := operation(operationName)

	maxValue, err := max(2, 3, 3, 4, 10, 2, 4, -10)

	require.Equal(t, "", err)
	require.Equal(t, 10, maxValue)
}

func TestNotExistingOperation(t *testing.T) {
	operationName := "not existing operation"

	operation := operation(operationName)

	require.Nil(t, operation)
}
