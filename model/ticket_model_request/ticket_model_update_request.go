package ticket_model_request

import "main/model"

type TicketModelUpdateRequest struct {
	UserId string            `json:"user_id"`
	Ticket model.TicketModel `json:"ticket"`
}
