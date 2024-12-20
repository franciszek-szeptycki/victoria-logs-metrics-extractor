package connectors

type fetchStreamsResponseValueDTO struct {
	Value string `json:"value"`
	Hits  int    `json:"hits"`
}

type FetchStreamsResponseValueDTO struct {
	Values []fetchStreamsResponseValueDTO `json:"values"`
}
