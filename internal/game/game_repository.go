package game

type Store interface {
	Insert(collection string, item interface{}) error
	FindById(collection string, document string, id int) error
}

type Service struct {
	store Store
}

func New(store Store) *Service {

	return &Service{
		store: store,
	}
}

func (s *Service) CreateGame(game *Game) error {

	if err := s.store.Insert("games", game); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetGameByID(id int) error {

	if err := s.store.FindById("tetris", "games", id); err != nil {
		return err
	}

	return nil
}
