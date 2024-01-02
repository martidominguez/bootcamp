package main

import "fmt"

func main() {
	result := average(10.0, 10.0, 9.0)
	fmt.Printf("The average is: %f\n", result)
}

func average(values ...float64) float64 {
	average := 0.0
	counter := 0

	for _, value := range values {
		average += value
		counter++
	}

	return average / float64(counter)
}

/*
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.
*/
