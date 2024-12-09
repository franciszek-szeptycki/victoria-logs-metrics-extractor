package operations

import (
	"main/internal/external"
	"reflect"
	"testing"
)

func TestAnalyzeLogStreams(t *testing.T) {

	allLogStreams := []external.LogStreamDTO{
		{KubernetesNamespace: "abc", KubernetesContainerName: "def", Hits: 10},
	}
	positiveLogStreams := []external.LogStreamDTO{
		{KubernetesNamespace: "abc", KubernetesContainerName: "def", Hits: 8},
	}

	results, _ := AnalyzeLogStreams(allLogStreams, positiveLogStreams, 0.1)

	expectedResults := []ResultLogStreamDTO{
		{ContainerName: "def", Namespace: "abc", TotalErrors: 2, Total: 10, HealthScore: 0.8, ErrorThreshold: 0.1, Healthy: 1},
	}

	if !reflect.DeepEqual(results, expectedResults) {
		t.Errorf("Results do not match expected values.\nExpected: %+v\nGot: %+v", expectedResults, results)
	}
}
