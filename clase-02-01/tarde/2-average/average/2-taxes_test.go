package average

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAveragePositiveGrades(t *testing.T) {
	expected := 9.0
	errExpected := ""

	result, err := average(10.0, 8.0)

	require.Equal(t, errExpected, err)
	require.Equal(t, expected, result)
}

func TestAverageNegativeGrades(t *testing.T) {
	expected := "No se pueden introducir notas negativas"

	result, err := average(-10.0, 8.0)

	require.NotNil(t, err)
	require.Equal(t, 0.0, result)
	require.Equal(t, expected, err)
}
