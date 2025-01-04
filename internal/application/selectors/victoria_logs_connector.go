package selectors

type FetchStreamsResponse struct {
	Values []struct {
		Value string `json:"value"`
		Hits  int    `json:"hits"`
	} `json:"values"`
}

type LastLogReponse struct {
	CustomErrorThreshold *float32 `json:"kubernetes.pod_labels.custom_error_threshold"`
}
