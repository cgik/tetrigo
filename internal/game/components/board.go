package game

import "main/internal/game"

func initBoard() *game.Board {
	board := new(game.Board)

	board.Width = 10
	board.Height = 20

	board.Blocks = []*game.Block{}

	for height := 0; height < int(board.Height)-1; height++ {

	}

	return board
}

func createRow(rowLength int) []*game.Block {
	for x := 0; x < rowLength; x++ {

	}

}
