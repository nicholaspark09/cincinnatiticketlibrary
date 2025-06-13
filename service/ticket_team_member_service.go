package service

import (
	json2 "encoding/json"
	"errors"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	network2 "github.com/nicholaspark09/awsgorocket/network"
	"github.com/nicholaspark09/awsgorocket/utils"
	model2 "github.com/nicholaspark09/cincinnatiticketlibrary/model"
	"github.com/nicholaspark09/cincinnatiticketlibrary/model/ticket_team_member_model_request"
	"log"
)

type TicketTeamMemberService struct {
	endpoint       string
	apiKey         string
	contentType    string
	controllerName string
	metricsManager metrics2.MetricsManagerContract
}

func ProvideTicketTeamMemberService(endpoint string,
	apiKey string,
	metrics metrics2.MetricsManagerContract) TicketTeamMemberService {
	return TicketTeamMemberService{
		endpoint:       endpoint,
		apiKey:         apiKey,
		contentType:    "application/json",
		controllerName: "teammembers",
		metricsManager: metrics,
	}
}

func (memberService *TicketTeamMemberService) Create(createRequest ticket_team_member_model_request.TicketTeamMemberModelCreateRequest) response.Response[model2.TicketTeamMemberModel] {
	methodName := "TicketTeamMemberService.Create"
	log.Printf("%s - STARTED - ClientId: %s, TeamId: %s, Title: %s",
		methodName, createRequest.ClientId, createRequest.TicketTeamId, createRequest.Title)

	params := map[string]string{
		"controller": memberService.controllerName,
		"action":     "create",
	}
	manager := network2.ProvideNetworkManager[model2.TicketTeamMemberModel](memberService.endpoint, params, &memberService.apiKey, &memberService.contentType)

	bytes, parseError := json2.Marshal(createRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, ClientId: %s, TeamId: %s",
			methodName, parseError, createRequest.ClientId, createRequest.TicketTeamId)
		return response.Response[model2.TicketTeamMemberModel]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, memberService.metricsManager, func() (*model2.TicketTeamMemberModel, *error) {
		callResponse := network2.Post[model2.TicketTeamMemberModel](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, ClientId: %s, TeamId: %s, Title: %s",
			methodName, callResponse.StatusCode, createRequest.ClientId, createRequest.TicketTeamId, createRequest.Title)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, ClientId: %s, TeamId: %s",
				methodName, callResponse.StatusCode, errorMsg, createRequest.ClientId, createRequest.TicketTeamId)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, ClientId: %s, TeamId: %s, Title: %s",
			methodName, callResponse.StatusCode, createRequest.ClientId, createRequest.TicketTeamId, createRequest.Title)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, ClientId: %s, TeamId: %s",
				methodName, genericError.StatusCode, genericError.Message, createRequest.ClientId, createRequest.TicketTeamId)
			return response.Response[model2.TicketTeamMemberModel]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, ClientId: %s, TeamId: %s",
			methodName, *networkError, createRequest.ClientId, createRequest.TicketTeamId)
		return response.Response[model2.TicketTeamMemberModel]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, ClientId: %s, TeamId: %s, Title: %s",
		methodName, createRequest.ClientId, createRequest.TicketTeamId, createRequest.Title)
	return response.Response[model2.TicketTeamMemberModel]{Data: networkResponse, StatusCode: 200}
}

