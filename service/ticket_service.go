package service

import (
	"encoding/json"
	"github.com/nicholaspark09/awsgorocket/metrics"
	"github.com/nicholaspark09/awsgorocket/network_v2"
	"github.com/nicholaspark09/cincinnatiticketlibrary/model"
	"github.com/nicholaspark09/cincinnatiticketlibrary/model/ticket_model_request"
	"log"
)

type TicketService struct {
	Endpoint       string
	ApiKey         string
	ClientId       string
	ContentType    string
	TeamId         string
	AutoCutKey     string
	metricsManager metrics.MetricsManagerContract
}

func ProvideTicketService(
	endpoint string,
	apiKey string,
	clientId string,
	teamId string,
	autoCutKey string,
	metricsManager metrics.MetricsManagerContract,
) TicketService {
	return TicketService{
		Endpoint:       endpoint,
		ApiKey:         apiKey,
		ClientId:       clientId,
		ContentType:    "application/json",
		TeamId:         teamId,
		metricsManager: metricsManager,
	}
}

func (ticketService *TicketService) CreateAutocut(
	title string,
	description string,
	files string,
	severity int,
) bool {
	params := map[string]string{
		"action": "create",
	}
	manager := network_v2.ProvideNetworkManagerV2[model.TicketModel](ticketService.Endpoint, params, &ticketService.ApiKey, &ticketService.ContentType)
	bytes, err := json.Marshal(ticket_model_request.TicketModelCreateRequest{
		ClientId:     ticketService.ClientId,
		TeamRangeKey: ticketService.TeamId,
		Title:        title,
		Description:  description,
		Files:        files,
		Severity:     severity,
		UserId:       ticketService.AutoCutKey,
		Status:       "OPEN",
	})
	if err != nil {
		log.Printf("TicketService.CreateFailure - Error in converting request to json: %s", err.Error())
		return false
	}
	networkResponse, _ := metrics.MeasureTimeWithError("CincinnatiTicketService.create", ticketService.metricsManager, func() (*model.TicketModel, *error) {
		callResponse, callErr := network_v2.Post[model.TicketModel](manager, bytes)
		if callErr != nil {
			log.Printf("TicketService.CreateFailure - Error in calling ticket service: %s", callErr)
			return nil, &callErr
		}
		return callResponse, nil
	})
	if networkResponse != nil {
		log.Printf("TicketService.CreateSuccess - Successfully created a ticket with PK: %s, RK: %s", networkResponse.PartitionKey, networkResponse.RangeKey)
	} else {
		log.Printf("TicketService.CreateFailure - Failed to create a ticket with PK: %s, RK: %s", networkResponse.PartitionKey, networkResponse.RangeKey)
	}
	return networkResponse != nil
}

func (ticketService *TicketService) Create(
	title string,
	description string,
	files string,
	severity int,
	userId string,
) {

}
