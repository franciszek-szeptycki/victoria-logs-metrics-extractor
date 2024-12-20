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
	victoriaLogsConnector    *connectors.VictoriaLogsStreamsConnector
	analyzeLogStreamsService *services.AnalyzeLogStreamsService
	jsonPresenter            *presenters.JSONPresenter
}

func NewConvertLogsToMetricsUseCase(victoriaLogsConnector *connectors.VictoriaLogsStreamsConnector, analyzeLogStreamsService *services.AnalyzeLogStreamsService, jsonPresenter *presenters.JSONPresenter) *ConvertLogsToMetricsUseCase {
	return &ConvertLogsToMetricsUseCase{
		victoriaLogsConnector:    victoriaLogsConnector,
		analyzeLogStreamsService: analyzeLogStreamsService,
		jsonPresenter:            jsonPresenter,
	}
}

func (a *ConvertLogsToMetricsUseCase) Execute(cfg config.Config) {
	logTimeframe := cfg.LogTimeframeMinutes
	// errorThreshold := cfg.ErrorThreshold
	positiveHitsQuery := constants.PositiveHitsQuery
	allHitsQuery := constants.AllStreamsHitsQuery

	allstreams, err := a.victoriaLogsConnector.FetchStreams(allHitsQuery, logTimeframe)
	if err != nil {
		log.Fatalln(err)
		return
	}
	positivestreams, err := a.victoriaLogsConnector.FetchStreams(positiveHitsQuery, logTimeframe)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("All streams: ", allstreams)
	log.Println("Positive streams: ", positivestreams)
	// results, err := a.analyzeLogStreamsService.AnalyzeLogStreams(allstreams, positivestreams, errorThreshold)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	// a.jsonPresenter.Present(results)
}
