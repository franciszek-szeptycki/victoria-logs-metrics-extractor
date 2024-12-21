package tests

import (
	"main/internal/application/selectors"
	"main/internal/infrastructure/connectors"
	"testing"
)

func TestVictoriaLogsConnectorFetchLastLog(t *testing.T) {

	connector := connectors.NewVictoriaLogsConnector()

	cfg := selectors.Config{
		VictoriaLogsURL:     "http://localhost:9428",
		LogTimeframeMinutes: 5,
		ErrorThreshold:      0.5,
	}

	logStream := selectors.LogStreamDTO{
		KubernetesNamespace:     "paris",
		KubernetesContainerName: "mariadb",
		Hits:                    27,
	}

	connector.FetchLastLog(cfg, logStream)

	// t.Error("Not implemented")
}
