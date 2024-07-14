package model

type DeleteRequest struct {
	PartitionKey string `json:"partition_key"`
	RangeKey     string `json:"range_key"`
	IsHardDelete bool   `json:"is_hard_delete"`
	UserId       string `json:"user_id"`
}
