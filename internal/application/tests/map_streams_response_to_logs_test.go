package services

import (
	"encoding/json"
	"main/internal/application/selectors"
	"main/internal/application/services"
	"testing"
)

func TestFetchLogsStreamsMapper(t *testing.T) {
	jsonText := `{"values":[{"value":"{}","hits":31},{"value":"{kubernetes.container_name=\"coredns\",kubernetes.pod_namespace=\"kube-system\"}","hits":4},{"value":"{kubernetes.container_name=\"mariadb\",kubernetes.pod_namespace=\"paris\"}","hits":5}]}`

	var input selectors.FetchStreamsResponse
	err := json.Unmarshal([]byte(jsonText), &input)
	if err != nil {
		t.Errorf("Error marshalling JSON: %s", err)
	}

	expectedOutput := []selectors.LogStreamDTO{
		{
			KubernetesNamespace:     "kube-system",
			KubernetesContainerName: "coredns",
			Hits:                    4,
		}, {
			KubernetesNamespace:     "paris",
			KubernetesContainerName: "mariadb",
			Hits:                    5,
		},
	}

	mapper := services.NewFetchLogsStreamsMapper()
	output := mapper.MapStreamsResponseToLogs(input)

	if !compareLogStreamDTOs(expectedOutput, output) {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func compareLogStreamDTOs(a, b []selectors.LogStreamDTO) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].KubernetesNamespace != b[i].KubernetesNamespace {
			return false
		}

		if a[i].KubernetesContainerName != b[i].KubernetesContainerName {
			return false
		}

		if a[i].Hits != b[i].Hits {
			return false
		}
	}

	return true
}
