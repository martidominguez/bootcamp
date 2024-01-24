package service

import (
	internal "app"
	"context"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	// create the context
	ctx := context.Background()
	// get all the tickets
	tickets, err := s.rp.Get(ctx)
	if err != nil {
		return
	}
	// get the total number of tickets
	total = len(tickets)
	return
}
