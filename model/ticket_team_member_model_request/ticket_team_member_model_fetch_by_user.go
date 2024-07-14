package ticket_team_member_model_request

type TicketTeamMemberByUserRequest struct {
	UserId           string  `json:"user_id"`
	LastPartitionKey *string `json:"last_partition_key"`
	LastRangeKey     *string `json:"last_range_key"`
}
