package ticket_model_request

type TicketModelFetchAllRequest struct {
	ClientId     string  `json:"client_id"`
	TeamId       string  `json:"team_id"`
	UserId       string  `json:"user_id"`
	LastRangeKey *string `json:"last_range_key"`
}
