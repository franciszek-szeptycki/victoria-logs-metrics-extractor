package selectors

type FetchStreamsResponse struct {
	Values []struct {
		Value string `json:"value"`
		Hits  int    `json:"hits"`
	} `json:"values"`
}
