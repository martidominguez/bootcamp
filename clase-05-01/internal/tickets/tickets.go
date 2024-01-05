package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string // i could change this to Time
	Price              int
}

var TicketsList []Ticket

// example 1
func GetTotalTickets(destination string) (int, error) {
	total := 0

	for _, ticket := range TicketsList {
		if ticket.DestinationCountry == destination {
			total++
		}
	}

	return total, nil
}

// example 2
func GetCountByPeriod(period string) (int, error) {
	total := 0

	if period != "early morning" && period != "morning" && period != "afternoon" && period != "night" {
		return 0, errors.New("period not found")
	}

	for _, ticket := range TicketsList {
		// it splits the time because we only care about the hour
		fields := strings.Split(ticket.FlightTime, ":")

		// it converts the string to int to be able to compare it
		time, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("cannot convert string to int", err)
			return 0, err
		}

		// if it matches the period, it adds 1 to the total
		switch {
		case period == "early morning" && time >= 0 && time <= 6:
			total++
		case period == "morning" && time >= 7 && time <= 12:
			total++
		case period == "afternoon" && time >= 13 && time <= 19:
			total++
		case period == "night" && time >= 20 && time <= 23:
			total++
		}
	}

	return total, nil
}

// example 3
func AverageDestination(destination string) (float64, error) {
	total := TotalTickets()

	totalDestination, err := GetTotalTickets(destination)
	if err != nil {
		return 0.0, errors.New("cannot get total tickets for destination")
	}

	return float64(totalDestination) / float64(total) * 100, nil
}

func TotalTickets() int {
	return len(TicketsList)
}

func ReadTickets(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	for _, row := range rows {
		ticket := CreateTicket(row)
		TicketsList = append(TicketsList, ticket)
	}

	return
}

func CreateTicket(row []string) (ticket Ticket) {
	id, err := strconv.Atoi(row[0])
	if err != nil {
		fmt.Println("cannot convert string to int", err)
	}

	price, err := strconv.Atoi(row[5])
	if err != nil {
		fmt.Println("cannot convert string to int", err)
	}

	ticket = Ticket{
		ID:                 id,
		Name:               row[1],
		Email:              row[2],
		DestinationCountry: row[3],
		FlightTime:         row[4],
		Price:              price,
	}
	return
}
