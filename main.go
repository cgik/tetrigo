package main

import (
	"github.com/spf13/viper"
	"log/slog"
	"main/internal/datastore"
	"net/http"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	slog.Info("Getting config for application...")

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Unable to load config file: %w", err)
		panic(err)
	}
}

func main() {
	slog.Info("Starting application...")

	mongo := datastore.ConnectMongo(
		viper.GetString(`mongo.uri`),
		viper.GetString(`mongo.database`))

	collections := viper.GetStringSlice(`mongo.collections`)
	if err := mongo.SetupDatabase(collections); err != nil {
		slog.Error("Unable to setup database: ", err)
		panic(err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Up!"))
	})

	http.HandleFunc("/api/v1/game/create", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s := &http.Server{
		Addr: viper.GetString(`app.port`),
	}
	err := s.ListenAndServe()

	slog.Info("Application started, listening on port: ", s.Addr)

	if err != nil {
		slog.Error("Unable to start application: ", err)
		panic(err)
	}
}
