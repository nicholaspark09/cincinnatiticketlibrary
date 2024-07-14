package ticket_team_model_request

import "github.com/nicholaspark09/cincinnatiticketlibrary/model"

type TicketTeamModelUpdateRequest struct {
	UserId string                `json:"user_id"`
	Team   model.TicketTeamModel `json:"team"`
}
