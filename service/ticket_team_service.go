package service

import (
	json2 "encoding/json"
	"errors"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	network2 "github.com/nicholaspark09/awsgorocket/network"
	"github.com/nicholaspark09/awsgorocket/utils"
	model2 "github.com/nicholaspark09/cincinnatiticketlibrary/model"
	"github.com/nicholaspark09/cincinnatiticketlibrary/model/ticket_team_model_request"
	"log"
)

type TicketTeamService struct {
	endpoint       string
	apiKey         string
	contentType    string
	controllerName string
	metricsManager metrics2.MetricsManagerContract
}

func ProvideTicketTeamService(
	endpoint string,
	apiKey string,
	metricsManager metrics2.MetricsManagerContract,
) TicketTeamService {
	return TicketTeamService{
		endpoint:       endpoint,
		apiKey:         apiKey,
		contentType:    "application/json",
		controllerName: "teams",
		metricsManager: metricsManager,
	}
}

func (teamService *TicketTeamService) Create(createRequest ticket_team_model_request.TicketTeamModelCreateRequest) response.Response[model2.TicketTeamModel] {
	methodName := "TicketTeamService.Create"
	log.Printf("%s - STARTED - ClientId: %s, Title: %s", methodName, createRequest.ClientId, createRequest.Title)

	params := map[string]string{
		"controller": teamService.controllerName,
		"action":     "create",
	}
	manager := network2.ProvideNetworkManager[model2.TicketTeamModel](teamService.endpoint, params, &teamService.apiKey, &teamService.contentType)

	bytes, parseError := json2.Marshal(createRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, ClientId: %s", methodName, parseError, createRequest.ClientId)
		return response.Response[model2.TicketTeamModel]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, teamService.metricsManager, func() (*model2.TicketTeamModel, *error) {
		callResponse := network2.Post[model2.TicketTeamModel](manager, bytes)

		// Log the response details for debugging
		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, ClientId: %s, Title: %s",
			methodName, callResponse.StatusCode, createRequest.ClientId, createRequest.Title)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, ClientId: %s, Title: %s",
				methodName, callResponse.StatusCode, errorMsg, createRequest.ClientId, createRequest.Title)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, ClientId: %s, Title: %s",
			methodName, callResponse.StatusCode, createRequest.ClientId, createRequest.Title)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, ClientId: %s",
				methodName, genericError.StatusCode, genericError.Message, createRequest.ClientId)
			return response.Response[model2.TicketTeamModel]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, ClientId: %s", methodName, *networkError, createRequest.ClientId)
		return response.Response[model2.TicketTeamModel]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, ClientId: %s, Title: %s", methodName, createRequest.ClientId, createRequest.Title)
	return response.Response[model2.TicketTeamModel]{Data: networkResponse, StatusCode: 200}
}

