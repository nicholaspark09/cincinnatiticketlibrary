package model

type TicketTeamModelsResponse struct {
	Results          []*TicketTeamModel `json:"results"`
	LastPartitionKey *string            `json:"last_partition_key"`
	LastRangeKey     *string            `json:"last_range_key"`
}
