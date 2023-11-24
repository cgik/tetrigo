package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log/slog"
	"main/internal/datastore"
	game "main/internal/game"
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
	app := fiber.New()

	mongo := datastore.ConnectMongo(
		viper.GetString(`mongo.uri`),
		viper.GetString(`mongo.database`))

	game.New(mongo)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("I'm awake, I'm awake...")
	})

	app.Get("/api/v1/game/init", func(c *fiber.Ctx) error {
		initGame := game.Init()

		return c.SendString(string(initGame))
	})

	slog.Info("Application started.")
	err := app.Listen(viper.GetString(`app.port`))

	if err != nil {
		slog.Error("Unable to start application: ", err)
		panic(err)
	}
}
