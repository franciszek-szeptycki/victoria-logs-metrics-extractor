package main

import (
	"main/internal/infrastructure/factories"
)

func main() {
	factory := factories.NewConvertLogsToMetricsFactory()

	useCase := factory.Execute()

	useCase.Execute()
}
