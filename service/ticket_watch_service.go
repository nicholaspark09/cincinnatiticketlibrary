package service

import (
	json2 "encoding/json"
	"errors"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	network2 "github.com/nicholaspark09/awsgorocket/network"
	"github.com/nicholaspark09/awsgorocket/utils"
	model2 "github.com/nicholaspark09/cincinnatiticketlibrary/model"
	"github.com/nicholaspark09/cincinnatiticketlibrary/model/ticket_watch_request"
	"log"
)

type TicketWatchService struct {
	endpoint       string
	apiKey         string
	contentType    string
	controllerName string
	metricsManager metrics2.MetricsManagerContract
}

func ProvideTicketWatchService(
	endpoint string,
	apiKey string,
	metricsManager metrics2.MetricsManagerContract,
) TicketWatchService {
	return TicketWatchService{
		endpoint:       endpoint,
		apiKey:         apiKey,
		contentType:    "application/json",
		controllerName: "watchers",
		metricsManager: metricsManager,
	}
}

func (watchService *TicketWatchService) AddWatcher(addRequest ticket_watch_request.TicketWatchAddRequest) response.Response[model2.TicketWatchModel] {
	methodName := "TicketWatchService.AddWatcher"
	log.Printf("%s - STARTED - UserId: %s, TicketPK: %s, TicketRK: %s",
		methodName, addRequest.UserId, addRequest.TicketPartitionKey, addRequest.TicketRangeKey)

	params := map[string]string{
		"controller": watchService.controllerName,
		"action":     "addWatcher",
	}
	manager := network2.ProvideNetworkManager[model2.TicketWatchModel](watchService.endpoint, params, &watchService.apiKey, &watchService.contentType)

	bytes, parseError := json2.Marshal(addRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, TicketPK: %s",
			methodName, parseError, addRequest.UserId, addRequest.TicketPartitionKey)
		return response.Response[model2.TicketWatchModel]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, watchService.metricsManager, func() (*model2.TicketWatchModel, *error) {
		callResponse := network2.Post[model2.TicketWatchModel](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, addRequest.UserId, addRequest.TicketPartitionKey, addRequest.TicketRangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, TicketPK: %s",
				methodName, callResponse.StatusCode, errorMsg, addRequest.UserId, addRequest.TicketPartitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, addRequest.UserId, addRequest.TicketPartitionKey, addRequest.TicketRangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, TicketPK: %s",
				methodName, genericError.StatusCode, genericError.Message, addRequest.UserId, addRequest.TicketPartitionKey)
			return response.Response[model2.TicketWatchModel]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, TicketPK: %s",
			methodName, *networkError, addRequest.UserId, addRequest.TicketPartitionKey)
		return response.Response[model2.TicketWatchModel]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, TicketPK: %s, TicketRK: %s",
		methodName, addRequest.UserId, addRequest.TicketPartitionKey, addRequest.TicketRangeKey)
	return response.Response[model2.TicketWatchModel]{Data: networkResponse, StatusCode: 200}
}