func (memberService *TicketTeamMemberService) Update(userId string, memberModel model2.TicketTeamMemberModel) response.Response[bool] {
	methodName := "TicketTeamMemberService.Update"
	log.Printf("%s - STARTED - UserId: %s, PK: %s, RK: %s",
		methodName, userId, memberModel.PartitionKey, memberModel.RangeKey)

	params := map[string]string{
		"controller": memberService.controllerName,
		"action":     "update",
	}
	manager := network2.ProvideNetworkManager[bool](memberService.endpoint, params, &memberService.apiKey, &memberService.contentType)

	updateRequest := ticket_team_member_model_request.TicketTeamMemberUpdateRequest{
		UserId:     userId,
		TeamMember: memberModel,
	}

	bytes, parseError := json2.Marshal(updateRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, PK: %s",
			methodName, parseError, userId, memberModel.PartitionKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, memberService.metricsManager, func() (*bool, *error) {
		callResponse := network2.Post[bool](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, userId, memberModel.PartitionKey, memberModel.RangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, PK: %s",
				methodName, callResponse.StatusCode, errorMsg, userId, memberModel.PartitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, userId, memberModel.PartitionKey, memberModel.RangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, PK: %s",
				methodName, genericError.StatusCode, genericError.Message, userId, memberModel.PartitionKey)
			return response.Response[bool]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, PK: %s",
			methodName, *networkError, userId, memberModel.PartitionKey)
		return response.Response[bool]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, PK: %s, RK: %s",
		methodName, userId, memberModel.PartitionKey, memberModel.RangeKey)
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (memberService *TicketTeamMemberService) Delete(deleteRequest model2.DeleteRequest) response.Response[bool] {
	methodName := "TicketTeamMemberService.Delete"
	log.Printf("%s - STARTED - PK: %s, RK: %s, UserId: %s",
		methodName, deleteRequest.PartitionKey, deleteRequest.RangeKey, deleteRequest.UserId)

	params := map[string]string{
		"controller": memberService.controllerName,
		"action":     "delete",
	}
	manager := network2.ProvideNetworkManager[bool](memberService.endpoint, params, &memberService.apiKey, &memberService.contentType)

	bytes, parseError := json2.Marshal(deleteRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, PK: %s, RK: %s",
			methodName, parseError, deleteRequest.PartitionKey, deleteRequest.RangeKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, memberService.metricsManager, func() (*bool, *error) {
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

func (memberService *TicketTeamMemberService) FetchAll(fetchAllRequest ticket_team_member_model_request.TicketTeamMemberModelFetchAllRequest) response.Response[model2.TicketTeamMemberModelsResponse] {
	methodName := "TicketTeamMemberService.FetchAll"
	log.Printf("%s - STARTED - ClientId: %s, TeamId: %s",
		methodName, fetchAllRequest.ClientId, fetchAllRequest.TicketTeamId)

	params := map[string]string{
		"controller": memberService.controllerName,
		"action":     "fetchAll",
	}
	manager := network2.ProvideNetworkManager[model2.TicketTeamMemberModelsResponse](memberService.endpoint, params, &memberService.apiKey, &memberService.contentType)

	bytes, parseError := json2.Marshal(fetchAllRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, ClientId: %s, TeamId: %s",
			methodName, parseError, fetchAllRequest.ClientId, fetchAllRequest.TicketTeamId)
		return response.Response[model2.TicketTeamMemberModelsResponse]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, memberService.metricsManager, func() (*model2.TicketTeamMemberModelsResponse, *error) {
		callResponse := network2.Post[model2.TicketTeamMemberModelsResponse](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, ClientId: %s, TeamId: %s",
			methodName, callResponse.StatusCode, fetchAllRequest.ClientId, fetchAllRequest.TicketTeamId)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, ClientId: %s, TeamId: %s",
				methodName, callResponse.StatusCode, errorMsg, fetchAllRequest.ClientId, fetchAllRequest.TicketTeamId)
			return nil, callResponse.Error
		}

		resultCount := 0
		if callResponse.Data != nil && callResponse.Data.Results != nil {
			resultCount = len(callResponse.Data.Results)
		}
		log.Printf("%s - SUCCESS - StatusCode: %d, ClientId: %s, TeamId: %s, ResultCount: %d",
			methodName, callResponse.StatusCode, fetchAllRequest.ClientId, fetchAllRequest.TicketTeamId, resultCount)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, ClientId: %s, TeamId: %s",
				methodName, genericError.StatusCode, genericError.Message, fetchAllRequest.ClientId, fetchAllRequest.TicketTeamId)
			return response.Response[model2.TicketTeamMemberModelsResponse]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, ClientId: %s, TeamId: %s",
			methodName, *networkError, fetchAllRequest.ClientId, fetchAllRequest.TicketTeamId)
		return response.Response[model2.TicketTeamMemberModelsResponse]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	resultCount := 0
	if networkResponse != nil && networkResponse.Results != nil {
		resultCount = len(networkResponse.Results)
	}
	log.Printf("%s - COMPLETED - StatusCode: 200, ClientId: %s, TeamId: %s, ResultCount: %d",
		methodName, fetchAllRequest.ClientId, fetchAllRequest.TicketTeamId, resultCount)
	return response.Response[model2.TicketTeamMemberModelsResponse]{Data: networkResponse, StatusCode: 200}
}

func (memberService *TicketTeamMemberService) FetchByUser(fetchRequest ticket_team_member_model_request.TicketTeamMemberByUserRequest) response.Response[model2.TicketTeamMemberModelsResponse] {
	methodName := "TicketTeamMemberService.FetchByUser"
	log.Printf("%s - STARTED - UserId: %s", methodName, fetchRequest.UserId)

	params := map[string]string{
		"controller": memberService.controllerName,
		"action":     "fetchByUser",
	}
	manager := network2.ProvideNetworkManager[model2.TicketTeamMemberModelsResponse](memberService.endpoint, params, &memberService.apiKey, &memberService.contentType)

	bytes, parseError := json2.Marshal(fetchRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s",
			methodName, parseError, fetchRequest.UserId)
		return response.Response[model2.TicketTeamMemberModelsResponse]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, memberService.metricsManager, func() (*model2.TicketTeamMemberModelsResponse, *error) {
		callResponse := network2.Post[model2.TicketTeamMemberModelsResponse](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s",
			methodName, callResponse.StatusCode, fetchRequest.UserId)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s",
				methodName, callResponse.StatusCode, errorMsg, fetchRequest.UserId)
			return nil, callResponse.Error
		}

		resultCount := 0
		if callResponse.Data != nil && callResponse.Data.Results != nil {
			resultCount = len(callResponse.Data.Results)
		}
		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, ResultCount: %d",
			methodName, callResponse.StatusCode, fetchRequest.UserId, resultCount)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s",
				methodName, genericError.StatusCode, genericError.Message, fetchRequest.UserId)
			return response.Response[model2.TicketTeamMemberModelsResponse]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s",
			methodName, *networkError, fetchRequest.UserId)
		return response.Response[model2.TicketTeamMemberModelsResponse]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	resultCount := 0
	if networkResponse != nil && networkResponse.Results != nil {
		resultCount = len(networkResponse.Results)
	}
	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, ResultCount: %d",
		methodName, fetchRequest.UserId, resultCount)
	return response.Response[model2.TicketTeamMemberModelsResponse]{Data: networkResponse, StatusCode: 200}
}

