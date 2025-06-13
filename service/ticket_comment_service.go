package service

import (
	json2 "encoding/json"
	"errors"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	network2 "github.com/nicholaspark09/awsgorocket/network"
	"github.com/nicholaspark09/awsgorocket/utils"
	model2 "github.com/nicholaspark09/cincinnatiticketlibrary/model"
	"github.com/nicholaspark09/cincinnatiticketlibrary/model/ticket_comment_request"
	"log"
)

type TicketCommentService struct {
	endpoint       string
	apiKey         string
	contentType    string
	controllerName string
	metricsManager metrics2.MetricsManagerContract
}

func ProvideTicketCommentService(
	endpoint string,
	apiKey string,
	metricsManager metrics2.MetricsManagerContract,
) TicketCommentService {
	return TicketCommentService{
		endpoint:       endpoint,
		apiKey:         apiKey,
		contentType:    "application/json",
		controllerName: "ticket-comments",
		metricsManager: metricsManager,
	}
}

func (commentService *TicketCommentService) Create(createRequest ticket_comment_request.TicketCommentModelCreateRequest) response.Response[model2.TicketCommentModel] {
	methodName := "TicketCommentService.Create"
	log.Printf("%s - STARTED - UserId: %s, TicketPK: %s, TicketRK: %s, MessageLength: %d",
		methodName, createRequest.UserId, createRequest.TicketPartitionKey, createRequest.TicketRangeKey, len(createRequest.Message))

	params := map[string]string{
		"controller": commentService.controllerName,
		"action":     "create",
	}
	manager := network2.ProvideNetworkManager[model2.TicketCommentModel](commentService.endpoint, params, &commentService.apiKey, &commentService.contentType)

	bytes, parseError := json2.Marshal(createRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, TicketPK: %s",
			methodName, parseError, createRequest.UserId, createRequest.TicketPartitionKey)
		return response.Response[model2.TicketCommentModel]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, commentService.metricsManager, func() (*model2.TicketCommentModel, *error) {
		callResponse := network2.Post[model2.TicketCommentModel](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, createRequest.UserId, createRequest.TicketPartitionKey, createRequest.TicketRangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, TicketPK: %s",
				methodName, callResponse.StatusCode, errorMsg, createRequest.UserId, createRequest.TicketPartitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s, CommentPK: %s, CommentRK: %s",
			methodName, callResponse.StatusCode, createRequest.UserId, createRequest.TicketPartitionKey, createRequest.TicketRangeKey,
			callResponse.Data.PartitionKey, callResponse.Data.RangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, TicketPK: %s",
				methodName, genericError.StatusCode, genericError.Message, createRequest.UserId, createRequest.TicketPartitionKey)
			return response.Response[model2.TicketCommentModel]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, TicketPK: %s",
			methodName, *networkError, createRequest.UserId, createRequest.TicketPartitionKey)
		return response.Response[model2.TicketCommentModel]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, TicketPK: %s, TicketRK: %s, CommentPK: %s, CommentRK: %s",
		methodName, createRequest.UserId, createRequest.TicketPartitionKey, createRequest.TicketRangeKey,
		networkResponse.PartitionKey, networkResponse.RangeKey)
	return response.Response[model2.TicketCommentModel]{Data: networkResponse, StatusCode: 200}
}

