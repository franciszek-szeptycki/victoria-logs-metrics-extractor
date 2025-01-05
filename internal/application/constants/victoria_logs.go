package constants

const (
	LogsQLQueryAllStreams      = "*"
	LogsQLQueryPositiveStreams = "stream:stdout"
)

const (
	VictoriaLogsApiPathStreams = "/select/logsql/streams"
	VictoriaLogsApiPathQuery   = "/select/logsql/query"
)

const DefaultErrorThreshold = 0.01
