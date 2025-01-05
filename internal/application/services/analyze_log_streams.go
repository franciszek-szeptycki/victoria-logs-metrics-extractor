package services

import "main/internal/application/selectors"

type AnalyzeMetricsService struct{}

func (a *AnalyzeMetricsService) Execute(resourceMetricsWithThresholds []selectors.ResourceMetricsWithErrorThresholdDTO) []selectors.MetricsOutputDTO {

	outputArray := []selectors.MetricsOutputDTO{}
	for _, dto := range resourceMetricsWithThresholds {
		outputArray = append(outputArray, selectors.MetricsOutputDTO{
			Resource:       dto.Resource,
			All:            dto.AllHits,
			Succeded:       dto.PositiveHits,
			Errors:         a.calculateErrors(dto),
			ErrorRate:      a.calculateErrorRate(dto),
			HealthScore:    a.calculateHealthScore(dto),
			IsHealthy:      a.checkIsHealthy(dto),
			ErrorThreshold: dto.ErrorThreshold,
		})
	}

	return outputArray
}

func (a *AnalyzeMetricsService) calculateErrors(dto selectors.ResourceMetricsWithErrorThresholdDTO) int {
	return dto.AllHits - dto.PositiveHits
}

func (a *AnalyzeMetricsService) calculateErrorRate(dto selectors.ResourceMetricsWithErrorThresholdDTO) float32 {
	return float32(a.calculateErrors(dto)) / float32(dto.AllHits)
}

func (a *AnalyzeMetricsService) calculateHealthScore(dto selectors.ResourceMetricsWithErrorThresholdDTO) float32 {
	return 1 - a.calculateErrorRate(dto)
}

func (a *AnalyzeMetricsService) checkIsHealthy(dto selectors.ResourceMetricsWithErrorThresholdDTO) int {
	if a.calculateErrorRate(dto) <= dto.ErrorThreshold {
		return 1
	}
	return 0
}
