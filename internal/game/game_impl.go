package game

import "go.mongodb.org/mongo-driver/bson"

type DataStore interface {
	Insert(collection string, item interface{}) error
	FindById(collection string, document string, id int) (bson.M, error)
}

type Implementation struct {
	store DataStore
}

func NewImplementation(store DataStore) *Implementation {
	return &Implementation{
		store: store,
	}
}

func (s *Implementation) CreateGame(game *Game) (*Game, error) {

	if err := s.store.Insert("games", game); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Implementation) GetGameByID(id int) (*Game, error) {

	if _, err := s.store.FindById("tetris", "games", id); err != nil {
		return nil, err
	}

	return nil, nil
}
