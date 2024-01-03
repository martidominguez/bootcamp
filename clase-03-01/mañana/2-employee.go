package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func main() {
	person := Person{
		ID:          1,
		Name:        "Martina Dominguez",
		DateOfBirth: "2003-05-14",
	}

	employee := Employee{
		ID:       1,
		Position: "Software Engineer",
		Person:   person,
	}

	employee.PrintEmployee()
}

func (employee Employee) PrintEmployee() {
	fmt.Printf("ID: %d\n", employee.ID)
	fmt.Printf("Position: %s\n", employee.Position)
	fmt.Printf("Person ID: %d\n", employee.Person.ID)
	fmt.Printf("Name: %s\n", employee.Person.Name)
	fmt.Printf("Date of Birth: %s\n", employee.Person.DateOfBirth)
	return
}
