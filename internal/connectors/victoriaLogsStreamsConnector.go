package connectors

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"main/internal/constants"
)

type LogQueryResponse struct {
	KubernetesNamespace     string `json:"kubernetes.namespace"`
	KubernetesContainerName string `json:"kubernetes.container_name"`
	Hits                    int    `json:"hits"`
}

type VictoriaLogsStreamsConnector struct {
	url string
}

func NewVictoriaLogsStreamsConnector(url string) *VictoriaLogsStreamsConnector {
	return &VictoriaLogsStreamsConnector{
		url: url,
	}
}

func (v *VictoriaLogsStreamsConnector) FetchAllStreamsHits() {
	path := constants.StreamsPath
	query := constants.AllStreamsHitsQuery

	response, err := v.makePostRequest(path, query)
	if err != nil {
		fmt.Printf("Error fetching streams: %v\n", err)
		return
	}
	fmt.Printf("Response: %s\n", response)
}

func (v *VictoriaLogsStreamsConnector) FetchPositiveStreamsHits() {
	path := constants.StreamsPath
	query := constants.PositiveHitsQuery

	response, err := v.makePostRequest(path, query)
	if err != nil {
		fmt.Printf("Error fetching streams: %v\n", err)
		return
	}
	fmt.Printf("Response: %s\n", response)
}

func (v *VictoriaLogsStreamsConnector) makePostRequest(path string, query string) (string, error) {
	fullURL := fmt.Sprintf("%s%s", v.url, path)
	fmt.Println("Full URL:", fullURL)

	client := resty.New()

	data := map[string]string{
		"query": query,
		"start": "1m",
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(data).
		Post(fullURL)

	if err != nil {
		fmt.Println("Błąd:", err)
		return "", err
	}

	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Response Body:", resp.String())
	return resp.String(), nil
}
