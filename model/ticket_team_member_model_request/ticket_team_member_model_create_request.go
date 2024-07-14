package ticket_team_member_model_request

type TicketTeamMemberModelCreateRequest struct {
	ClientId        string `json:"client_id"`
	TicketTeamId    string `json:"ticket_team_id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Email           string `json:"email"`
	RequesterUserId string `json:"requester_user_id"`
	UserId          string `json:"user_id"`
	Status          string `json:"status"`
	Level           int    `json:"level"`
}
