package pkg

import (
	"encoding/json"
)

func Marshal(data interface{}) ([]byte, error) {
	// Marshal the struct to JSON
	byteArr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return byteArr, nil
}
