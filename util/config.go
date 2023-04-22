package util

import (
	"os"

	"github.com/spf13/viper"
)

var (
	EnvKey     = "ENVIRONMENT"
	EnvProd    = "production"
	EnvDev     = "development"
	EnvTesting = "testing"

	ReqCounterKey = "req_counter"
)

var conf Config

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	AppPort     string `mapstructure:"APP_PORT"`
	RedisURL    string `mapstructure:"REDIS_URL"`

	MaxMetricsRecords int `mapstructure:"MAX_METRICS"`
	TTLShortenKey     int `mapstructure:"TTL_SHORTEN_URL_IN_SECONDS"`
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

	setMissingIfAny()

	return
}

func setMissingIfAny() {
	if GetConfig().MaxMetricsRecords == 0 {
		// Default max metrics
		conf.MaxMetricsRecords = 3
	}

	if GetConfig().TTLShortenKey == 0 {
		// Default TTLShortenKey
		conf.TTLShortenKey = 3600
	}
}

func GetConfig() Config {
	return conf
}

func IsProdMode() bool {
	return GetConfig().Environment == EnvProd ||
		os.Getenv(EnvKey) == EnvProd
}

func IsTestingMode() bool {
	return GetConfig().Environment == EnvTesting ||
		os.Getenv(EnvKey) == EnvTesting
}
