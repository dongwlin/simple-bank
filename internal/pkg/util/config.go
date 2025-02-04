package util

import (
	"github.com/spf13/viper"
	"time"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or env variables.
type Config struct {
	DBDriver            string        `mapstructure:"dbDriver"`
	DBSource            string        `mapstructure:"dbSource"`
	ServerAddress       string        `mapstructure:"serverAddress"`
	TokenSymmetricKey   string        `mapstructure:"tokenSymmetricKey"`
	AccessTokenDuration time.Duration `mapstructure:"accessTokenDuration"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
