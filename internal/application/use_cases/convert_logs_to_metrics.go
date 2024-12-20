package use_cases

import (
	"log"
	"main/internal/application/constants"
	"main/internal/application/selectors"
	"main/internal/application/services"
	"main/internal/infrastructure/connectors"
	"main/internal/infrastructure/presenters"
)

type ConvertLogsToMetricsUseCase struct {
	victoriaLogsConnector    *connectors.VictoriaLogsStreamsConnector
	mapStreamsResponseToLogs *services.FetchLogsStreamsMapper
	analyzeLogStreamsService *services.AnalyzeLogStreamsService
	jsonPresenter            *presenters.JSONPresenter
}

func NewConvertLogsToMetricsUseCase(
	victoriaLogsConnector *connectors.VictoriaLogsStreamsConnector,
	mapStreamsResponseToLogs *services.FetchLogsStreamsMapper,
	analyzeLogStreamsService *services.AnalyzeLogStreamsService,
	jsonPresenter *presenters.JSONPresenter,
) *ConvertLogsToMetricsUseCase {
	return &ConvertLogsToMetricsUseCase{
		victoriaLogsConnector:    victoriaLogsConnector,
		analyzeLogStreamsService: analyzeLogStreamsService,
		mapStreamsResponseToLogs: mapStreamsResponseToLogs,
		jsonPresenter:            jsonPresenter,
	}
}

func (c *ConvertLogsToMetricsUseCase) Execute(cfg selectors.Config) {
	allstreams := c.victoriaLogsConnector.FetchStreams(cfg, constants.AllStreamsHitsQuery)
	positivestreams := c.victoriaLogsConnector.FetchStreams(cfg, constants.PositiveHitsQuery)

	allStreamsLogs := c.mapStreamsResponseToLogs.MapStreamsResponseToLogs(allstreams)
	positiveStreamsLogs := c.mapStreamsResponseToLogs.MapStreamsResponseToLogs(positivestreams)

	log.Println("All streams: ", allstreams)
	log.Println("Positive streams: ", positivestreams)

	errorThreshold := cfg.ErrorThreshold

	results, err := c.analyzeLogStreamsService.AnalyzeLogStreams(allStreamsLogs, positiveStreamsLogs, errorThreshold)
	if err != nil {
		log.Fatalln(err)
		return
	}

	c.jsonPresenter.Present(results)
}