func (watchService *TicketWatchService) RemoveWatcher(removeRequest ticket_watch_request.TicketWatchRemoveRequest) response.Response[bool] {
	methodName := "TicketWatchService.RemoveWatcher"
	log.Printf("%s - STARTED - UserId: %s, TicketKey: %s",
		methodName, removeRequest.UserId, removeRequest.RangeKey)

	params := map[string]string{
		"controller": watchService.controllerName,
		"action":     "removeWatcher",
	}
	manager := network2.ProvideNetworkManager[bool](watchService.endpoint, params, &watchService.apiKey, &watchService.contentType)

	bytes, parseError := json2.Marshal(removeRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, TicketPK: %s",
			methodName, parseError, removeRequest.UserId, removeRequest.RangeKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, watchService.metricsManager, func() (*bool, *error) {
		callResponse := network2.Post[bool](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, TicketPK: %s",
			methodName, callResponse.StatusCode, removeRequest.UserId, removeRequest.RangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, TicketPK: %s",
				methodName, callResponse.StatusCode, errorMsg, removeRequest.UserId, removeRequest.RangeKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, TicketPK: %s",
			methodName, callResponse.StatusCode, removeRequest.UserId, removeRequest.RangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, TicketPK: %s",
				methodName, genericError.StatusCode, genericError.Message, removeRequest.UserId, removeRequest.RangeKey)
			return response.Response[bool]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, TicketPK: %s",
			methodName, *networkError, removeRequest.UserId, removeRequest.RangeKey)
		return response.Response[bool]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, TicketPK: %s",
		methodName, removeRequest.UserId, removeRequest.RangeKey)
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (watchService *TicketWatchService) GetUserWatchList(fetchRequest ticket_watch_request.TicketWatchUserListRequest) response.Response[model2.TicketWatchModelsResponse] {
	methodName := "TicketWatchService.GetUserWatchList"
	log.Printf("%s - STARTED - UserId: %s, LastRangeKey: %s", methodName, fetchRequest.UserId, fetchRequest.LastRangeKey)

	params := map[string]string{
		"controller": watchService.controllerName,
		"action":     "getUserWatchList",
		"userId":     fetchRequest.UserId,
	}
	if fetchRequest.LastRangeKey != nil && len(*fetchRequest.LastRangeKey) > 0 {
		params["lastRangeKey"] = *fetchRequest.LastRangeKey
	}
	manager := network2.ProvideNetworkManager[model2.TicketWatchModelsResponse](watchService.endpoint, params, &watchService.apiKey, &watchService.contentType)

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, watchService.metricsManager, func() (*model2.TicketWatchModelsResponse, *error) {
		callResponse := network2.Get[model2.TicketWatchModelsResponse](manager)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, LastRangeKey: %s",
			methodName, callResponse.StatusCode, fetchRequest.UserId, fetchRequest.LastRangeKey)

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
			return response.Response[model2.TicketWatchModelsResponse]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s", methodName, *networkError, fetchRequest.UserId)
		return response.Response[model2.TicketWatchModelsResponse]{
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
	return response.Response[model2.TicketWatchModelsResponse]{Data: networkResponse, StatusCode: 200}
}

func (watchService *TicketWatchService) GetUserUnreadList(fetchRequest ticket_watch_request.TicketWatchUserListRequest) response.Response[model2.TicketWatchModelsResponse] {
	methodName := "TicketWatchService.GetUserUnreadList"
	log.Printf("%s - STARTED - UserId: %s, LastRangeKey: %s", methodName, fetchRequest.UserId, fetchRequest.LastRangeKey)

	params := map[string]string{
		"controller": watchService.controllerName,
		"action":     "getUserUnreadList",
		"userId":     fetchRequest.UserId,
	}
	manager := network2.ProvideNetworkManager[model2.TicketWatchModelsResponse](watchService.endpoint, params, &watchService.apiKey, &watchService.contentType)

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, watchService.metricsManager, func() (*model2.TicketWatchModelsResponse, *error) {
		callResponse := network2.Get[model2.TicketWatchModelsResponse](manager)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, LastRangeKey: %s",
			methodName, callResponse.StatusCode, fetchRequest.UserId, fetchRequest.LastRangeKey)

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
		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, UnreadCount: %d",
			methodName, callResponse.StatusCode, fetchRequest.UserId, resultCount)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s",
				methodName, genericError.StatusCode, genericError.Message, fetchRequest.UserId)
			return response.Response[model2.TicketWatchModelsResponse]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s", methodName, *networkError, fetchRequest.UserId)
		return response.Response[model2.TicketWatchModelsResponse]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	resultCount := 0
	if networkResponse != nil && networkResponse.Results != nil {
		resultCount = len(networkResponse.Results)
	}
	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, UnreadCount: %d",
		methodName, fetchRequest.UserId, resultCount)
	return response.Response[model2.TicketWatchModelsResponse]{Data: networkResponse, StatusCode: 200}
}

