package cmd

import (
	"fmt"
	"main/internal/constants"
	"main/internal/external"
	"main/internal/operations"
	"main/internal/output"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the tool with data from environment variables",
	Run: func(cmd *cobra.Command, args []string) {

		allStreams, _ := external.FetchStreams(cfg, constants.AllStreamsHitsQuery)
		positiveStreams, _ := external.FetchStreams(cfg, constants.PositiveHitsQuery)

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