func (memberService *TicketTeamMemberService) Fetch(email string, partitionKey string, rangeKey string) response.Response[model2.TicketTeamMemberModel] {
	methodName := "TicketTeamMemberService.Fetch"
	log.Printf("%s - STARTED - Email: %s, PK: %s, RK: %s", methodName, email, partitionKey, rangeKey)

	params := map[string]string{
		"controller": memberService.controllerName,
		"action":     "fetch",
	}
	manager := network2.ProvideNetworkManager[model2.TicketTeamMemberModel](memberService.endpoint, params, &memberService.apiKey, &memberService.contentType)

	fetchRequest := model2.FetchRequest{
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		UserId:       email,
	}

	bytes, parseError := json2.Marshal(fetchRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, Email: %s, PK: %s, RK: %s",
			methodName, parseError, email, partitionKey, rangeKey)
		return response.Response[model2.TicketTeamMemberModel]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, memberService.metricsManager, func() (*model2.TicketTeamMemberModel, *error) {
		callResponse := network2.Post[model2.TicketTeamMemberModel](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, Email: %s, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, email, partitionKey, rangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, Email: %s, PK: %s, RK: %s",
				methodName, callResponse.StatusCode, errorMsg, email, partitionKey, rangeKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, Email: %s, PK: %s, RK: %s",
			methodName, callResponse.StatusCode, email, partitionKey, rangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, Email: %s, PK: %s, RK: %s",
				methodName, genericError.StatusCode, genericError.Message, email, partitionKey, rangeKey)
			return response.Response[model2.TicketTeamMemberModel]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, Email: %s, PK: %s, RK: %s",
			methodName, *networkError, email, partitionKey, rangeKey)
		return response.Response[model2.TicketTeamMemberModel]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, Email: %s, PK: %s, RK: %s",
		methodName, email, partitionKey, rangeKey)
	return response.Response[model2.TicketTeamMemberModel]{Data: networkResponse, StatusCode: 200}
}
