package selectors

// deprecated
type LogStreamDTO struct {
	KubernetesNamespace     string `json:"kubernetes.namespace"`
	KubernetesContainerName string `json:"kubernetes.container_name"`
	Hits                    int    `json:"hits"`
}

type ResourceMetricsDTO struct {
	Resource     string
	AllHits      int
	PositiveHits int
}

type ResourceMetricsWithErrorThresholdDTO struct {
	ResourceMetricsDTO
	ErrorThreshold float32
}

type MetricsOutputDTO struct {
	Resource       string
	All            int
	Succeded       int
	Errors         int
	ErrorRate      float32
	HealthScore    float32
	IsHealthy      int
	ErrorThreshold float32
}

// type MetricsDTO struct {
// 	Container      string  `json:"container"`
// 	Namespace      string  `json:"namespace"`
// 	TotalErrors    int     `json:"totalErrors"`
// 	Total          int     `json:"total"`
// 	HealthScore    float32 `json:"healthScore"`
// 	ErrorThreshold float32 `json:"errorThreshold"`
// 	Healthy        int     `json:"healthy"`
// }
