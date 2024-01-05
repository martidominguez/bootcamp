package tickets_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTicketsBrazil(t *testing.T) {
	destination := "Brazil"
	expected := 45

	tickets.ReadTickets("../../tickets.csv")

	result, err := tickets.GetTotalTickets(destination)

	require.Equal(t, expected, result)
	require.Nil(t, err)
}

func TestGetTotalTicketsUnknownDestination(t *testing.T) {
	destination := "Unknown"
	expected := 0

	// tickets.ReadTickets("../../tickets.csv") -- only necesary if we run the test alone

	result, err := tickets.GetTotalTickets(destination)

	require.Equal(t, expected, result)
	require.Nil(t, err)
}

func TestGetCountByPeriodMorning(t *testing.T) {
	period := "morning"
	expected := 256

	// tickets.ReadTickets("../../tickets.csv") -- only necesary if we run the test alone

	result, err := tickets.GetCountByPeriod(period)

	require.Equal(t, expected, result)
	require.Nil(t, err)
}

func TestGetCountByPeriodUnknown(t *testing.T) {
	period := "unknown"
	expected := 0

	// tickets.ReadTickets("../../tickets.csv") -- only necesary if we run the test alone

	result, err := tickets.GetCountByPeriod(period)

	require.Equal(t, expected, result)
	require.NotNil(t, err)
}

func TestAverageDestinationChina(t *testing.T) {
	destination := "China"
	expected := 17.8

	// tickets.ReadTickets("../../tickets.csv") -- only necesary if we run the test alone

	result, err := tickets.AverageDestination(destination)

	require.Equal(t, expected, result)
	require.Nil(t, err)
}

func TestAverageDestinationUnknown(t *testing.T) {
	destination := "Unknown"
	expected := 0.0

	// tickets.ReadTickets("../../tickets.csv") -- only necesary if we run the test alone

	result, err := tickets.AverageDestination(destination)

	require.Equal(t, expected, result)
	require.Nil(t, err)
}
