package connectors

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/application/selectors"
	"main/internal/constants"
	"main/internal/infrastructure/config"

	"github.com/go-resty/resty/v2"
)

type httpRequest struct {
	URL  string
	Body map[string]string
}

type httpResponse struct {
	Status int
	Body   string
}

type VictoriaLogsStreamsConnector struct{}

func NewVictoriaLogsStreamsConnector() *VictoriaLogsStreamsConnector {
	return &VictoriaLogsStreamsConnector{}
}

func (v *VictoriaLogsStreamsConnector) FetchStreams(cfg config.Config, query string) selectors.FetchStreamsResponse {
	fullURL := fmt.Sprintf("%s%s", cfg.VictoriaLogsURL, constants.StreamsPath)
	payload := map[string]string{
		"query": query,
		"start": fmt.Sprintf("%dm", cfg.LogTimeframeMinutes),
	}

	httpResponse := v.post(httpRequest{
		URL:  fullURL,
		Body: payload,
	})

	if httpResponse.Status != 200 {
		log.Fatalf("Error fetching streams: %s", httpResponse.Body)
	}

	var streamsResponse selectors.FetchStreamsResponse
	err := json.Unmarshal([]byte(httpResponse.Body), &streamsResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling streams response: %s", err)
	}

	return streamsResponse
}

func (v *VictoriaLogsStreamsConnector) post(httpRequest httpRequest) httpResponse {
	client := resty.New().R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(httpRequest.Body)

	resp, err := client.Post(httpRequest.URL)

	if err != nil {
		log.Fatalf("Error making POST request: %s", err)
	}

	return httpResponse{
		Status: resp.StatusCode(),
		Body:   resp.String(),
	}
}
