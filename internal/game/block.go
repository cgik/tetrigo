package game

type Block struct {
	X         int   `json:"x"`
	Y         int   `json:"y"`
	Color     int32 `json:"color"`
	Destroyed bool  `json:"destroyed"`
	Active    bool  `json:"active"`
	Falling   bool  `json:"falling"`
}

func moveUp(b *Block) {
	b.Y -= 1
}

func moveDown(b *Block) {
	b.Y += 1
}

func moveLeft(b *Block) {
	b.X -= 1
}

func moveRight(b *Block) {
	b.X += 1
}

func newBlock(x int, y int, color int32) *Block {
	return &Block{
		X:         x,
		Y:         y,
		Color:     color,
		Destroyed: false,
		Falling:   false,
	}
}
