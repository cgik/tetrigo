package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	log.Info("Getting config for application...")

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Unable to load config file: %w", err)
		panic(err)
	}

	return viper.GetViper()
}
