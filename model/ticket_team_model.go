package model

type TicketTeamModel struct {
	// ClientId
	PartitionKey string `json:"partition_key"`
	// UUID so we can sort
	RangeKey    string `json:"range_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"user_id"`
	Category    string `json:"category"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
	Status      string `json:"status"`
}