func (commentService *TicketCommentService) FetchAll(fetchRequest ticket_comment_request.TicketCommentModelFetchAllRequest) response.Response[model2.TicketCommentModelsResponse] {
	methodName := "TicketCommentService.FetchAll"
	lastRangeKeyStr := "nil"
	if fetchRequest.LastRangeKey != nil {
		lastRangeKeyStr = *fetchRequest.LastRangeKey
	}
	log.Printf("%s - STARTED - UserId: %s, TicketPK: %s, TicketRK: %s, LastRangeKey: %s",
		methodName, fetchRequest.UserId, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey, lastRangeKeyStr)

	params := map[string]string{
		"controller": commentService.controllerName,
		"action":     "fetchAll",
	}
	manager := network2.ProvideNetworkManager[model2.TicketCommentModelsResponse](commentService.endpoint, params, &commentService.apiKey, &commentService.contentType)

	bytes, parseError := json2.Marshal(fetchRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, TicketPK: %s",
			methodName, parseError, fetchRequest.UserId, fetchRequest.TicketPartitionKey)
		return response.Response[model2.TicketCommentModelsResponse]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, commentService.metricsManager, func() (*model2.TicketCommentModelsResponse, *error) {
		callResponse := network2.Post[model2.TicketCommentModelsResponse](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, fetchRequest.UserId, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, TicketPK: %s",
				methodName, callResponse.StatusCode, errorMsg, fetchRequest.UserId, fetchRequest.TicketPartitionKey)
			return nil, callResponse.Error
		}

		resultCount := 0
		if callResponse.Data != nil && callResponse.Data.Results != nil {
			resultCount = len(callResponse.Data.Results)
		}
		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s, CommentCount: %d",
			methodName, callResponse.StatusCode, fetchRequest.UserId, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey, resultCount)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, TicketPK: %s",
				methodName, genericError.StatusCode, genericError.Message, fetchRequest.UserId, fetchRequest.TicketPartitionKey)
			return response.Response[model2.TicketCommentModelsResponse]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, TicketPK: %s",
			methodName, *networkError, fetchRequest.UserId, fetchRequest.TicketPartitionKey)
		return response.Response[model2.TicketCommentModelsResponse]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	resultCount := 0
	if networkResponse != nil && networkResponse.Results != nil {
		resultCount = len(networkResponse.Results)
	}
	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, TicketPK: %s, TicketRK: %s, CommentCount: %d",
		methodName, fetchRequest.UserId, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey, resultCount)
	return response.Response[model2.TicketCommentModelsResponse]{Data: networkResponse, StatusCode: 200}
}

func (commentService *TicketCommentService) Fetch(partitionKey string, rangeKey string, userId string) response.Response[model2.TicketCommentModel] {
	methodName := "TicketCommentService.Fetch"
	log.Printf("%s - STARTED - UserId: %s, CommentPK: %s, CommentRK: %s", methodName, userId, partitionKey, rangeKey)

	params := map[string]string{
		"controller": commentService.controllerName,
		"action":     "fetch",
	}
	manager := network2.ProvideNetworkManager[model2.TicketCommentModel](commentService.endpoint, params, &commentService.apiKey, &commentService.contentType)

	fetchRequest := model2.FetchRequest{
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		UserId:       userId,
	}

	bytes, parseError := json2.Marshal(fetchRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, CommentPK: %s",
			methodName, parseError, userId, partitionKey)
		return response.Response[model2.TicketCommentModel]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, commentService.metricsManager, func() (*model2.TicketCommentModel, *error) {
		callResponse := network2.Post[model2.TicketCommentModel](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, CommentPK: %s, CommentRK: %s",
			methodName, callResponse.StatusCode, userId, partitionKey, rangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, CommentPK: %s",
				methodName, callResponse.StatusCode, errorMsg, userId, partitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, CommentPK: %s, CommentRK: %s",
			methodName, callResponse.StatusCode, userId, partitionKey, rangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, CommentPK: %s",
				methodName, genericError.StatusCode, genericError.Message, userId, partitionKey)
			return response.Response[model2.TicketCommentModel]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, CommentPK: %s",
			methodName, *networkError, userId, partitionKey)
		return response.Response[model2.TicketCommentModel]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, CommentPK: %s, CommentRK: %s",
		methodName, userId, partitionKey, rangeKey)
	return response.Response[model2.TicketCommentModel]{Data: networkResponse, StatusCode: 200}
}

