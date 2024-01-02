package average

import "fmt"

func main() {
	result, err := average(10.0, 10.0, 9.0)

	if err != "" {
		fmt.Println(err)
		return
	}

	fmt.Printf("The average is: %f\n", result)
}

func average(values ...float64) (float64, string) {
	average := 0.0
	counter := 0

	for _, value := range values {
		if value < 0 {
			return 0.0, "No se pueden introducir notas negativas"
		}
		average += value
		counter++
	}

	return average / float64(counter), ""
}

/*
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.
*/
