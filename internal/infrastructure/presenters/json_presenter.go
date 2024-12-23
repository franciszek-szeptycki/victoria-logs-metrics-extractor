package presenters

import (
	"encoding/json"
	"log"
	"main/internal/application/selectors"
	"os"
)

type JSONPresenter struct{}

func NewJSONPresenter() *JSONPresenter {
	return &JSONPresenter{}
}

func (j *JSONPresenter) Present(results []selectors.ResultLogStreamDTO) {
	jsonOutput, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}
	os.Stdout.Write(jsonOutput)
}
