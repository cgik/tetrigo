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

	board.Blocks = []*Block{}

	for height := 0; height < int(board.Height)-1; height++ {
		createRow(int(board.Width), height)
	}

	return board
}
