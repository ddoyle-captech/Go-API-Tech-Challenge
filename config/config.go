package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
ENV: dev
LOG_LEVEL: DEBUG

DATABASE_CONTAINER_NAME: courses-db-container
DATABASE_NAME: coursesDB
DATABASE_USER: courses-db-user
DATABASE_PASSWORD: courses-db-password
DATABASE_HOST: localhost
DATABASE_PORT: 5432
DATABASE_RETRY_DURATION_SECONDS: 3

HTTP_DOMAIN: localhost
HTTP_PORT: :8080
*/

type Config struct {
	Env      string `mapstructure:"ENV"`
	LogLevel string `mapstructure:"LOG_LEVEL"`

	DBContainerName string `mapstructure:"DATABASE_CONTAINER_NAME"`
	DBName          string `mapstructure:"DATABASE_NAME"`
	DBUser          string `mapstructure:"DATABASE_USER"`
	DBPassword      string `mapstructure:"DATABASE_PASSWORD"`
	DBHost          string `mapstructure:"DATABASE_HOST"`
	DBPort          string `mapstructure:"DATABASE_PORT"`
	DBRetryDuration string `mapstructure:"DATABASE_RETRY_DURATION_SECONDS"`

	HTTPDomain string `mapstructure:"HTTP_DOMAIN"`
	HTTPPort   string `mapstructure:"HTTP_PORT"`
}

func Load(path string) (*Config, error) {
	// Tell Viper where to find the config file and what type it is
	viper.SetConfigFile(path)
	viper.SetConfigType("env")

	// Parse .env file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading in .env config file: %s", err.Error())
	}
	viper.AutomaticEnv()

	// Read values into Config struct
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("error parsing .env file into Config struct: %s", err.Error())
	}
	return cfg, nil
}
