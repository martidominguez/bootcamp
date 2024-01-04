package main

import (
	"errors"
	"fmt"
)

func main() {
	hours := 100.0
	value := 100.0

	salary, err := calculateSalary(hours, value)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("The salary is: %.1f\n", salary)
}

func calculateSalary(hours float64, value float64) (float64, error) {
	salary := hours * value

	switch {
	case hours < 80:
		return 0, errors.New("the worker cannot have worked less than 80 hours per month")
	case salary >= 150000:
		return salary * 0.9, nil
	default:
		return salary, nil
	}
}

/*
Ejercicio 5 -  Impuestos de salario #5

Vamos a hacer que nuestro programa sea un poco más complejo y útil.

Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error.
El mismo tendrá que indicar “Error: the worker cannot have worked less than 80 hours per month”.
*/
