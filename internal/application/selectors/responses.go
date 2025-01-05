package selectors

type FetchStreamsResponse struct {
	Values []struct {
		Value string `json:"value"`
		Hits  int    `json:"hits"`
	} `json:"values"`
}

type LastLogReponse struct {
	// CustomErrorThreshold string `json:"kubernetes.pod_labels.custom_error_threshold"`
	CustomErrorThreshold string `json:"kubernetes.pod_labels.apps.kubernetes.io/pod-index"`
}
