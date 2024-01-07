package game

import (
	"github.com/labstack/echo/v4"
)

type Impl interface {
	CreateGame() (*Game, error)
	GetGameByID(id int) (*Game, error)
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
	g, err := s.impl.CreateGame()

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, g)
}
