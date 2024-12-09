package operations

import (
	"main/internal/external"
)

type ResultLogStreamDTO struct {
	ContainerName  string
	Namespace      string
	TotalErrors    int
	Total          int
	HealthScore    float32
	ErrorThreshold float32
	Healthy        int
}

func AnalyzeLogStreams(allStreams []external.LogStreamDTO, positiveStreams []external.LogStreamDTO, errorThreshold float32) ([]ResultLogStreamDTO, error) {

	var results []ResultLogStreamDTO
	for _, stream := range allStreams {
		positiveStream := retrievePositiveLogStreams(positiveStreams, stream)

		containerName := stream.KubernetesContainerName
		namespace := stream.KubernetesNamespace
		totalHits := stream.Hits
		totalErrors := calculateTotalErrors(positiveStream.Hits, totalHits)
		healthScore := calculateHealthScore(totalErrors, totalHits)
		healthy := isHealthy(healthScore, errorThreshold)
		results = append(results, ResultLogStreamDTO{
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

func retrievePositiveLogStreams(positiveStreams []external.LogStreamDTO, stream external.LogStreamDTO) external.LogStreamDTO {
	for _, positiveStream := range positiveStreams {
		if positiveStream.KubernetesNamespace == stream.KubernetesNamespace && positiveStream.KubernetesContainerName == stream.KubernetesContainerName {
			return positiveStream
		}
	}
	return external.LogStreamDTO{}
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
