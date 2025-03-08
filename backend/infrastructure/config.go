package infrastructure

import (
	"github.com/spf13/viper"
)

type Provider struct {
	URL string
	Name string
	Type string
}

type Config struct {
	Database Database
	Providers []Provider
}

func NewConfig() (Config, error) {
	viper.SetConfigFile("env.yaml")
	viper.ReadInConfig()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}