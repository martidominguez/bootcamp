package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Benjamin's age is %d\n", employees["Benjamin"])

	counter := 0

	for _, age := range employees {
		if age >= 21 {
			counter++
		}
	}

	fmt.Printf("%d employees are older than 21\n", counter)

	employees["Federico"] = 25

	delete(employees, "Pedro")
}
