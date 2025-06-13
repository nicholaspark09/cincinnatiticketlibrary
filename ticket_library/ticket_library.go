package ticket_library

import (
	"github.com/nicholaspark09/awsgorocket/metrics"
	"github.com/nicholaspark09/cincinnatiticketlibrary/service"
)

type TicketLibrary struct {
	clientId                string
	teamId                  string
	ticketEndpoint          string
	ticketApiKey            string
	autoCutKey              string
	TicketService           service.TicketService
	TicketCommentService    service.TicketCommentService
	TicketWatchService      service.TicketWatchService
	TicketTeamService       service.TicketTeamService
	TicketTeamMemberService service.TicketTeamMemberService
}

func ProvideTicketLibrary(
	clientId string,
	teamId string,
	ticketEndpoint string,
	ticketApiKey string,
	autoCutKey string,
	metricsManager metrics.MetricsManagerContract) TicketLibrary {
	return TicketLibrary{
		clientId:       clientId,
		teamId:         teamId,
		ticketEndpoint: ticketEndpoint,
		ticketApiKey:   ticketApiKey,
		TicketService: service.ProvideTicketService(
			ticketEndpoint,
			ticketApiKey,
			clientId,
			teamId,
			autoCutKey,
			metricsManager),
		TicketCommentService:    service.ProvideTicketCommentService(ticketEndpoint, ticketApiKey, metricsManager),
		TicketWatchService:      service.ProvideTicketWatchService(ticketEndpoint, ticketApiKey, metricsManager),
		TicketTeamService:       service.ProvideTicketTeamService(ticketEndpoint, ticketApiKey, metricsManager),
		TicketTeamMemberService: service.ProvideTicketTeamMemberService(ticketEndpoint, ticketApiKey, metricsManager),
	}
}
