package taxes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculatorTaxesUnder50(t *testing.T) {
	salary := 20000.0
	expected := 0.0

	result := calculatorTaxes(salary)

	require.Equal(t, expected, result)
}

func TestCalculatorTaxesOver50(t *testing.T) {
	salary := 100000.0
	expected := 17000.0

	result := calculatorTaxes(salary)

	require.Equal(t, expected, result)
}

func TestCalculatorTaxesOver150(t *testing.T) {
	salary := 200000.0
	expected := 54000.0

	result := calculatorTaxes(salary)

	require.Equal(t, expected, result)
}

/*
La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados al momento de depositar
el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos impuestos están correctos.
Para esto nos encargaron el trabajo de realizar los test correspondientes para:

Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/
