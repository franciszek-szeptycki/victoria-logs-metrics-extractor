package connectors

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/application/constants"
	"main/internal/application/selectors"

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

type VictoriaLogsConnector struct{}

func NewVictoriaLogsConnector() *VictoriaLogsConnector {
	return &VictoriaLogsConnector{}
}

func (v *VictoriaLogsConnector) FetchStreams(cfg selectors.Config, query string) selectors.FetchStreamsResponse {
	fullURL := fmt.Sprintf("%s%s", cfg.VictoriaLogsURL, constants.VictoriaLogsApiPathStreams)
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

func (v *VictoriaLogsConnector) post(httpRequest httpRequest) httpResponse {
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

func (v *VictoriaLogsConnector) FetchLastLog(cfg selectors.Config, logStreamDTO selectors.LogStreamDTO) {
	fullURL := fmt.Sprintf("%s%s", cfg.VictoriaLogsURL, constants.VictoriaLogsApiPathQuery)

	query := fmt.Sprintf("kubernetes.pod_namespace:%s AND kubernetes.container_name:%s", logStreamDTO.KubernetesNamespace, logStreamDTO.KubernetesContainerName)
	payload := map[string]string{
		"query": query,
		"limit": "1",
	}

	httpResponse := v.post(httpRequest{
		URL:  fullURL,
		Body: payload,
	})

	if httpResponse.Status != 200 {
		log.Fatalf("Error fetching logs: %s", httpResponse.Body)
	}

	var logsResponse selectors.FetchStreamsResponse
	err := json.Unmarshal([]byte(httpResponse.Body), &logsResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling logs response: %s", err)
	}
}