func (commentService *TicketCommentService) Update(updateRequest ticket_comment_request.TicketCommentModelUpdateRequest) response.Response[bool] {
	methodName := "TicketCommentService.Update"
	log.Printf("%s - STARTED - UserId: %s, CommentPK: %s, CommentRK: %s",
		methodName, updateRequest.UserId, updateRequest.Comment.PartitionKey, updateRequest.Comment.RangeKey)

	params := map[string]string{
		"controller": commentService.controllerName,
		"action":     "update",
	}
	manager := network2.ProvideNetworkManager[bool](commentService.endpoint, params, &commentService.apiKey, &commentService.contentType)

	bytes, parseError := json2.Marshal(updateRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, CommentPK: %s",
			methodName, parseError, updateRequest.UserId, updateRequest.Comment.PartitionKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, commentService.metricsManager, func() (*bool, *error) {
		callResponse := network2.Post[bool](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, CommentPK: %s, CommentRK: %s",
			methodName, callResponse.StatusCode, updateRequest.UserId, updateRequest.Comment.PartitionKey, updateRequest.Comment.RangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, CommentPK: %s",
				methodName, callResponse.StatusCode, errorMsg, updateRequest.UserId, updateRequest.Comment.PartitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, CommentPK: %s, CommentRK: %s",
			methodName, callResponse.StatusCode, updateRequest.UserId, updateRequest.Comment.PartitionKey, updateRequest.Comment.RangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, CommentPK: %s",
				methodName, genericError.StatusCode, genericError.Message, updateRequest.UserId, updateRequest.Comment.PartitionKey)
			return response.Response[bool]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, CommentPK: %s",
			methodName, *networkError, updateRequest.UserId, updateRequest.Comment.PartitionKey)
		return response.Response[bool]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, CommentPK: %s, CommentRK: %s",
		methodName, updateRequest.UserId, updateRequest.Comment.PartitionKey, updateRequest.Comment.RangeKey)
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (commentService *TicketCommentService) Delete(deleteRequest model2.DeleteRequest) response.Response[bool] {
	methodName := "TicketCommentService.Delete"
	log.Printf("%s - STARTED - UserId: %s, CommentPK: %s, CommentRK: %s",
		methodName, deleteRequest.UserId, deleteRequest.PartitionKey, deleteRequest.RangeKey)

	params := map[string]string{
		"controller": commentService.controllerName,
		"action":     "delete",
	}
	manager := network2.ProvideNetworkManager[bool](commentService.endpoint, params, &commentService.apiKey, &commentService.contentType)

	bytes, parseError := json2.Marshal(deleteRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, CommentPK: %s",
			methodName, parseError, deleteRequest.UserId, deleteRequest.PartitionKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, commentService.metricsManager, func() (*bool, *error) {
		callResponse := network2.Post[bool](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, CommentPK: %s, CommentRK: %s",
			methodName, callResponse.StatusCode, deleteRequest.UserId, deleteRequest.PartitionKey, deleteRequest.RangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, CommentPK: %s",
				methodName, callResponse.StatusCode, errorMsg, deleteRequest.UserId, deleteRequest.PartitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, CommentPK: %s, CommentRK: %s",
			methodName, callResponse.StatusCode, deleteRequest.UserId, deleteRequest.PartitionKey, deleteRequest.RangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, CommentPK: %s",
				methodName, genericError.StatusCode, genericError.Message, deleteRequest.UserId, deleteRequest.PartitionKey)
			return response.Response[bool]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, CommentPK: %s",
			methodName, *networkError, deleteRequest.UserId, deleteRequest.PartitionKey)
		return response.Response[bool]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, CommentPK: %s, CommentRK: %s",
		methodName, deleteRequest.UserId, deleteRequest.PartitionKey, deleteRequest.RangeKey)
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (commentService *TicketCommentService) FetchByUser(fetchRequest ticket_comment_request.TicketCommentModelByUserRequest) response.Response[model2.TicketCommentModelsResponse] {
	methodName := "TicketCommentService.FetchByUser"
	lastRangeKeyStr := "nil"
	if fetchRequest.LastRangeKey != nil {
		lastRangeKeyStr = *fetchRequest.LastRangeKey
	}
	log.Printf("%s - STARTED - UserId: %s, LastRangeKey: %s", methodName, fetchRequest.UserId, lastRangeKeyStr)

	params := map[string]string{
		"controller": commentService.controllerName,
		"action":     "fetchByUser",
	}
	manager := network2.ProvideNetworkManager[model2.TicketCommentModelsResponse](commentService.endpoint, params, &commentService.apiKey, &commentService.contentType)

	bytes, parseError := json2.Marshal(fetchRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s",
			methodName, parseError, fetchRequest.UserId)
		return response.Response[model2.TicketCommentModelsResponse]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, commentService.metricsManager, func() (*model2.TicketCommentModelsResponse, *error) {
		callResponse := network2.Post[model2.TicketCommentModelsResponse](manager, bytes)

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
		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, CommentCount: %d",
			methodName, callResponse.StatusCode, fetchRequest.UserId, resultCount)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s",
				methodName, genericError.StatusCode, genericError.Message, fetchRequest.UserId)
			return response.Response[model2.TicketCommentModelsResponse]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s", methodName, *networkError, fetchRequest.UserId)
		return response.Response[model2.TicketCommentModelsResponse]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	resultCount := 0
	if networkResponse != nil && networkResponse.Results != nil {
		resultCount = len(networkResponse.Results)
	}
	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, CommentCount: %d",
		methodName, fetchRequest.UserId, resultCount)
	return response.Response[model2.TicketCommentModelsResponse]{Data: networkResponse, StatusCode: 200}
}
