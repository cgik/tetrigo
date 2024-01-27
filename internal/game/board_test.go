package game

import "testing"

func TestInitBoard(t *testing.T) {
	board := InitBoard()

	if board == nil {
		t.Errorf("Board is nil")
	}
}

func TestMoveCursor(t *testing.T) {
	board := InitBoard()

	err := board.MoveCursor(0)

	if err != nil {
		t.Errorf("Error moving cursor right")
	}

	err = board.MoveCursor(1)

	if err != nil {
		t.Errorf("Error moving cursor left")
	}

	err = board.MoveCursor(2)

	if err != nil {
		t.Errorf("Error moving cursor down")
	}

	err = board.MoveCursor(3)

	if err != nil {
		t.Errorf("Error moving cursor up")
	}
}
