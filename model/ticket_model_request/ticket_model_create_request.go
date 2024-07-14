package ticket_model_request

type TicketModelCreateRequest struct {
	ClientId     string `json:"client_id"`
	TeamRangeKey string `json:"team_range_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Files        string `json:"files"`
	Severity     int    `json:"severity"`
	UserId       string `json:"user_id"`
	Status       string `json:"status"`
}
