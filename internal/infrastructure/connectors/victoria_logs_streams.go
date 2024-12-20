package connectors

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/constants"

	"github.com/go-resty/resty/v2"
)

type VictoriaLogsStreamsConnector struct {
	url string
}

func NewVictoriaLogsStreamsConnector(url string) *VictoriaLogsStreamsConnector {
	return &VictoriaLogsStreamsConnector{
		url: url,
	}
}

type streamsResponseValueDTO struct {
	Value string `json:"value"`
	Hits  int    `json:"hits"`
}

type streamsResponseDTO struct {
	Values []streamsResponseValueDTO `json:"values"`
}

func (v *VictoriaLogsStreamsConnector) FetchStreams(query string, logTimeframeMinute int) (streamsResponseDTO, error) {
	fullURL := fmt.Sprintf("%s%s", v.url, constants.StreamsPath)

	payload := map[string]string{
		"query": query,
		"start": fmt.Sprintf("%dm", logTimeframeMinute),
	}

	response, err := makePostRequest(fullURL, payload)
	if err != nil {
		log.Fatalf("Error fetching streams: %s", err)
		return streamsResponseDTO{}, err
	}

	var streamsResponse streamsResponseDTO
	if err := json.Unmarshal([]byte(response), &streamsResponse); err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
		return streamsResponseDTO{}, err
	}
	log.Println(streamsResponse)
	return streamsResponse, nil
}

func makePostRequest(fullURL string, payload map[string]string) (string, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(payload).
		Post(fullURL)

	if err != nil {
		log.Fatalf("Error making POST request: %s", err)
		return "", err
	}

	return resp.String(), nil
}
