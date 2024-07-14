package ticket_team_model_request

import "main/model"

type TicketTeamModelUpdateRequest struct {
	UserId string                `json:"user_id"`
	Team   model.TicketTeamModel `json:"team"`
}
