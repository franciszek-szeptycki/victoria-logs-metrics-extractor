package connectors

type FetchStreamsResponse struct {
	Values []struct {
		Value string `json:"value"`
		Hits  int    `json:"hits"`
	} `json:"values"`
}

type httpRequest struct {
	URL  string
	Body map[string]string
}

type httpResponse struct {
	Status int
	Body   string
}
