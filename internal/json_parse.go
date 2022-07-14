package internal

import (
	"encoding/json"
	"io"
)

// Users struct for parse requests
type JsonUrl struct {
	Url string `json:"url"`
}

// Decode r.body to json
func decoder_json(s interface{}, request io.ReadCloser) error {
	decode := json.NewDecoder(request)
	decode.DisallowUnknownFields()
	return decode.Decode(s)
}