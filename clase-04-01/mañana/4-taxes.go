package main

import "fmt"

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
		return fmt.Errorf("the minimum taxable amount is 150000 and the salary entered is: %d", salary)
	}
	return nil
}

/*
Ejercicio 4 - Impuestos de salario #4

Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro
el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir:
“Error: the minimum taxable amount is 150,000 and the salary entered is: [salary]”,
siendo [salary] el valor de tipo int pasado por parámetro).
*/
