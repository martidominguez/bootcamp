package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 9000

	err := checkSalary(salary)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Must pay tax\n")
}

func checkSalary(salary int) error {
	if salary <= 10000 {
		return errors.New("salary is less than 10000")
	}
	return nil
}

/*
Ejercicio 3 - Impuestos de salario #3

Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()”.
*/
