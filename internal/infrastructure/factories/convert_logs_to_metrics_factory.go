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
	// Load configuration
	cfg := config.LoadConfig()

	// Create instances of the required components
	connector := connectors.NewVictoriaLogsStreamsConnector(cfg.VictoriaLogsURL, cfg.LogTimeframeMinutes)
	analyzeLogStreamsService := services.NewAnalyzeLogStreamsService()
	jsonPresenter := presenters.NewJSONPresenter()

	// Return the new use case
	return use_cases.NewConvertLogsToMetricsUseCase(connector, analyzeLogStreamsService, jsonPresenter)
}
