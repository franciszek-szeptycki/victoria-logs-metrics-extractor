package factories

import (
	"main/internal/application/use_cases"
	"main/internal/infrastructure/config"
	"main/internal/infrastructure/connectors"
)

type AnalyzeMetricsFactory struct{}

func NewAnalyzeMetricsFactory() *AnalyzeMetricsFactory {
	return &AnalyzeMetricsFactory{}
}

func (f *AnalyzeMetricsFactory) Execute() *AnalyzeMetrics {

	cfg := config.LoadConfig()

	connector := connectors.NewVictoriaLogsStreamsConnector(cfg.VictoriaLogsURL, cfg.LogTimeframeMinutes)

	convertLogsToMetricsUseCase := use_cases.ConvertLogsToMetricsUseCase{
		victoriaLogsStreamsConnector: connector,
	}

}
