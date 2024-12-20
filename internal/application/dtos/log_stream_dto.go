package dtos

type LogStreamDTO struct {
	KubernetesNamespace     string `json:"kubernetes.namespace"`
	KubernetesContainerName string `json:"kubernetes.container_name"`
	Hits                    int    `json:"hits"`
}
