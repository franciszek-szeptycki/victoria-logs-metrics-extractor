package services

import (
	"main/internal/application/constants"
	"main/internal/application/selectors"
	"strconv"
)

type RetrieveResourceMetricsWithErrorThresholdService struct {
	connector selectors.VictoriaLogsConnector
}

func NewRetrieveResourceMetricsWithErrorThresholdService(connector selectors.VictoriaLogsConnector) *RetrieveResourceMetricsWithErrorThresholdService {
	return &RetrieveResourceMetricsWithErrorThresholdService{
		connector: connector,
	}
}

func (r *RetrieveResourceMetricsWithErrorThresholdService) Execute(resourceMetrics []selectors.ResourceMetricsDTO) []selectors.ResourceMetricsWithErrorThresholdDTO {

	outputArray := []selectors.ResourceMetricsWithErrorThresholdDTO{}
	for _, dto := range resourceMetrics {

		errorThreshold := r.getErrorThreshold(dto)

		outputArray = append(outputArray, selectors.ResourceMetricsWithErrorThresholdDTO{
			ResourceMetricsDTO: dto,
			ErrorThreshold:     errorThreshold,
		})
	}
	return outputArray
}

func (r *RetrieveResourceMetricsWithErrorThresholdService) getErrorThreshold(dto selectors.ResourceMetricsDTO) float32 {
	fetchLastLogResponse := r.connector.FetchLastLog(dto.Resource)
	customErrorThreshold := fetchLastLogResponse.CustomErrorThreshold

	if customErrorThreshold == "" {
		return constants.DefaultErrorThreshold
	}

	errorThreshold, err := strconv.ParseFloat(customErrorThreshold, 32)
	if err != nil {
		return constants.DefaultErrorThreshold
	}

	return float32(errorThreshold)
}
