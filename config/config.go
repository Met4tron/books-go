package config

import (
	"github.com/Met4tron/books-go/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

var config *Config

func LoadConfig() (*Config, error) {
	var configuration *Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cmd")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		logger.Error("Error to obtain env vars", err)
		return nil, err
	}

	err = viper.Unmarshal(&configuration)

	if err != nil {
		logger.Error("Error to parse env vars", err)
		return nil, err
	}

	config = configuration

	return config, nil
}

func GetConfig() *Config {
	return config
}
