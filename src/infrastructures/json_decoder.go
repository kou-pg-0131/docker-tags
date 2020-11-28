package infrastructures

import (
	"encoding/json"
	"io"
)

// IJSONDecoder ...
type IJSONDecoder interface {
	Decode(r io.Reader, obj interface{}) error
}

// JSONDecoder ...
type JSONDecoder struct{}

// NewJSONDecoder ...
func NewJSONDecoder() *JSONDecoder {
	return new(JSONDecoder)
}

// Decode ...
func (d *JSONDecoder) Decode(r io.Reader, obj interface{}) error {
	return json.NewDecoder(r).Decode(obj)
}
