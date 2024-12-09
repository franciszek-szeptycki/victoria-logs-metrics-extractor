package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/internal/config"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the tool with data from environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		fmt.Printf("VictoriaLogsURL: %s\n", cfg.VictoriaLogsURL)
		fmt.Printf("LogTimeframeMinutes: %d\n", cfg.LogTimeframeMinutes)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
