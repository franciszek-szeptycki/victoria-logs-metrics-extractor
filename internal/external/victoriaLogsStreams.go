package external

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"main/internal/config"
	"main/internal/constants"
)

type LogQueryResponse struct {
	KubernetesNamespace     string `json:"kubernetes.namespace"`
	KubernetesContainerName string `json:"kubernetes.container_name"`
	Hits                    int    `json:"hits"`
}

func FetchStreamsHits(cfg config.Config, query string) {
	fullURL := fmt.Sprintf("%s%s", cfg.VictoriaLogsURL, constants.StreamsPath)

	payload := map[string]string{
		"query": query,
		"start": fmt.Sprintf("%dm", cfg.LogTimeframeMinutes),
	}

	response, err := makePostRequest(fullURL, payload)
	if err != nil {
		fmt.Printf("Error fetching streams: %v\n", err)
		return
	}
	fmt.Printf("Response: %s\n", response)
}

func makePostRequest(fullURL string, payload map[string]string) (string, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(payload).
		Post(fullURL)

	if err != nil {
		fmt.Println("Błąd:", err)
		return "", err
	}

	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Response Body:", resp.String())
	return resp.String(), nil
}
