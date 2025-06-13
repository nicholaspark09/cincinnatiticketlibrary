package model

type TicketWatchModel struct {
	PartitionKey  string `json:"partition_key"` // "{UserId}"
	RangeKey      string `json:"range_key"`     // "{TicketPK}_{TicketRK}"
	Role          string `json:"role"`
	TicketTitle   string `json:"ticket_title"`
	TicketStatus  string `json:"ticket_status"`
	LastUpdated   string `json:"last_updated"`
	UnreadUpdates int    `json:"unread_updates"`
	WatchingSince string `json:"watching_since"`
	Created       string `json:"created"`
	Modified      string `json:"modified"`
}
