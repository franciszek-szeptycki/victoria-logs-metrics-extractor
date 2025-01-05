package use_cases

import (
	"main/internal/application/selectors"
	"main/internal/application/services"
)

type ConvertLogsToMetricsUseCase struct {
	retrieveMetricsService                           *services.RetrieveResourceMetricsService
	retrieveResourceMetricsWithErrorThresholdService *services.RetrieveResourceMetricsWithErrorThresholdService
	analyzeMetricsService                            *services.AnalyzeMetricsService
	presenter                                        selectors.PresenterInterface
}

func NewConvertLogsToMetricsUseCase(
	retrieveMetricsService *services.RetrieveResourceMetricsService,
	retrieveResourceMetricsWithErrorThresholdService *services.RetrieveResourceMetricsWithErrorThresholdService,
	analyzeMetricsService *services.AnalyzeMetricsService,
	presenter selectors.PresenterInterface,
) *ConvertLogsToMetricsUseCase {
	return &ConvertLogsToMetricsUseCase{
		retrieveMetricsService:                           retrieveMetricsService,
		retrieveResourceMetricsWithErrorThresholdService: retrieveResourceMetricsWithErrorThresholdService,
		analyzeMetricsService:                            analyzeMetricsService,
		presenter:                                        presenter,
	}
}

func (c *ConvertLogsToMetricsUseCase) Execute() {

	resourceMetrics := c.retrieveMetricsService.Execute()

	resourceMetricsWithErrorThreshold := c.retrieveResourceMetricsWithErrorThresholdService.Execute(resourceMetrics)

	output := c.analyzeMetricsService.Execute(resourceMetricsWithErrorThreshold)

	c.presenter.Present(output)
}
