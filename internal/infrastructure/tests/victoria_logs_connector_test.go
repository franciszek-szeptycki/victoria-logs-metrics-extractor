package tests

import (
	"main/internal/application/constants"
	"main/internal/application/selectors"
	"main/internal/infrastructure/connectors"
	"strings"
	"testing"
)

func TestVictoriaLogsConnector(t *testing.T) {

	connector := connectors.NewVictoriaLogsConnector()

	cfg := selectors.Config{
		VictoriaLogsURL:     "http://localhost:9428",
		LogTimeframeMinutes: 5,
		ErrorThreshold:      0.5,
	}

	query := constants.AllStreamsHitsQuery
	output := connector.FetchStreams(cfg, query)

	if len(output.Values) == 0 {
		t.Errorf("Expected streams to be returned, got: %v", output)
	}

	item := output.Values[0].Value
	phrases := []string{"kubernetes.pod_namespace", "kubernetes.container_name"}

	for _, phrase := range phrases {
		if !strings.Contains(item, phrase) {
			t.Errorf("Expected item to contain %s", phrase)
		}
	}
}
