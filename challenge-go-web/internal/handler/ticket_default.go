package handler

import (
	internal "app"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

// HandlerTicketsDefault represents the default handler of the tickets
type HandlerTicketsDefault struct {
	// sv represents the service of the tickets
	sv internal.ServiceTicket
}

// NewHandlerTicketsDefault creates a new default handler of the tickets
func NewHandlerTicketsDefault(sv internal.ServiceTicket) *HandlerTicketsDefault {
	return &HandlerTicketsDefault{
		sv: sv,
	}
}

func (h *HandlerTicketsDefault) GetTotalTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the total number of tickets
		total, err := h.sv.GetTotalAmountTickets()
		if err != nil {
			// error response
			response.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
		// response
		response.JSON(w, http.StatusOK, map[string]any{"message": "tickets found", "total": total})
	}
}
