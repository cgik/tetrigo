package main

import (
	"log/slog"
	"main/internal/config"
	"main/internal/datastore"
	"main/internal/game"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {
	slog.Info("Starting application...")
	cfg := config.LoadConfig()

	httpServer := echo.New()
	httpServer.Use(middleware.CORS())

	storage := mongoSetup(cfg)
	game.NewHttpInterface(httpServer, game.NewImplementation(storage))

	cfgServerPort := cfg.GetString(`app.port`)

	if err := httpServer.Start(cfgServerPort); err != nil {
		log.Error("Unable to start application: ", err)
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
