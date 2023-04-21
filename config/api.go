package config

import "fmt"

type APIConfig struct {
	Port               string `mapstructure:"PORT" default:"1323"`
	LogLevel           string `mapstructure:"LOG_LEVEL" default:"info"`
	AuthToken          string `mapstructure:"AUTH_TOKEN" validate:"required"`
	RedisURL           string `mapstructure:"REDIS_URL" validate:"required"`
	MySQLDSN           string `mapstructure:"MYSQL_DSN" validate:"required"`
	JWTSecret          string `mapstructure:"JWT_SECRET" validate:"required"`
	InitData           string `mapstructure:"INIT_DATA" validate:"required"`
	DefaultDevPassword string `mapstructure:"DEFAULT_DEV_PASSWORD" validate:"required"`
}

func NewAPIConfig() *APIConfig {
	config := new(APIConfig)
	AutoLoadConfig(config)
	return config
}

func (cfg APIConfig) ListenAddr() string {
	return fmt.Sprintf(":%s", cfg.Port)
}
