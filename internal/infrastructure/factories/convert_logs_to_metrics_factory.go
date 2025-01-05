package factories

import (
	"main/internal/application/services"
	"main/internal/application/use_cases"
	"main/internal/infrastructure/config"
	"main/internal/infrastructure/connectors"
	"main/internal/infrastructure/presenters"
)

type ConvertLogsToMetricsFactory struct{}

func NewConvertLogsToMetricsFactory() *ConvertLogsToMetricsFactory {
	return &ConvertLogsToMetricsFactory{}
}

func (f *ConvertLogsToMetricsFactory) Execute() *use_cases.ConvertLogsToMetricsUseCase {

	cfg := config.LoadConnectorConfig()

	connector := connectors.NewVictoriaLogsConnector(
		cfg.VictoriaLogsURL,
		cfg.LogTimeframeMinutes,
	)

	retrieveResourceMetricsService := services.NewRetrieveResourceMetricsService(connector)
	retrieveResourceMetricsWithErrorThresholdService := services.NewRetrieveResourceMetricsWithErrorThresholdService(connector)

	analyzeMetricsService := &services.AnalyzeMetricsService{}

	k8sJsonPresenter := &presenters.K8sJsonPresenter{}

	return use_cases.NewConvertLogsToMetricsUseCase(
		retrieveResourceMetricsService,
		retrieveResourceMetricsWithErrorThresholdService,
		analyzeMetricsService,
		k8sJsonPresenter,
	)
}
