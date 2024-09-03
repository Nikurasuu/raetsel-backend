package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Port int
	}
	Mongo struct {
		Host     string
		Port     int
		DataBase string
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
