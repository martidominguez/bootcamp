package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 100000

	err := checkSalary(salary)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Must pay tax\n")
}

func checkSalary(salary int) error {
	if salary < 150000 {
		return errors.New("The salary entered does not reach the taxable minimum")
	}
	return nil
}

/*
Ejercicio 1 - Impuestos de salario #1

En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error:
the salary entered does not reach the taxable minimum" y lanzalo en caso de que “salary” sea menor a 150.000.
De lo contrario, tendrás que imprimir por consola el mensaje “Must pay tax”.
*/
