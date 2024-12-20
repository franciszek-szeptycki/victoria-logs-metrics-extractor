package services

import (
	"encoding/json"
	"fmt"
	"main/internal/application/dtos"
	"main/internal/infrastructure/connectors"
	"regexp"
)

type RetrieveStreamsService struct {
	victoriaLogsConnector *connectors.VictoriaLogsStreamsConnector
}

func NewRetrieveStreamsService(victoriaLogsConnector *connectors.VictoriaLogsStreamsConnector) *RetrieveStreamsService {
	return &RetrieveStreamsService{
		victoriaLogsConnector: victoriaLogsConnector,
	}
}

// func (a *RetrieveStreamsService) Execute(query string, logTimeframe int) (map[string]int, error) {

// }

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
