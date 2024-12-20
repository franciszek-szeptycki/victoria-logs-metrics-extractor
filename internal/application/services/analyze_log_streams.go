package services

import "main/internal/application/selectors"

type AnalyzeLogStreamsService struct{}

func (a *AnalyzeLogStreamsService) AnalyzeLogStreams(allStreams []selectors.LogsStreamsDTO, positiveStreams []selectors.LogsStreamsDTO, errorThreshold float32) ([]selectors.ResultLogStreamDTO, error) {

	var results []selectors.ResultLogStreamDTO
	for _, stream := range allStreams {
		positiveStream := retrievePositiveLogStreams(positiveStreams, stream)

		containerName := stream.KubernetesContainerName
		namespace := stream.KubernetesNamespace
		totalHits := stream.Hits
		totalErrors := calculateTotalErrors(positiveStream.Hits, totalHits)
		healthScore := calculateHealthScore(totalErrors, totalHits)
		healthy := isHealthy(healthScore, errorThreshold)
		results = append(results, selectors.ResultLogStreamDTO{
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

func retrievePositiveLogStreams(positiveStreams []selectors.LogsStreamsDTO, stream selectors.LogsStreamsDTO) selectors.LogsStreamsDTO {
	for _, positiveStream := range positiveStreams {
		if positiveStream.KubernetesNamespace == stream.KubernetesNamespace && positiveStream.KubernetesContainerName == stream.KubernetesContainerName {
			return positiveStream
		}
	}
	return selectors.LogsStreamsDTO{}
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
