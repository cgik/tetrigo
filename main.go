package main

import (
	"github.com/spf13/viper"
	"log"

	"main/internal/datastore"
	game "main/internal/game"

	"github.com/gofiber/fiber/v2"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	log.Println("Getting config for application...")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to load config file: %w", err)
	}
}

func main() {
	app := fiber.New()
	mongo := datastore.ConnectMongo(viper.GetString(`mongo.uri`))

	log.Print("Connected to mongo: ", mongo)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("I'm awake, I'm awake...")
	})

	app.Get("/api/v1/game/init", func(c *fiber.Ctx) error {
		initGame := game.Init()

		return c.SendString(string(initGame))
	})

	log.Fatal(app.Listen(viper.GetString(`app.port`)))
}
