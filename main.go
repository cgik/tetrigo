package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log/slog"
	"main/internal/config"
	"main/internal/datastore"
	"main/internal/game"
)

func main() {
	slog.Info("Starting application...")
	cfg := config.LoadConfig()

	httpServer := echo.New()

	storage := mongoSetup(cfg)
	game.NewHttpInterface(httpServer, game.NewImplementation(storage))

	cfgServerPort := cfg.GetString(`app.port`)

	if err := httpServer.Start(cfgServerPort); err != nil {
		slog.Error("Unable to start application: ", err)
	}

}

func mongoSetup(cfg *viper.Viper) *datastore.DataStore {
	if cfg.GetString(`feature.mongo`) == `true` {
		cfgMongoUri := cfg.GetString(`mongo.uri`)
		cfgMongoDatabase := cfg.GetString(`mongo.database`)

		mongo := datastore.ConnectMongo(
			cfgMongoUri,
			cfgMongoDatabase)

		cfgCollections := cfg.GetStringSlice(`mongo.collections`)
		mongo.SetupDatabase(cfgCollections)

		return mongo
	}

	return nil
}
