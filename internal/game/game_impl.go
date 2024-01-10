package game

import (
	"encoding/json"
	"log/slog"
)

type DataStore interface {
	Insert(collection string, item interface{}) error
	FindById(collection string, id string) ([]byte, error)
}

type Implementation struct {
	store DataStore
}

type Results struct {
	Game *Game `json:"game"`
}

func NewImplementation(store DataStore) *Implementation {
	return &Implementation{
		store: store,
	}
}

func (s *Implementation) CreateGame() (*Game, error) {
	game := NewGame()
	if err := s.store.Insert("games", game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *Implementation) GetGameByID(id string) (*Game, error) {
	game, err := s.store.FindById("games", id)

	if err != nil {
		return nil, err
	}

	g := ConvertToStruct(game)

	return g, nil
}

func ConvertToStruct(j []byte) *Game {
	game := new(Game)

	json.Unmarshal(j, &game)

	slog.Info("ConvertToStruct", "msg", string(j))

	return game
}