func (teamService *TicketTeamService) Update(userId string, teamModel model2.TicketTeamModel) response.Response[bool] {
	methodName := "TicketTeamService.Update"
	log.Printf("%s - STARTED - UserId: %s, PK: %s, RK: %s", methodName, userId, teamModel.PartitionKey, teamModel.RangeKey)

	params := map[string]string{
		"controller": teamService.controllerName,
		"action":     "update",
	}
	manager := network2.ProvideNetworkManager[bool](teamService.endpoint, params, &teamService.apiKey, &teamService.contentType)

	updateRequest := ticket_team_model_request.TicketTeamModelUpdateRequest{
		UserId: userId,
		Team:   teamModel,
	}

	bytes, parseError := json2.Marshal(updateRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, PK: %s",
			methodName, parseError, userId, teamModel.PartitionKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, teamService.metricsManager, func() (*bool, *error) {
		callResponse := network2.Post[bool](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, userId, teamModel.PartitionKey, teamModel.RangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, PK: %s",
				methodName, callResponse.StatusCode, errorMsg, userId, teamModel.PartitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, userId, teamModel.PartitionKey, teamModel.RangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, PK: %s",
				methodName, genericError.StatusCode, genericError.Message, userId, teamModel.PartitionKey)
			return response.Response[bool]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, PK: %s", methodName, *networkError, userId, teamModel.PartitionKey)
		return response.Response[bool]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, PK: %s, RK: %s",
		methodName, userId, teamModel.PartitionKey, teamModel.RangeKey)
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (teamService *TicketTeamService) Fetch(partitionKey string, rangeKey string, userId string) response.Response[model2.TicketTeamModel] {
	methodName := "TicketTeamService.Fetch"
	log.Printf("%s - STARTED - PK: %s, RK: %s, UserId: %s", methodName, partitionKey, rangeKey, userId)

	params := map[string]string{
		"controller": teamService.controllerName,
		"action":     "fetch",
	}
	manager := network2.ProvideNetworkManager[model2.TicketTeamModel](teamService.endpoint, params, &teamService.apiKey, &teamService.contentType)

	fetchRequest := model2.FetchRequest{
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		UserId:       userId,
	}

	bytes, parseError := json2.Marshal(fetchRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, PK: %s, RK: %s",
			methodName, parseError, partitionKey, rangeKey)
		return response.Response[model2.TicketTeamModel]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, teamService.metricsManager, func() (*model2.TicketTeamModel, *error) {
		callResponse := network2.Post[model2.TicketTeamModel](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, PK: %s, RK: %s, UserId: %s",
			methodName, callResponse.StatusCode, partitionKey, rangeKey, userId)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, PK: %s, RK: %s",
				methodName, callResponse.StatusCode, errorMsg, partitionKey, rangeKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, partitionKey, rangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, PK: %s, RK: %s",
				methodName, genericError.StatusCode, genericError.Message, partitionKey, rangeKey)
			return response.Response[model2.TicketTeamModel]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, PK: %s, RK: %s", methodName, *networkError, partitionKey, rangeKey)
		return response.Response[model2.TicketTeamModel]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, PK: %s, RK: %s", methodName, partitionKey, rangeKey)
	return response.Response[model2.TicketTeamModel]{Data: networkResponse, StatusCode: 200}
}

func (teamService *TicketTeamService) FetchAll(clientId string, lastRangeKey *string) response.Response[model2.TicketTeamModelsResponse] {
	methodName := "TicketTeamService.FetchAll"
	lastRangeKeyStr := "nil"
	if lastRangeKey != nil {
		lastRangeKeyStr = *lastRangeKey
	}
	log.Printf("%s - STARTED - ClientId: %s, LastRangeKey: %s", methodName, clientId, lastRangeKeyStr)

	params := map[string]string{
		"controller": teamService.controllerName,
		"action":     "fetchAll",
	}
	manager := network2.ProvideNetworkManager[model2.TicketTeamModelsResponse](teamService.endpoint, params, &teamService.apiKey, &teamService.contentType)

	fetchAllRequest := ticket_team_model_request.TicketTeamModelFetchAllRequest{
		ClientId:     clientId,
		LastRangeKey: lastRangeKey,
	}

	bytes, parseError := json2.Marshal(fetchAllRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, ClientId: %s",
			methodName, parseError, clientId)
		return response.Response[model2.TicketTeamModelsResponse]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, teamService.metricsManager, func() (*model2.TicketTeamModelsResponse, *error) {
		callResponse := network2.Post[model2.TicketTeamModelsResponse](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, ClientId: %s, LastRangeKey: %s",
			methodName, callResponse.StatusCode, clientId, lastRangeKeyStr)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, ClientId: %s",
				methodName, callResponse.StatusCode, errorMsg, clientId)
			return nil, callResponse.Error
		}

		resultCount := 0
		if callResponse.Data != nil && callResponse.Data.Results != nil {
			resultCount = len(callResponse.Data.Results)
		}
		log.Printf("%s - SUCCESS - StatusCode: %d, ClientId: %s, ResultCount: %d",
			methodName, callResponse.StatusCode, clientId, resultCount)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, ClientId: %s",
				methodName, genericError.StatusCode, genericError.Message, clientId)
			return response.Response[model2.TicketTeamModelsResponse]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, ClientId: %s", methodName, *networkError, clientId)
		return response.Response[model2.TicketTeamModelsResponse]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	resultCount := 0
	if networkResponse != nil && networkResponse.Results != nil {
		resultCount = len(networkResponse.Results)
	}
	log.Printf("%s - COMPLETED - StatusCode: 200, ClientId: %s, ResultCount: %d",
		methodName, clientId, resultCount)
	return response.Response[model2.TicketTeamModelsResponse]{Data: networkResponse, StatusCode: 200}
}

// Delete method implementation (uncommented and fixed)
func (teamService *TicketTeamService) Delete(deleteRequest model2.DeleteRequest) response.Response[bool] {
	methodName := "TicketTeamService.Delete"
	log.Printf("%s - STARTED - PK: %s, RK: %s, UserId: %s",
		methodName, deleteRequest.PartitionKey, deleteRequest.RangeKey, deleteRequest.UserId)

	params := map[string]string{
		"controller": teamService.controllerName,
		"action":     "delete",
	}
	manager := network2.ProvideNetworkManager[bool](teamService.endpoint, params, &teamService.apiKey, &teamService.contentType)

	bytes, parseError := json2.Marshal(deleteRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, PK: %s, RK: %s",
			methodName, parseError, deleteRequest.PartitionKey, deleteRequest.RangeKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, teamService.metricsManager, func() (*bool, *error) {
		callResponse := network2.Post[bool](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, deleteRequest.PartitionKey, deleteRequest.RangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, PK: %s, RK: %s",
				methodName, callResponse.StatusCode, errorMsg, deleteRequest.PartitionKey, deleteRequest.RangeKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, deleteRequest.PartitionKey, deleteRequest.RangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, PK: %s, RK: %s",
				methodName, genericError.StatusCode, genericError.Message, deleteRequest.PartitionKey, deleteRequest.RangeKey)
			return response.Response[bool]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, PK: %s, RK: %s",
			methodName, *networkError, deleteRequest.PartitionKey, deleteRequest.RangeKey)
		return response.Response[bool]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, PK: %s, RK: %s",
		methodName, deleteRequest.PartitionKey, deleteRequest.RangeKey)
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}
