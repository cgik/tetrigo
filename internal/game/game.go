package game

type Block struct {
	X         int32 `json:"x"`
	Y         int32 `json:"y"`
	Color     int32 `json:"color"`
	Destroyed bool  `json:"destroyed"`
	Active    bool  `json:"active"`
	Falling   bool  `json:"falling"`
}

type Board struct {
	Blocks []*Block `json:"blocks"`
	Width  int32    `json:"width"`
	Height int32    `json:"height"`
}

type Game struct {
	Board *Board `json:"board"`
}
