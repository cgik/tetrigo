package game

import (
	"errors"
)

type Board struct {
	Id       int      `json:"id"`
	Blocks   []*Block `json:"blocks"`
	Cursor   []int    `json:"selector"`
	Width    int      `json:"width"`
	Height   int      `json:"height"`
	Position int      `json:"position"`
}

func createRow(rowLength int, y int) []*Block {
	var row []*Block

	for x := 0; x < rowLength; x++ {
		b := newBlock(y, x, 1)
		row = append(row, b)
	}
	return row
}

func InitBoard() *Board {
	board := new(Board)

	board.Width = 6
	board.Height = 12
	board.Position = 0
	board.Cursor = []int{0, 0}

	for y := 0; y < board.Height; y++ {
		board.Blocks = append(board.Blocks, createRow(board.Width, y)...)
	}

	return board
}

func (b *Board) GetBlock(x int, y int) *Block {
	return nil
}

func (b *Board) MoveCursor(move int) error {
	if move == 0 && b.Cursor[0] < b.Width {
		// right
		b.Cursor[0]++
		return nil
	}

	if move == 1 && b.Cursor[0] > 0 {
		// left
		b.Cursor[0]--
		return nil
	}

	if move == 3 && b.Cursor[1] < b.Height {
		// up
		b.Cursor[1]++
		return nil
	}

	if move == 4 && b.Cursor[1] > 0 {
		// down
		b.Cursor[1]--
		return nil
	}

	return errors.New("cursor moved beyond bounds")
}

func (b *Board) SwitchBlocks() error {
	// find the blocks based on the selector
	// switch the blocks
	return nil
}
