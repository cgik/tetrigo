package game

import "main/internal/auth"

type Game struct {
	Id         int          `json:"id"`
	Board      *Board       `json:"board"`
	Score      int          `json:"score"`
	User       *auth.User   `json:"user"`
	Spectators []*auth.User `json:"spectators"`
}

func NewGame() *Game {
	game := new(Game)

	board := InitBoard()

	game.Board = board
	game.Score = 0

	return game
}
