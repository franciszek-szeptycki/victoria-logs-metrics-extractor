package services

import (
	"fmt"
	"log"
	"main/internal/application/selectors"
	"regexp"
)

type FetchLogsStreamsMapper struct{}

func NewFetchLogsStreamsMapper() *FetchLogsStreamsMapper {
	return &FetchLogsStreamsMapper{}
}

func (m *FetchLogsStreamsMapper) MapStreamsResponseToLogs(input selectors.FetchStreamsResponse) []selectors.LogStreamDTO {

	fmt.Println("input", input)

	outputArray := []selectors.LogStreamDTO{}
	for _, stream := range input.Values {

		regexContainerName := regexp.MustCompile(`kubernetes\.container_name="([^"]+)"`)
		containerNameMatch := regexContainerName.FindStringSubmatch(stream.Value)

		regexNamespace := regexp.MustCompile(`kubernetes\.pod_namespace="([^"]+)"`)
		namespaceMatch := regexNamespace.FindStringSubmatch(stream.Value)

		if len(containerNameMatch) == 0 || len(namespaceMatch) == 0 {
			log.Printf("Error with: containerNameMatch: %v, namespaceMatch: %v", containerNameMatch, namespaceMatch)
			continue
		}

		outputArray = append(outputArray, selectors.LogStreamDTO{
			KubernetesNamespace:     namespaceMatch[1],
			KubernetesContainerName: containerNameMatch[1],
			Hits:                    stream.Hits,
		})
	}

	return outputArray
}
