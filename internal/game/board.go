package game

type Board struct {
	Id       int32    `json:"id"`
	Blocks   []*Block `json:"blocks"`
	Width    int32    `json:"width"`
	Height   int32    `json:"height"`
	Position int32    `json:"position"`
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

	board.Width = 10
	board.Height = 20
	board.Position = 0

	for y := 0; y < int(board.Height); y++ {
		board.Blocks = append(board.Blocks, createRow(int(board.Width), y)...)
	}

	return board
}
