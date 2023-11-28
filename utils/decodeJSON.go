package utils

import (
	"encoding/json"
	"io"
)

func DecodeJSON(body io.Reader, v interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(v)
}
