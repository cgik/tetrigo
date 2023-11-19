package game

type Store interface {
	Insert(collection string, item interface{}) error
	FindById(collection string, document string, id int) error
}

type Service struct {
	store Store
}

func NewGameService(store Store) *Service {
	return &Service{
		store: store,
	}
}

func (g *Service) CreateGame(game *Game) error {
	if err := g.store.Insert("game", game); err != nil {
		return err
	}
	return nil
}

func (g *Service) GetGameByID(id int) error {
	if err := g.store.FindById("game", "game", id); err != nil {
		return err
	}
	return nil
}
