package dtos

type ResultLogStreamDTO struct {
	ContainerName  string  `json:"containerName"`
	Namespace      string  `json:"namespace"`
	TotalErrors    int     `json:"totalErrors"`
	Total          int     `json:"total"`
	HealthScore    float32 `json:"healthScore"`
	ErrorThreshold float32 `json:"errorThreshold"`
	Healthy        int     `json:"healthy"`
}
