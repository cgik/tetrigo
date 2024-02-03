package game

type DataStore interface {
	Insert(collection string, item interface{}) error
	FindById(collection string, id string) ([]byte, error)
	FindAll(collection string, key string) ([]byte, error)
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

func (s *Implementation) CreateGame() ([]byte, error) {
	game := NewGame()
	if err := s.store.Insert("games", game); err != nil {
		return nil, err
	}

	g, err := GameToJson(game)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (s *Implementation) GetGameByID(id string) ([]byte, error) {
	game, err := s.store.FindById("games", id)

	if err != nil {
		return nil, err
	}

	g, err := ValidateGame(game)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (s *Implementation) ListGames() ([]byte, error) {
	games, err := s.store.FindAll("games", "_id")

	if err != nil {
		return nil, err
	}

	return games, nil
}

func (s *Implementation) MoveCursor(game *Game, move int) error {
	return nil
}

func (s *Implementation) SwitchBlocks(game *Game) error {
	return nil
}
