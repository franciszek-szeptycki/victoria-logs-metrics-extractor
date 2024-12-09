package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/internal/config"
	"main/internal/connectors"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the tool with data from environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		fmt.Printf("VictoriaLogsURL: %s\n", cfg.VictoriaLogsURL)
		fmt.Printf("LogTimeframeMinutes: %d\n", cfg.LogTimeframeMinutes)

		victoriaLogsStreamsConnector := connectors.NewVictoriaLogsStreamsConnector(cfg.VictoriaLogsURL)
		victoriaLogsStreamsConnector.FetchAllStreamsHits()
		victoriaLogsStreamsConnector.FetchPositiveStreamsHits()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
