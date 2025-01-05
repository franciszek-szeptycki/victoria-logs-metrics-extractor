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

type victoriaLogsConnector struct {
	url          string
	logTimeframe int
}

func NewVictoriaLogsConnector(url string, logTimeframe int) *victoriaLogsConnector {
	return &victoriaLogsConnector{
		url:          url,
		logTimeframe: logTimeframe,
	}
}

func (v *victoriaLogsConnector) post(httpRequest httpRequest) httpResponse {
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

func (v *victoriaLogsConnector) FetchStreams(query string) selectors.FetchStreamsResponse {
	fullURL := fmt.Sprintf("%s%s", v.url, constants.VictoriaLogsApiPathStreams)
	payload := map[string]string{
		"query": query,
		"start": fmt.Sprintf("%dm", v.logTimeframe),
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

func (v *victoriaLogsConnector) FetchLastLog(query string) selectors.LastLogReponse {
	fullURL := fmt.Sprintf("%s%s", v.url, constants.VictoriaLogsApiPathQuery)

	payload := map[string]string{
		"query": query,
		"limit": "1",
	}

	httpResponse := v.post(httpRequest{
		URL:  fullURL,
		Body: payload,
	})

	if httpResponse.Status != 200 {
		log.Fatalf("Error fetching logs: %s, status: %d", httpResponse.Body, httpResponse.Status)
	}

	var lastLogResponse selectors.LastLogReponse
	err := json.Unmarshal([]byte(httpResponse.Body), &lastLogResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling logs response: %s", err)
	}
	return lastLogResponse
}
