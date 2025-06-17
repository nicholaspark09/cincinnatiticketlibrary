package service

import (
	"encoding/json"
	"errors"
	"github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	network2 "github.com/nicholaspark09/awsgorocket/network"
	"github.com/nicholaspark09/awsgorocket/network_v2"
	"github.com/nicholaspark09/awsgorocket/utils"
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
		AutoCutKey:     autoCutKey,
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

func (ticketService *TicketService) Fetch(partitionKey string, rangeKey string) response.Response[model.TicketModel] {
	params := map[string]string{
		"controller": "tickets",
		"action":     "fetch",
	}
	manager := network2.ProvideNetworkManager[model.TicketModel](ticketService.Endpoint, params, &ticketService.ApiKey, &ticketService.ContentType)
	bytes, parseError := json.Marshal(model.FetchRequest{
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		UserId:       "",
	})
	if parseError != nil {
		log.Printf("TicketService.FetchAllFailure - Failed to parse FetchRequest properly: %v", parseError)
		return response.Response[model.TicketModel]{StatusCode: 400, Message: "Incorrect body"}
	}
	networkResponse, networkError := metrics.MeasureTimeWithError("TicketService.FetchAll", ticketService.metricsManager, func() (*model.TicketModel, *error) {
		callResponse := network2.Post[model.TicketModel](manager, bytes)
		if callResponse.StatusCode != 200 {
			log.Printf("TicketService.FetchFailure - StatusCode: %v, Failed to fetch podmembers because: %v", callResponse.StatusCode, callResponse.Error)
			return nil, callResponse.Error
		} else {
			log.Printf("TicketService.FetchSuccess - StatusCode: %v, Successfully fetched for PK: %s, RK: %s", callResponse.StatusCode, partitionKey, rangeKey)
		}
		return callResponse.Data, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		log.Printf("TicketService.FetchFailure - Failed to fetch any podmembers. Error: %v", networkError)
		return response.Response[model.TicketModel]{StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	return response.Response[model.TicketModel]{Data: networkResponse, StatusCode: 200}
}
