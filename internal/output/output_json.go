package output

import (
	"encoding/json"
	"log"
	"main/internal/operations"
	"os"
)

func OutputJSON(results []operations.ResultLogStreamDTO) {

	jsonFormattedResults := []map[string]interface{}{}
	for _, result := range results {
		jsonFormattedResults = append(jsonFormattedResults, map[string]interface{}{
			"containerName":  result.ContainerName,
			"namespace":      result.Namespace,
			"totalErrors":    result.TotalErrors,
			"total":          result.Total,
			"healthScore":    result.HealthScore,
			"errorThreshold": result.ErrorThreshold,
		})
	}
	jsonOutput, err := json.MarshalIndent(jsonFormattedResults, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}
	os.Stdout.Write(jsonOutput)
}
