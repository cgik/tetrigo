package wasm

import (
	"main/internal/game"
)

func CreateGame() (*game.Game, error) {
	return game.NewGame(), nil
}

func Test() {
	println("Hello from wasm")
}
