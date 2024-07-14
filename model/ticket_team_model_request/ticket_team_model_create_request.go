package ticket_team_model_request

type TicketTeamModelCreateRequest struct {
	ClientId    string `json:"client_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	UserId      string `json:"user_id"`
	Status      string `json:"status"`
}
