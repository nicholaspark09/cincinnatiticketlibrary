package model

type TicketTeamMemberModel struct {
	// ClientId_TicketTeamModelId
	PartitionKey string `json:"partition_key"`
	// UUID
	RangeKey        string `json:"range_key"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	ObfuscatedEmail string `json:"obfuscated_email"`
	// This will be dependent on the client using the service
	UserId          string `json:"user_id"`
	Created         string `json:"created"`
	Modified        string `json:"modified"`
	AssignedTickets int    `json:"assigned_tickets"`
	// What the user can do; 5= admin, 4 = manager, 3 = hr, 2 = engineer, 1 = intern
	Level int `json:"level"`
}
