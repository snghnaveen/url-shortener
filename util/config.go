package util

import (
	"github.com/spf13/viper"
)

var (
	EnvKey     = "ENVIRONMENT"
	EnvProd    = "production"
	EnvDev     = "development"
	EnvTesting = "testing"
)

var conf Config

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	AppPort     string `mapstructure:"APP_PORT"`
	RedisURL    string `mapstructure:"REDIS_URL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("application")
	viper.SetConfigType("env")

	// Automatically read environment variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err == nil {
		conf = config
	}
	return
}

func GetConfig() Config {
	return conf
}

func IsProdMode() bool {
	return GetConfig().Environment == EnvProd
}
