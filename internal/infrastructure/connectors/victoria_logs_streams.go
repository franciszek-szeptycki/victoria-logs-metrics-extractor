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

type VictoriaLogsStreamsConnector struct {
	url                 string
	logTimeframeMinutes int
}

func NewVictoriaLogsStreamsConnector() *VictoriaLogsStreamsConnector {
	return &VictoriaLogsStreamsConnector{}
}

func (v *VictoriaLogsStreamsConnector) FetchStreams(cfg config.Config, query string) (FetchStreamsResponseValueDTO, error) {
	fullURL := fmt.Sprintf("%s%s", cfg.VictoriaLogsURL, constants.StreamsPath)

	payload := map[string]string{
		"query": query,
		"start": fmt.Sprintf("%dm", cfg.LogTimeframeMinutes),
	}

	response, err := makePostRequest(fullURL, payload)
	if err != nil {
		log.Fatalf("Error fetching streams: %s", err)
		return FetchStreamsResponseValueDTO{}, err
	}

	var streamsResponse FetchStreamsResponseValueDTO
	if err := json.Unmarshal([]byte(response), &streamsResponse); err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}
	return streamsResponse, nil
}

func makePostRequest(fullURL string, payload map[string]string) (string, error) {
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(payload).
		Post(fullURL)

	if err != nil {
		log.Fatalf("Error making POST request: %s", err)
	}

	return resp.String(), nil
}

func (v *VictoriaLogsStreamsConnector) mapStreamResponse(reponseDTO FetchStreamsResponseValueDTO) ([]selectors.LogsStreamsDTO, error) {

	// re := regexp.MustCompile(`kubernetes\.container_name="([^"]+)",kubernetes\.pod_namespace="([^"]+)"`)

	var results []selectors.LogsStreamsDTO
	log.Println(reponseDTO)
	log.Fatal(results)
	return results, nil
	// for _, item := range responseDTO.Values {
	// 	matches := re.FindStringSubmatch(item.Value)
	// 	if len(matches) == 3 {
	// 		results = append(results, selectors.LogsStreamsDTO{
	// 			KubernetesContainerName: matches[1],
	// 			KubernetesNamespace:     matches[2],
	// 			Hits:                    item.Hits,
	// 		})
	// 	}
	// }

	// return results, nil
}
