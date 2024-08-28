package config

import "github.com/spf13/viper"

type Config struct {
	Broker struct {
		Host     string
		Port     int
		ClientID string
	}
	Server struct {
		Port int
	}
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
