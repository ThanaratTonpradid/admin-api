package cmd

import (
	"os"

	"github.com/dollarsignteam/go-logger"
	"github.com/spf13/cobra"
	"go.uber.org/fx/fxevent"

	"mini-api/config"
	"mini-api/helper"
)

var cfgFile string
var log *logger.Logger

var rootCmd = &cobra.Command{
	Use:     "mini-api",
	Short:   "mini api",
	Version: config.Version,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	log = logger.NewLogger(logger.LoggerOptions{
		Name:       helper.GetPackageName(),
		HideCaller: true,
	})

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is local.env)")
}

func initConfig() {
	config.AutoReadConfig(cfgFile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func FxCmdLogger() fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log.Desugar()}
}
