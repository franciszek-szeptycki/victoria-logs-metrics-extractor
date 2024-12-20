package use_cases

import (
	"log"
	"main/internal/application/services"
	"main/internal/constants"
	"main/internal/infrastructure/config"
	"main/internal/infrastructure/connectors"
	"main/internal/infrastructure/presenters"
)

type ConvertLogsToMetricsUseCase struct {
	victoriaLogsConnector    connectors.VictoriaLogsStreamsConnector
	analyzeLogStreamsService services.AnalyzeLogStreamsService
	jsonPresenter            presenters.JSONPresenter
}

func (a *ConvertLogsToMetricsUseCase) Execute(cfg config.Config) {
	allstreams, _ := a.victoriaLogsConnector.FetchStreams(constants.AllStreamsHitsQuery)
	positivestreams, _ := a.victoriaLogsConnector.FetchStreams(constants.PositiveHitsQuery)

	results, err := a.analyzeLogStreamsService.AnalyzeLogStreams(allstreams, positivestreams, cfg.ErrorThreshold)
	if err != nil {
		log.Fatalln(err)
		return
	}

	output.presentjson(results)
}
