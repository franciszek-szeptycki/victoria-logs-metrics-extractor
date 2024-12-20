package main

import (
	"main/internal/infrastructure/config"
	"main/internal/infrastructure/factories"
)

func main() {
	cfg := config.Init()

	factory := factories.NewConvertLogsToMetricsFactory()

	useCase := factory.Execute()

	useCase.Execute(cfg)
}
