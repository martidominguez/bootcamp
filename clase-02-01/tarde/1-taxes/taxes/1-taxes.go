package taxes

import "fmt"

func main() {
	salary := 200000.0

	taxes := calculatorTaxes(salary)

	fmt.Printf("Taxes: %f\n", taxes)
}

func calculatorTaxes(salary float64) float64 {
	switch {
	case salary > 150000:
		return salary * 0.27
	case salary > 50000:
		return salary * 0.17
	default:
		return 0
	}
}

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo
y si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/
