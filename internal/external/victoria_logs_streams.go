package external

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"main/internal/config"
	"main/internal/constants"
	"regexp"
)

type LogStreamDTO struct {
	KubernetesNamespace     string `json:"kubernetes.namespace"`
	KubernetesContainerName string `json:"kubernetes.container_name"`
	Hits                    int    `json:"hits"`
}

func FetchStreamsHits(cfg config.Config, query string) ([]LogStreamDTO, error) {
	fullURL := fmt.Sprintf("%s%s", cfg.VictoriaLogsURL, constants.StreamsPath)

	payload := map[string]string{
		"query": query,
		"start": fmt.Sprintf("%dm", cfg.LogTimeframeMinutes),
	}

	response, err := makePostRequest(fullURL, payload)
	if err != nil {
		log.Fatalf("Error fetching streams: %s", err)
		return []LogStreamDTO{}, err
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

func mapStreamResponseToDTO(response string) ([]LogStreamDTO, error) {
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

	var results []LogStreamDTO
	for _, item := range input.Values {
		matches := re.FindStringSubmatch(item.Value)
		if len(matches) == 3 {
			results = append(results, LogStreamDTO{
				KubernetesContainerName: matches[1],
				KubernetesNamespace:     matches[2],
				Hits:                    item.Hits,
			})
		}
	}

	return results, nil
}
