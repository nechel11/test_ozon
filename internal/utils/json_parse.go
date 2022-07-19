package utils

import (
	"encoding/json"
	"io"
)



// Decode r.body to json
func Decoder_json(s interface{}, request io.ReadCloser) error {
	decode := json.NewDecoder(request)
	decode.DisallowUnknownFields()
	return decode.Decode(s)
}