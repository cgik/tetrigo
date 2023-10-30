package game

import "main/internal/game"

func newBlock(x int32, y int32, color int32) game.Block {
	return game.Block{
		X:         x,
		Y:         y,
		Color:     color,
		Destroyed: false,
		Falling:   false,
	}
}
