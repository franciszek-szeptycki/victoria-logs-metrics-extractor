package operations

import (
	"fmt"
	"main/internal/external"
)

type ResultLogStreamDTO struct {
	ContainerName  string
	Namespace      string
	TotalErrors    int
	Total          int
	HealthScore    float32
	ErrorThreshold float32
}

func AnalyzeLogStreams(allStreams []external.LogStreamDTO, positiveStreams []external.LogStreamDTO, errorThreshold float32) ([]ResultLogStreamDTO, error) {

	results := []ResultLogStreamDTO{}
	for _, stream := range allStreams {
		positiveStream := retrievePositiveLogStreams(positiveStreams, stream)

		containerName := stream.KubernetesContainerName
		namespace := stream.KubernetesNamespace
		totalHits := stream.Hits
		totalErrors := calculateTotalErrors(totalHits, positiveStream.Hits)
		healthScore := calculateHealthScore(totalErrors, totalHits)
		results = append(results, ResultLogStreamDTO{
			ContainerName:  containerName,
			Namespace:      namespace,
			TotalErrors:    totalErrors,
			Total:          totalHits,
			HealthScore:    healthScore,
			ErrorThreshold: errorThreshold,
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

func calculateTotalErrors(total int, positive int) int {
	fmt.Println(total, positive)
	fmt.Println(total - positive)
	return total - positive
}

func calculateHealthScore(totalErrors int, total int) float32 {
	if total != 0 {
		return 1 - (float32(totalErrors) / float32(total))
	}
	return 1
}