func (watchService *TicketWatchService) GetTicketWatchers(fetchRequest ticket_watch_request.TicketWatchersListRequest) response.Response[model2.TicketWatchModelsResponse] {
	methodName := "TicketWatchService.GetTicketWatchers"
	log.Printf("%s - STARTED - TicketPK: %s, TicketRK: %s",
		methodName, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey)

	params := map[string]string{
		"controller": watchService.controllerName,
		"action":     "getTicketWatchers",
		"ticketPK":   fetchRequest.TicketPartitionKey,
		"ticketRK":   fetchRequest.TicketRangeKey,
		"userId":     fetchRequest.UserId,
	}
	if fetchRequest.LastPartitionKey != nil && fetchRequest.LastRangeKey != nil && len(*fetchRequest.LastPartitionKey) > 0 {
		params["lastPartitionKey"] = *fetchRequest.LastPartitionKey
		params["lastRangeKey"] = *fetchRequest.LastRangeKey
	}
	manager := network2.ProvideNetworkManager[model2.TicketWatchModelsResponse](watchService.endpoint, params, &watchService.apiKey, &watchService.contentType)

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, watchService.metricsManager, func() (*model2.TicketWatchModelsResponse, *error) {
		callResponse := network2.Get[model2.TicketWatchModelsResponse](manager)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, TicketPK: %s, TicketRK: %s",
				methodName, callResponse.StatusCode, errorMsg, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey)
			return nil, callResponse.Error
		}

		resultCount := 0
		if callResponse.Data != nil && callResponse.Data.Results != nil {
			resultCount = len(callResponse.Data.Results)
		}
		log.Printf("%s - SUCCESS - StatusCode: %d, TicketPK: %s, TicketRK: %s, WatcherCount: %d",
			methodName, callResponse.StatusCode, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey, resultCount)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, TicketPK: %s, TicketRK: %s",
				methodName, genericError.StatusCode, genericError.Message, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey)
			return response.Response[model2.TicketWatchModelsResponse]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, TicketPK: %s, TicketRK: %s",
			methodName, *networkError, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey)
		return response.Response[model2.TicketWatchModelsResponse]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	resultCount := 0
	if networkResponse != nil && networkResponse.Results != nil {
		resultCount = len(networkResponse.Results)
	}
	log.Printf("%s - COMPLETED - StatusCode: 200, TicketPK: %s, TicketRK: %s, WatcherCount: %d",
		methodName, fetchRequest.TicketPartitionKey, fetchRequest.TicketRangeKey, resultCount)
	return response.Response[model2.TicketWatchModelsResponse]{Data: networkResponse, StatusCode: 200}
}

