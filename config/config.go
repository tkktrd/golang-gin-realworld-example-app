package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	MySQLHost string
	MySQLPort string
	MySQLUser string
	MySQLPassword string
	MySQLDatabaseName string
}

func LoadConfig(filename string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath("/config")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found, %v", err)
		}
	}

	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return &config, nil
}