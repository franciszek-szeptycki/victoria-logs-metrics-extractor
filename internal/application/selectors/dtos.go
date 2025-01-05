package selectors

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
