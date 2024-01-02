package main

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log/slog"
	"main/internal/config"
	"main/internal/datastore"
	"main/internal/game"
)

func main() {
	slog.Info("Starting application...")
	cfg := config.LoadConfig()

	m, _ := mongoSetup(cfg)

	cfgServerPort := cfg.GetString(`app.port`)

	e := echo.New()
	game.NewHttpInterface(e, game.NewImplementation(m))

	if err := e.Start(cfgServerPort); err != nil {
		slog.Error("Unable to start application: ", err)
	}

}

func mongoSetup(cfg *viper.Viper) (*datastore.DataStore, error) {
	if cfg.GetString(`feature.mongo`) == `true` {
		cfgMongoUri := cfg.GetString(`mongo.uri`)
		cfgMongoDatabase := cfg.GetString(`mongo.database`)

		mongo := datastore.ConnectMongo(
			cfgMongoUri,
			cfgMongoDatabase)

		cfgCollections := cfg.GetStringSlice(`mongo.collections`)
		mongo.SetupDatabase(cfgCollections)

		return mongo, nil
	}

	return nil, nil
}
