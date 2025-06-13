package model

type TicketCommentModel struct {
	// TicketPK_TicketRK
	PartitionKey string `json:"partition_key"`
	// Timestamp so we can sort
	RangeKey string `json:"range_key"`
	UserId   string `json:"user_id"`
	Message  string `json:"message"`
	Files    string `json:"files"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}
