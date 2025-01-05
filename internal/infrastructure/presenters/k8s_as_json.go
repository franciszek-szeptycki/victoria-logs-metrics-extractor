package presenters

import (
	"encoding/json"
	"main/internal/application/selectors"
	"os"
	"regexp"
)

type K8sJsonPresenter struct{}

type K8sJsonOutput struct {
	ContainerName  string  `json:"containerName"`
	Namespace      string  `json:"namespace"`
	Total          int     `json:"total"`
	TotalSuccess   int     `json:"totalSuccess"`
	TotalErrors    int     `json:"totalErrors"`
	ErrorThreshold float32 `json:"errorThreshold"`
	ErrorRate      float32 `json:"errorRate"`
	HealthScore    float32 `json:"healthScore"`
	IsHealthy      int     `json:"isHealthy"`
}

func (k *K8sJsonPresenter) Present(output []selectors.MetricsOutputDTO) {

	var k8sJsonOutput []K8sJsonOutput
	for _, o := range output {
		containerName := k.getContainerNameFromResource(o.Resource)
		namespace := k.getNamespaceFromResource(o.Resource)

		k8sJsonOutput = append(k8sJsonOutput, K8sJsonOutput{
			ContainerName:  containerName,
			Namespace:      namespace,
			Total:          o.All,
			TotalSuccess:   o.Succeded,
			TotalErrors:    o.Errors,
			ErrorThreshold: o.ErrorThreshold,
			ErrorRate:      o.ErrorRate,
			HealthScore:    o.HealthScore,
			IsHealthy:      o.IsHealthy,
		})
	}

	jsonOutput, err := json.MarshalIndent(k8sJsonOutput, "", "  ")
	if err != nil {
		os.Stderr.Write([]byte("Error marshalling JSON: " + err.Error()))
	}

	os.Stdout.Write(jsonOutput)

}

func (k *K8sJsonPresenter) getContainerNameFromResource(resourceValue string) string {
	regexContainerName := regexp.MustCompile(`kubernetes\.container_name="([^"]+)"`)
	containerNameMatch := regexContainerName.FindStringSubmatch(resourceValue)

	if len(containerNameMatch) == 0 {
		return ""
	}

	return containerNameMatch[1]
}

func (k *K8sJsonPresenter) getNamespaceFromResource(resourceValue string) string {
	regexNamespace := regexp.MustCompile(`kubernetes\.pod_namespace="([^"]+)"`)
	namespaceMatch := regexNamespace.FindStringSubmatch(resourceValue)

	if len(namespaceMatch) == 0 {
		return ""
	}

	return namespaceMatch[1]
}
