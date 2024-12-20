package factories

import (
	"main/internal/application/services"
	"main/internal/application/use_cases"
	"main/internal/infrastructure/connectors"
	"main/internal/infrastructure/presenters"
)

type ConvertLogsToMetricsFactory struct{}

func NewConvertLogsToMetricsFactory() *ConvertLogsToMetricsFactory {
	return &ConvertLogsToMetricsFactory{}
}

func (f *ConvertLogsToMetricsFactory) Execute() *use_cases.ConvertLogsToMetricsUseCase {
	// Create instances of the required components
	connector := connectors.NewVictoriaLogsStreamsConnector()
	mapStreamsResponseToLogs := &services.FetchLogsStreamsMapper{}
	analyzeLogStreamsService := &services.AnalyzeLogStreamsService{}
	jsonPresenter := &presenters.JSONPresenter{}

	// Return the new use case
	return use_cases.NewConvertLogsToMetricsUseCase(connector, mapStreamsResponseToLogs, analyzeLogStreamsService, jsonPresenter)
}
