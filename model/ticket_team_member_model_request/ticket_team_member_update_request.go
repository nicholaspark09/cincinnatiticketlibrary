package ticket_team_member_model_request

import "github.com/nicholaspark09/cincinnatiticketlibrary/model"

type TicketTeamMemberUpdateRequest struct {
	UserId     string                      `json:"user_id"`
	TeamMember model.TicketTeamMemberModel `json:"team_member"`
}
