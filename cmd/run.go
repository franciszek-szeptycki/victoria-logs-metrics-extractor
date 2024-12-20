package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the tool with data from environment variables",
	Run: func(cmd *cobra.Command, args []string) {

		},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
