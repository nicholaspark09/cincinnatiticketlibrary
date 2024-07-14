package ticket_library

import (
	"github.com/nicholaspark09/awsgorocket/metrics"
	"main/service"
)

type TicketLibrary struct {
	clientId       string
	teamId         string
	ticketEndpoint string
	ticketApiKey   string
	autoCutKey     string
	TicketService  service.TicketService
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
			metricsManager),
	}
}
