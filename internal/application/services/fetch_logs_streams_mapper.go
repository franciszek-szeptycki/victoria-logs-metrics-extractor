package services

import (
	"fmt"
	"main/internal/application/selectors"
	"main/internal/infrastructure/connectors"
)

type FetchLogsStreamsMapper struct{}

func NewFetchLogsStreamsMapper() *FetchLogsStreamsMapper {
	return &FetchLogsStreamsMapper{}
}

func (m *FetchLogsStreamsMapper) MapResponseToDTO(input []connectors.FetchStreamsResponseValueDTO) []selectors.LogsStreamsDTO {

	fmt.Println("input", input)

	return []selectors.LogsStreamsDTO{}
}
