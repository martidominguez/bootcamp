package loader

import (
	internal "app"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (tMap map[int]internal.TicketAttributes, err error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	tMap = make(map[int]internal.TicketAttributes)
	var record []string
	for {
		record, err = r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			err = fmt.Errorf("error reading record: %v", err)
			return
		}

		// serialize the record
		idStr := record[0]
		id, parseErr := strconv.Atoi(idStr)
		if parseErr != nil {
			err = fmt.Errorf("error converting ID to int: %v", parseErr)
			return
		}

		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
		}

		// Parse Price as a float64
		var price float64
		price, parseErr = strconv.ParseFloat(record[5], 64)
		if parseErr != nil {
			err = fmt.Errorf("error converting Price to float64: %v", parseErr)
			return
		}
		ticket.Price = price

		// add the ticket to the map
		tMap[id] = ticket
	}

	return tMap, nil
}
