package main

import "fmt"

func main() {
	minutes := 4800.0

	salary := salaryCalculator(minutes, "A")

	fmt.Printf("Salary: %f\n", salary)
}

func salaryCalculator(minutes float64, category string) float64 {
	hours := minutes / 60

	switch category {
	case "C":
		return hours * 1000
	case "B":
		salary := hours * 1500
		return salary + salary*0.2
	case "A":
		salary := hours * 3000
		return salary + salary*0.5
	default:
		return -1
	}

}

/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes
y la categoría.

Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes,
la categoría y que devuelva su salario.

*/
