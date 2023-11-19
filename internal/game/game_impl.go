package game

import (
	"encoding/json"
)

func Init() []byte {
	b := InitBoard()
	json_response, _ := json.Marshal(&b)

	return json_response
}
