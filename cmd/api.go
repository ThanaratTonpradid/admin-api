package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"mini-api/internal/api"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api gateway",
	Run:   apiRun,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func apiRun(cmd *cobra.Command, _ []string) {
	log.Info(cmd.Short)
	fx.New(api.Module, fx.WithLogger(FxCmdLogger)).Run()
}
