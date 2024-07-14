package ticket_team_model_request

type TicketTeamModelFetchAllRequest struct {
	ClientId     string  `json:"client_id"`
	LastRangeKey *string `json:"last_range_key"`
}
