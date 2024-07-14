package model

type HRNotificationModel struct {
	ClientId     string `json:"client_id"`
	From         string `json:"from"`
	To           string `json:"to"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	Type         string `json:"type"`
	TeamMemberId string `json:"team_member_id"`
	CompanyId    string `json:"company_id"`
}
