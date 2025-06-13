package ticket_watch_request

type TicketWatchAddRequest struct {
	UserId             string `json:"user_id"`
	TicketPartitionKey string `json:"ticket_partition_key"`
	TicketRangeKey     string `json:"ticket_range_key"`
	Role               string `json:"role"`
}

type TicketWatchRemoveRequest struct {
	UserId       string `json:"user_id"`
	PartitionKey string `json:"partition_key"`
	RangeKey     string `json:"range_key"`
}

type TicketWatchUserListRequest struct {
	UserId       string  `json:"user_id"`
	LastRangeKey *string `json:"last_range_key"`
}

type TicketWatchersListRequest struct {
	TicketPartitionKey string  `json:"ticket_partition_key"`
	TicketRangeKey     string  `json:"ticket_range_key"`
	LastPartitionKey   *string `json:"last_partition_key"`
	LastRangeKey       *string `json:"last_range_key"`
	UserId             string  `json:"user_id"`
}

type TicketWatchMarkReadRequest struct {
	UserId             string `json:"user_id"`
	TicketPartitionKey string `json:"ticket_partition_key"`
	TicketRangeKey     string `json:"ticket_range_key"`
}

type TicketWatchUpdateRequest struct {
	UserId             string `json:"user_id"`
	TicketPartitionKey string `json:"ticket_partition_key"`
	TicketRangeKey     string `json:"ticket_range_key"`
	TicketTitle        string `json:"ticket_title"`
	TicketStatus       string `json:"ticket_status"`
	LastUpdated        string `json:"last_updated"`
}
