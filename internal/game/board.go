package game

type Board struct {
	Id       int32    `json:"id"`
	Blocks   []*Block `json:"blocks"`
	Selector [2]int32 `json:"selector"`
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
	board.Selector = [2]int32{0, 0}

	for y := 0; y < int(board.Height); y++ {
		board.Blocks = append(board.Blocks, createRow(int(board.Width), y)...)
	}

	return board
}

func (b *Board) GetBlock(x int, y int) *Block {
	return nil
}

func (b *Board) MoveCursor(move int) error {
	if b.Selector[0] <= 0 || b.Selector[0] >= b.Width {
		return nil
	}

	if b.Selector[1] <= b.Height || b.Selector[1] >= b.Height {
		return nil
	}

	switch move {
	case 0: // right
		b.Selector[0]++
	case 1: // left
		b.Selector[0]--
	case 2: // down
		b.Selector[1]++
	case 3: // up
		b.Selector[1]--
	}

	return nil
}

func (b *Board) SwitchBlocks() error {
	// find the blocks based on the selector
	// switch the blocks
	return nil
}
