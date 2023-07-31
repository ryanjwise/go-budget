package io

import (
	"encoding/json"
	"os"
)

func SaveJSONToFile(data interface{}, filename string) error {
	jsonData, e := json.MarshalIndent(data, "", "  ")
	if e != nil {
		return e
	}

	return os.WriteFile(filename, jsonData, 0644)
}
