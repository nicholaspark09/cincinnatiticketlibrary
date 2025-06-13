package model

type TicketWatchModelsResponse struct {
	Results          []*TicketWatchModel `json:"results"`
	LastPartitionKey *string             `json:"last_partition_key"`
	LastRangeKey     *string             `json:"last_range_key"`
}
