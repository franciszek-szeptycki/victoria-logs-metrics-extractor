package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/internal/config"
	"main/internal/constants"
	"main/internal/external"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the tool with data from environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		fmt.Printf("VictoriaLogsURL: %s\n", cfg.VictoriaLogsURL)
		fmt.Printf("LogTimeframeMinutes: %d\n", cfg.LogTimeframeMinutes)

		external.FetchStreamsHits(cfg, constants.AllStreamsHitsQuery)
		external.FetchStreamsHits(cfg, constants.PositiveHitsQuery)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
