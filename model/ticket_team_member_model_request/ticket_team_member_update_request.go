package ticket_team_member_model_request

import "main/model"

type TicketTeamMemberUpdateRequest struct {
	UserId     string                      `json:"user_id"`
	TeamMember model.TicketTeamMemberModel `json:"team_member"`
}
