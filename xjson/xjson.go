package xjson

import (
	"encoding/json"
	"os"
)

// ParseFile parses json data from file at path and unmarshal into o.
func ParseFile(path string, o interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(o)
}
