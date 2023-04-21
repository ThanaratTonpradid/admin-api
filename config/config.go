package config

import (
	"github.com/dollarsignteam/go-logger"
	"github.com/spf13/viper"

	"mini-api/helper"
)

var log *logger.Logger

func init() {
	log = logger.NewLogger(logger.LoggerOptions{
		Name:       helper.GetPackageName(),
		HideCaller: true,
	})
}

func AutoReadConfig(cfgFile string) {
	configFile := "local.env"
	if cfgFile != "" {
		configFile = cfgFile
	}

	if err := helper.ReadConfig(configFile); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else if cfgFile != "" {
		log.Fatalf("Config file: %s, %s", viper.ConfigFileUsed(), err)
	}
}

func AutoLoadConfig(config interface{}) {
	if err := helper.LoadConfig(config); err != nil {
		log.Fatal(err.Error())
	}
	log.Infof("Config loaded: %+v", config)
}
