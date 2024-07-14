package ticket_model_request

type TicketModelByUserRequest struct {
	UserId           string  `json:"user_id"`
	LastPartitionKey *string `json:"last_partition_key"`
	LastRangeKey     *string `json:"last_range_key"`
}
