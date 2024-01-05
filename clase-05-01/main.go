package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// it reads the csv file and stores the tickets in the TicketsList variable
	tickets.ReadTickets("tickets.csv")

	fmt.Println("Total tickets: ", tickets.TotalTickets())

	// it gets the total tickets to Brazil
	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println("Error getting total tickets:", err)
		return
	}

	fmt.Println("Total tickets to Brazil: ", total)

	// it gets the total tickets in the morning
	total, err = tickets.GetCountByPeriod("morning")
	if err != nil {
		fmt.Println("Error getting total tickets:", err)
		return
	}

	fmt.Println("Total tickets in the morning: ", total)

	// it gets the average tickets to China
	average, err := tickets.AverageDestination("China")
	if err != nil {
		fmt.Println("Error getting average tickets:", err)
		return
	}

	fmt.Println("Average tickets to China: ", average)

	// it tries to get the total tickets in a period that doesn't exist and shows an error
	_, err = tickets.GetCountByPeriod("not a period")
	if err != nil {
		fmt.Println("Error getting total tickets:", err)
		return
	}
}
