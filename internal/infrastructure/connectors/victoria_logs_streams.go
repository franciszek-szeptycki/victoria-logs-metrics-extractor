package connectors

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/application/dtos"
	"main/internal/constants"
	"regexp"

	"github.com/go-resty/resty/v2"
)

type VictoriaLogsStreamsConnector struct {
	url                 string
	logTimeframeMinutes int
}

func NewVictoriaLogsStreamsConnector(url string, logTimeframeMinutes int) *VictoriaLogsStreamsConnector {
	return &VictoriaLogsStreamsConnector{
		url:                 url,
		logTimeframeMinutes: logTimeframeMinutes,
	}
}

func (v *VictoriaLogsStreamsConnector) FetchStreams(query string) ([]dtos.LogStreamDTO, error) {
	fullURL := fmt.Sprintf("%s%s", v.url, constants.StreamsPath)

	payload := map[string]string{
		"query": query,
		"start": fmt.Sprintf("%dm", v.logTimeframeMinutes),
	}

	response, err := makePostRequest(fullURL, payload)
	if err != nil {
		log.Fatalf("Error fetching streams: %s", err)
		return []dtos.LogStreamDTO{}, err
	}

	return mapStreamResponseToDTO(response)
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

func mapStreamResponseToDTO(response string) ([]dtos.LogStreamDTO, error) {
	type ValueItem struct {
		Value string `json:"value"`
		Hits  int    `json:"hits"`
	}
	type InputData struct {
		Values []ValueItem `json:"values"`
	}

	var input InputData
	if err := json.Unmarshal([]byte(response), &input); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	re := regexp.MustCompile(`kubernetes\.container_name="([^"]+)",kubernetes\.pod_namespace="([^"]+)"`)

	var results []dtos.LogStreamDTO
	for _, item := range input.Values {
		matches := re.FindStringSubmatch(item.Value)
		if len(matches) == 3 {
			results = append(results, dtos.LogStreamDTO{
				KubernetesContainerName: matches[1],
				KubernetesNamespace:     matches[2],
				Hits:                    item.Hits,
			})
		}
	}

	return results, nil
}
