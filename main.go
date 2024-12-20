package main

import (
	"main/internal/infrastructure/config"
	"main/internal/infrastructure/factories"
)

func main() {
	cfg := config.LoadEnv()

	factory := factories.NewConvertLogsToMetricsFactory()

	useCase := factory.Execute()

	useCase.Execute(cfg)
}
