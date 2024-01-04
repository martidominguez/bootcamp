package main

import (
	"errors"
	"fmt"
)

type errorSalary struct {
	message string
}

func (e *errorSalary) Error() string {
	return e.message
}

var errSalary = &errorSalary{message: "salary is less than 10000"}

func main() {
	var salary int = 9000

	err := checkSalary(salary)

	if errors.Is(err, errSalary) {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Must pay tax\n")
}

func checkSalary(salary int) error {
	if salary <= 10000 {
		return errSalary
	}
	return nil
}

/*
Ejercicio 2 - Impuestos de salario #2

En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: salary is less than 10000"
y lanzalo en caso de que “salary” sea menor o igual a  10000.
La validación debe ser hecha con la función “Is()” dentro del “main”.
*/
