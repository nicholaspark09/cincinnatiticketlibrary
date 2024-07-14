package model

type FetchRequest struct {
	PartitionKey string `json:"partition_key"`
	RangeKey     string `json:"range_key"`
	UserId       string `json:"user_id"`
}
