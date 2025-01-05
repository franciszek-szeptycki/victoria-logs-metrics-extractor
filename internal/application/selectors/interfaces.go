package selectors

type VictoriaLogsConnector interface {
	FetchStreams(query string) FetchStreamsResponse
	FetchLastLog(query string) LastLogReponse
}

type PresenterInterface interface {
	Present(output []MetricsOutputDTO)
}
