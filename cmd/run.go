package cmd

import (
	"main/internal/infrastructure/config"
	"main/internal/infrastructure/factories"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the tool with data from environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the environment variables
		cfg := config.Init()

		// Create the factory
		factory := factories.NewConvertLogsToMetricsFactory()

		// Create the use case
		useCase := factory.Execute()

		// Run the use case
		useCase.Execute(cfg)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
