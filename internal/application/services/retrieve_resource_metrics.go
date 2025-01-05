package services

import (
	"main/internal/application/constants"
	"main/internal/application/selectors"
)

type combinedStreams struct {
	allStreams      selectors.FetchStreamsResponse
	positiveStreams selectors.FetchStreamsResponse
}

type RetrieveResourceMetricsService struct {
	connector selectors.VictoriaLogsConnector
}

func NewRetrieveResourceMetricsService(connector selectors.VictoriaLogsConnector) *RetrieveResourceMetricsService {
	return &RetrieveResourceMetricsService{
		connector: connector,
	}
}

func (r *RetrieveResourceMetricsService) Execute() []selectors.ResourceMetricsDTO {
	allStreams := r.connector.FetchStreams(constants.LogsQLQueryAllStreams)
	positiveStreams := r.connector.FetchStreams(constants.LogsQLQueryPositiveStreams)

	return r.mapToResourceMetricsDTOs(combinedStreams{
		allStreams:      allStreams,
		positiveStreams: positiveStreams,
	})
}

func (r *RetrieveResourceMetricsService) mapToResourceMetricsDTOs(combinedStreams combinedStreams) []selectors.ResourceMetricsDTO {

	allStreamsMap := r.generateStreamsMap(combinedStreams.allStreams)
	positiveStreamsMap := r.generateStreamsMap(combinedStreams.positiveStreams)

	outputArray := []selectors.ResourceMetricsDTO{}
	for key, value := range allStreamsMap {
		positiveHits := 0
		if v, exists := positiveStreamsMap[key]; exists {
			positiveHits = v
		}
		outputArray = append(outputArray, selectors.ResourceMetricsDTO{
			Resource:     key,
			AllHits:      value,
			PositiveHits: positiveHits,
		})
	}
	return outputArray
}

func (r *RetrieveResourceMetricsService) generateStreamsMap(streamResponse selectors.FetchStreamsResponse) map[string]int {
	streamsMap := make(map[string]int)
	for _, item := range streamResponse.Values {
		streamsMap[item.Value] = item.Hits
	}

	return streamsMap
}
