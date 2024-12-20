package services

import (
	"main/internal/application/selectors"
	"testing"
)

func TestFetchLogsStreamsMapper(t *testing.T) {
	// {
	//   "values": [
	//     {
	//       "value": "{}",
	//       "hits": 31
	//     },
	//     {
	//       "value": "{kubernetes.container_name=\"coredns\",kubernetes.pod_namespace=\"kube-system\"}",
	//       "hits": 4
	//     },
	//     {
	//       "value": "{kubernetes.container_name=\"mariadb\",kubernetes.pod_namespace=\"paris\"}",
	//       "hits": 4
	//     }
	//   ]
	// }
	input := FetchLogsStreamsResponse{
		Values: []FetchLogsStreamsResponseValue{
			{
				Value: "{}",
				Hits:  31,
			},
			{
				Value: "{kubernetes.container_name=\"coredns\",kubernetes.pod_namespace=\"kube-system\"}",
				Hits:  4,
			},
			{
				Value: "{kubernetes.container_name=\"mariadb\",kubernetes.pod_namespace=\"paris\"}",
				Hits:  5,
			},
		},
	}

	expectedOutput := []selectors.LogsStreamsDTO{
		{
			KubernetesNamespace:     "",
			KubernetesContainerName: "",
			Hits:                    31,
		}, {
			KubernetesNamespace:     "kube-system",
			KubernetesContainerName: "coredns",
			Hits:                    4,
		}, {
			KubernetesNamespace:     "paris",
			KubernetesContainerName: "mariadb",
			Hits:                    5,
		},
	}

	mapper := NewFetchLogsStreamsMapper()
	output := mapper.MapResponseToDTO(input)

	if !compareLogsStreamsDTOs(output, expectedOutput) {
		t.Errorf("Expected output: %v, got: %v", expectedOutput, output)
	}
}

func compareLogsStreamsDTOs(a, b []selectors.LogsStreamsDTO) bool {
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
