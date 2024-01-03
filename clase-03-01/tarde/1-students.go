package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	Dni       int
	EntryDate string
}

func (s Student) Details() {
	fmt.Printf("First name: %s\n", s.FirstName)
	fmt.Printf("Last name: %s\n", s.LastName)
	fmt.Printf("Dni: %d\n", s.Dni)
	fmt.Printf("Entry date: %s\n", s.EntryDate)
}

func main() {
	student := Student{
		FirstName: "Martina",
		LastName:  "Dominguez",
		Dni:       44749012,
		EntryDate: "2023-12-20",
	}

	student.Details()
}
