package services

import (
	"main/internal/application/dtos"
)

type AnalyzeLogStreamsService struct{}

func NewAnalyzeLogStreamsService() *AnalyzeLogStreamsService {
	return &AnalyzeLogStreamsService{}
}

func (a *AnalyzeLogStreamsService) AnalyzeLogStreams(allStreams []dtos.LogStreamDTO, positiveStreams []dtos.LogStreamDTO, errorThreshold float32) ([]dtos.ResultLogStreamDTO, error) {

	var results []dtos.ResultLogStreamDTO
	for _, stream := range allStreams {
		positiveStream := retrievePositiveLogStreams(positiveStreams, stream)

		containerName := stream.KubernetesContainerName
		namespace := stream.KubernetesNamespace
		totalHits := stream.Hits
		totalErrors := calculateTotalErrors(positiveStream.Hits, totalHits)
		healthScore := calculateHealthScore(totalErrors, totalHits)
		healthy := isHealthy(healthScore, errorThreshold)
		results = append(results, dtos.ResultLogStreamDTO{
			ContainerName:  containerName,
			Namespace:      namespace,
			TotalErrors:    totalErrors,
			Total:          totalHits,
			HealthScore:    healthScore,
			ErrorThreshold: errorThreshold,
			Healthy:        healthy,
		})
	}

	return results, nil
}

func retrievePositiveLogStreams(positiveStreams []dtos.LogStreamDTO, stream dtos.LogStreamDTO) dtos.LogStreamDTO {
	for _, positiveStream := range positiveStreams {
		if positiveStream.KubernetesNamespace == stream.KubernetesNamespace && positiveStream.KubernetesContainerName == stream.KubernetesContainerName {
			return positiveStream
		}
	}
	return dtos.LogStreamDTO{}
}

func calculateHealthScore(totalErrors int, total int) float32 {
	if total != 0 {
		return 1 - (float32(totalErrors) / float32(total))
	}
	return 1
}

func calculateTotalErrors(positiveHits int, totalHits int) int {
	return totalHits - positiveHits
}

func isHealthy(healthScore float32, errorThreshold float32) int {
	if healthScore >= 1-errorThreshold {
		return 1
	}
	return 0
}
