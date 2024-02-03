package game

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
)

func ValidateGame(j []byte) ([]byte, error) {
	game := new(Game)

	if err := json.Unmarshal(j, &game); err != nil {
		log.Error(err)
		return nil, err
	}

	rawGame, err := GameToJson(game)

	if err != nil {
		return nil, err
	}

	return rawGame, nil
}

func GameToJson(g *Game) ([]byte, error) {
	j, err := json.Marshal(g)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return j, nil
}
