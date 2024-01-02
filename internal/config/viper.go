package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

func LoadConfig() *viper.Viper {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	slog.Info("Getting config for application...")

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Unable to load config file: %w", err)
		panic(err)
	}

	return viper.GetViper()
}
