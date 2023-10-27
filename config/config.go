package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DBHost     string
	DBPort     int
	DBUsername string
	DBPassword string
	DBDatabase string
	DBSSLMode  string
}

func LoadConfig(configFile string) (*AppConfig, error) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cool_service")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *AppConfig) DBConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBDatabase, c.DBSSLMode)
}
