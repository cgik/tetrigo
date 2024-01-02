package game

import (
	"github.com/labstack/echo"
)

type Impl interface {
	CreateGame(game *Game) (*Game, error)
	GetGameByID(id int) (*Game, error)
}

type ResponseError struct {
	Message string `json:"message"`
}

type Interface struct {
	impl  Impl
	serve *echo.Echo
}

func NewHttpInterface(e *echo.Echo, impl Impl) {
	i := &Interface{
		impl: impl,
	}
	e.GET("/game/create", i.CreateGame)
}

func (s *Interface) CreateGame(c echo.Context) error {
	return c.JSON(200, "Hello World")
}
