package ticket_model_request

import "github.com/nicholaspark09/cincinnatiticketlibrary/model"

type TicketModelUpdateRequest struct {
	UserId string            `json:"user_id"`
	Ticket model.TicketModel `json:"ticket"`
}
