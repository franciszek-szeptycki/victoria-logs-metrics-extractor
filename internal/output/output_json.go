package output

import (
	"encoding/json"
	"log"
	"main/internal/operations"
	"os"
)

func PresentJSON(results []operations.ResultLogStreamDTO) {
	jsonOutput, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}
	os.Stdout.Write(jsonOutput)
}
