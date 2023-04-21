package helper

import (
	"fmt"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func ReadConfig(cfgFile string) error {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func LoadConfig(config interface{}) error {
	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("unmarshal config: %s", err)
	}

	if err := defaults.Set(config); err != nil {
		return fmt.Errorf("defaults config: %s", err)
	}

	return validator.New().Struct(config)
}
