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

func NewConvertLogsToMetricsUseCase(
	victoriaLogsConnector *connectors.VictoriaLogsStreamsConnector,
	analyzeLogStreamsService *services.AnalyzeLogStreamsService,
	jsonPresenter *presenters.JSONPresenter,
) *ConvertLogsToMetricsUseCase {
	return &ConvertLogsToMetricsUseCase{
		victoriaLogsConnector:    victoriaLogsConnector,
		analyzeLogStreamsService: analyzeLogStreamsService,
		jsonPresenter:            jsonPresenter,
	}
}

func (c *ConvertLogsToMetricsUseCase) Execute(cfg config.Config) {
	allstreams, err := c.victoriaLogsConnector.FetchStreams(cfg, constants.AllStreamsHitsQuery)
	if err != nil {
		log.Fatalln(err)
		return
	}
	positivestreams, err := c.victoriaLogsConnector.FetchStreams(cfg, constants.PositiveHitsQuery)
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