func (watchService *TicketWatchService) MarkAsRead(markReadRequest ticket_watch_request.TicketWatchMarkReadRequest) response.Response[bool] {
	methodName := "TicketWatchService.MarkAsRead"
	log.Printf("%s - STARTED - UserId: %s, TicketPK: %s, TicketRK: %s",
		methodName, markReadRequest.UserId, markReadRequest.TicketPartitionKey, markReadRequest.TicketRangeKey)

	params := map[string]string{
		"controller": watchService.controllerName,
		"action":     "markAsRead",
	}
	manager := network2.ProvideNetworkManager[bool](watchService.endpoint, params, &watchService.apiKey, &watchService.contentType)

	bytes, parseError := json2.Marshal(markReadRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, TicketPK: %s",
			methodName, parseError, markReadRequest.UserId, markReadRequest.TicketPartitionKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, watchService.metricsManager, func() (*bool, *error) {
		callResponse := network2.Post[bool](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, markReadRequest.UserId, markReadRequest.TicketPartitionKey, markReadRequest.TicketRangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, TicketPK: %s",
				methodName, callResponse.StatusCode, errorMsg, markReadRequest.UserId, markReadRequest.TicketPartitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, markReadRequest.UserId, markReadRequest.TicketPartitionKey, markReadRequest.TicketRangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, TicketPK: %s",
				methodName, genericError.StatusCode, genericError.Message, markReadRequest.UserId, markReadRequest.TicketPartitionKey)
			return response.Response[bool]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, TicketPK: %s",
			methodName, *networkError, markReadRequest.UserId, markReadRequest.TicketPartitionKey)
		return response.Response[bool]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, TicketPK: %s, TicketRK: %s",
		methodName, markReadRequest.UserId, markReadRequest.TicketPartitionKey, markReadRequest.TicketRangeKey)
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (watchService *TicketWatchService) UpdateWatchEntry(updateRequest ticket_watch_request.TicketWatchUpdateRequest) response.Response[bool] {
	methodName := "TicketWatchService.UpdateWatchEntry"
	log.Printf("%s - STARTED - UserId: %s, TicketPK: %s, TicketRK: %s",
		methodName, updateRequest.UserId, updateRequest.TicketPartitionKey, updateRequest.TicketRangeKey)

	params := map[string]string{
		"controller": watchService.controllerName,
		"action":     "updateWatchEntry",
	}
	manager := network2.ProvideNetworkManager[bool](watchService.endpoint, params, &watchService.apiKey, &watchService.contentType)

	bytes, parseError := json2.Marshal(updateRequest)
	if parseError != nil {
		log.Printf("%s - PARSE_ERROR - Failed to marshal request: %v, UserId: %s, TicketPK: %s",
			methodName, parseError, updateRequest.UserId, updateRequest.TicketPartitionKey)
		return response.Response[bool]{StatusCode: 400, Message: "Invalid request body"}
	}

	networkResponse, networkError := metrics2.MeasureTimeWithError(methodName, watchService.metricsManager, func() (*bool, *error) {
		callResponse := network2.Post[bool](manager, bytes)

		log.Printf("%s - NETWORK_RESPONSE - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, updateRequest.UserId, updateRequest.TicketPartitionKey, updateRequest.TicketRangeKey)

		if callResponse.StatusCode != 200 {
			errorMsg := "Unknown error"
			if callResponse.Error != nil {
				errorMsg = (*callResponse.Error).Error()
			}
			log.Printf("%s - NETWORK_ERROR - StatusCode: %d, Error: %s, UserId: %s, TicketPK: %s",
				methodName, callResponse.StatusCode, errorMsg, updateRequest.UserId, updateRequest.TicketPartitionKey)
			return nil, callResponse.Error
		}

		log.Printf("%s - SUCCESS - StatusCode: %d, UserId: %s, TicketPK: %s, TicketRK: %s",
			methodName, callResponse.StatusCode, updateRequest.UserId, updateRequest.TicketPartitionKey, updateRequest.TicketRangeKey)
		return callResponse.Data, nil
	})

	if networkError != nil {
		var genericError utils.GenericError
		if errors.As(*networkError, &genericError) {
			log.Printf("%s - GENERIC_ERROR - StatusCode: %d, Message: %s, UserId: %s, TicketPK: %s",
				methodName, genericError.StatusCode, genericError.Message, updateRequest.UserId, updateRequest.TicketPartitionKey)
			return response.Response[bool]{
				StatusCode: genericError.StatusCode,
				Message:    genericError.Message,
			}
		}

		log.Printf("%s - UNKNOWN_ERROR - Error: %v, UserId: %s, TicketPK: %s",
			methodName, *networkError, updateRequest.UserId, updateRequest.TicketPartitionKey)
		return response.Response[bool]{
			StatusCode: 500,
			Message:    "Internal service error",
		}
	}

	log.Printf("%s - COMPLETED - StatusCode: 200, UserId: %s, TicketPK: %s, TicketRK: %s",
		methodName, updateRequest.UserId, updateRequest.TicketPartitionKey, updateRequest.TicketRangeKey)
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}
