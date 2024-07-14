package model

type TicketTeamMemberModelsResponse struct {
	Results          []*TicketTeamMemberModel `json:"results"`
	LastPartitionKey *string                  `json:"last_partition_key"`
	LastRangeKey     *string                  `json:"last_range_key"`
}
