package main

import (
	"log"

	game "main/internal/game"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("I'm awake, I'm awake...")
	})

	app.Get("/api/v1/game/init", func(c *fiber.Ctx) error {
		inital_game := game.Init()

		return c.SendString(string(inital_game))
	})

	log.Fatal(app.Listen(":18080"))
}
