package connectors

import (
	"main/internal/constants"
	"testing"
)

func TestNewVictoriaLogsStreamsConnector(t *testing.T) {
	t.Run("should return a new VictoriaLogsStreamsConnector", func(t *testing.T) {
		got := NewVictoriaLogsStreamsConnector()
		if got == nil {
			t.Errorf("NewVictoriaLogsStreamsConnector() = %v; want a new VictoriaLogsStreamsConnector", got)
		}
	})
}

func TestVictoriaLogsStreamsConnector_FetchStreams(t *testing.T) {
	t.Run("should fetch streams", func(t *testing.T) {

		connector := NewVictoriaLogsStreamsConnector()
		query := constants.AllStreamsHitsQuery
		got, err := connector.FetchStreams(query)
		if err != nil {
			t.Errorf("FetchStreams(%s) = %v; want a list of streams", query, got)
		}
	})
}
