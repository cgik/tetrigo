package game

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Impl interface {
	CreateGame() ([]byte, error)
	GetGameByID(id string) ([]byte, error)
	ListGames() ([]byte, error)
}

type Interface struct {
	impl  Impl
	serve *echo.Echo
}

type Data struct {
	Game    json.RawMessage `json:"game"`
	Message string          `json:"message"`
}

type Response struct {
	Success    string `json:"success"`
	DevMessage string `json:"dev_message"`
	Data       Data   `json:"data"`
}

func NewHttpInterface(e *echo.Echo, impl Impl) {
	i := &Interface{
		impl:  impl,
		serve: e,
	}
	i.serve.GET("/game/create", i.CreateGame)
	i.serve.GET("/game/load/:id", i.LoadGame)
	i.serve.GET("/game/list", i.ListGames)
}

func (s *Interface) CreateGame(c echo.Context) error {
	res := Response{}

	g, err := s.impl.CreateGame()

	if err != nil {
		res.Success = "false"
		res.Data.Message = fmt.Sprint(err)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Success = "true"
	res.Data.Game = g

	return c.JSON(http.StatusOK, res)
}

func (s *Interface) LoadGame(c echo.Context) error {
	id := c.Param("id")
	res := Response{}

	g, err := s.impl.GetGameByID(id)

	if err != nil {
		res.Success = "false"
		res.Data.Message = fmt.Sprint(err)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Success = "true"
	res.Data.Game = g

	return c.JSON(http.StatusOK, res)
}

func (s *Interface) ListGames(c echo.Context) error {
	res := Response{}

	games, err := s.impl.ListGames()

	if err != nil {
		res.Success = "false"
		res.Data.Message = fmt.Sprint(err)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Success = "true"
	res.Data.Game = games

	return c.JSON(http.StatusOK, res)
}
