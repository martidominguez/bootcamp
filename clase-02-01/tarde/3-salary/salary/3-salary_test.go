package salary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSalaryCalculatorCategoryC(t *testing.T) {
	minutes := 4800.0
	category := "C"
	expected := 80000.0

	salary := salaryCalculator(minutes, category)

	require.Equal(t, expected, salary)
}

func TestSalaryCalculatorCategoryB(t *testing.T) {
	minutes := 4800.0
	category := "B"
	expected := 144000.0

	salary := salaryCalculator(minutes, category)

	require.Equal(t, expected, salary)
}

func TestSalaryCalculatorCategoryA(t *testing.T) {
	minutes := 4800.0
	category := "A"
	expected := 360000.0

	salary := salaryCalculator(minutes, category)

	require.Equal(t, expected, salary)
}
