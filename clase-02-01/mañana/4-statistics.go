package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	operationName := minimum

	min := operation(operationName)

	minValue, err := min(2, 3, 3, 4, 10, 2, 4, -10)

	if err != "" {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Min: %d\n", minValue)

}

func operation(operation string) func(...int) (int, string) {
	switch operation {
	case minimum:
		return minFunc
	case average:
		return averageFunc
	case maximum:
		return maxFunc
	default:
		return nil
	}
}

func minFunc(values ...int) (min int, err string) {
	min = values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return
}

func averageFunc(values ...int) (average int, err string) {
	average = 0
	counter := 0
	for _, value := range values {
		average += value
		counter++
	}
	average /= counter
	return
}

func maxFunc(values ...int) (max int, err string) {
	max = values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return
}

/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes
de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una
cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.

Ejemplo:
const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)
...
minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)
...
minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/
