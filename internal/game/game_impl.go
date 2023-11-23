package game

import (
	"encoding/json"
)

func Init() []byte {
	b := InitBoard()
	jsonResponse, _ := json.Marshal(&b)

	return jsonResponse
}
