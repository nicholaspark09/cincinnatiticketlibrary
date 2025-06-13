package model

type TicketCommentModelsResponse struct {
	Results      []*TicketCommentModel `json:"results"`
	LastRangeKey *string               `json:"last_range_key"`
}
