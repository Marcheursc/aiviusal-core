package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type RabbitMQConfig struct {
	URL string `mapstructure:"url"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}