package model

type TicketModel struct {
	// ClientId_TicketTeamModelId
	PartitionKey string `json:"partition_key"`
	// Time.UUID so we can sort
	RangeKey             string `json:"range_key"`
	Title                string `json:"title"`
	Description          string `json:"description"`
	Category             string `json:"category"`
	Comments             string `json:"comments"`
	Files                string `json:"files"`
	Severity             int    `json:"severity"`
	Status               string `json:"status"`
	StatusHistory        string `json:"status_history"`
	AssignedUserId       string `json:"assigned_user_id"`
	UserId               string `json:"user_id"`
	Created              string `json:"created"`
	Modified             string `json:"modified"`
	ResolutionLimit      string `json:"resolution_limit"`
	CampaignPartitionKey string `json:"campaign_partition_key"`
	CampaignRangeKey     string `json:"campaign_range_key"`
}
