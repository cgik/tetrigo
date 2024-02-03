package game

import (
	"reflect"
	"testing"
)

func TestInitBoard(t *testing.T) {
	board := InitBoard()

	if board == nil {
		t.Errorf("Board is nil")
	}
}

func TestMoveCursor(t *testing.T) {
	board := InitBoard()

	for i := 0; i < board.Width+2; i++ {
		previous := board.Cursor
		err := board.MoveCursor(0)

		if err != nil {
			if !reflect.DeepEqual(previous, board.Cursor) {
				t.Error(err)
			}
			break
		}
	}

	for i := 0; i < board.Width+2; i++ {
		previous := board.Cursor
		err := board.MoveCursor(1)

		if err != nil {
			if !reflect.DeepEqual(previous, board.Cursor) {
				t.Error(err)
			}
			break
		}
	}

}

func TestSwitchBlocks(t *testing.T) {
	board := InitBoard()

	//board.SwitchBlocks()

	t.Log(board.Blocks)
}
