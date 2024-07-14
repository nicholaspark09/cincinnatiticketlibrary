package ticket_team_member_model_request

type TicketTeamMemberModelFetchAllRequest struct {
	ClientId     string  `json:"client_id"`
	TicketTeamId string  `json:"ticket_team_id"`
	UserId       string  `json:"user_id"`
	LastRangeKey *string `json:"last_range_key"`
}
