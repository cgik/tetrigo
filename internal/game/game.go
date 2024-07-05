package game

import (
	"main/internal/auth"
	"main/internal/datastore"
)

type Game struct {
	Id         string       `json:"id"`
	Board      *Board       `json:"board"`
	Score      int          `json:"score"`
	User       *auth.User   `json:"user"`
	Spectators []*auth.User `json:"spectators"`
}

func NewGame() *Game {
	game := new(Game)

	board := InitBoard()

	game.Id, _ = datastore.GenerateId()
	game.Board = board
	game.Score = 0

	return game
}
