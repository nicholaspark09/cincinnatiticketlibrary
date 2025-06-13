package ticket_comment_request

import (
	model2 "github.com/nicholaspark09/cincinnatiticketlibrary/model"
)

type TicketCommentModelCreateRequest struct {
	TicketPartitionKey string `json:"ticket_partition_key"`
	TicketRangeKey     string `json:"ticket_range_key"`
	UserId             string `json:"user_id"`
	Message            string `json:"message"`
	Files              string `json:"files"`
}

type TicketCommentModelFetchAllRequest struct {
	TicketPartitionKey string  `json:"ticket_partition_key"`
	TicketRangeKey     string  `json:"ticket_range_key"`
	UserId             string  `json:"user_id"`
	LastRangeKey       *string `json:"last_range_key"`
}

type TicketCommentModelByUserRequest struct {
	UserId           string  `json:"user_id"`
	LastPartitionKey *string `json:"last_partition_key"`
	LastRangeKey     *string `json:"last_range_key"`
}

type TicketCommentModelUpdateRequest struct {
	UserId  string                    `json:"user_id"`
	Comment model2.TicketCommentModel `json:"comment"`
}
