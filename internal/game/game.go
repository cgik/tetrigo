package game

import "main/internal/auth"

type Game struct {
	Id         int32        `json:"id"`
	Board      *Board       `json:"board"`
	Score      int32        `json:"score"`
	User       *auth.User   `json:"user"`
	Spectators []*auth.User `json:"spectators"`
}

func NewGame() *Game {
	game := new(Game)

	game.Board = InitBoard()

	return game
}
