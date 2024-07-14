package model

type TicketModelsResponse struct {
	Results          []*TicketModel `json:"results"`
	LastPartitionKey *string        `json:"last_partition_key"`
	LastRangeKey     *string        `json:"last_range_key"`
}
