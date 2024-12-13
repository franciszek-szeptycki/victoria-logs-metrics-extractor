package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/internal/config"
	"main/internal/constants"
	"main/internal/external"
	"main/internal/operations"
	"main/internal/output"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the tool with data from environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()

		allStreams, _ := external.FetchStreamsHits(cfg, constants.AllStreamsHitsQuery)
		positiveStreams, _ := external.FetchStreamsHits(cfg, constants.PositiveHitsQuery)

		results, err := operations.AnalyzeLogStreams(allStreams, positiveStreams, cfg.ErrorThreshold)
		if err != nil {
			fmt.Println(err)
			return
		}

		output.PresentJSON(results)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
