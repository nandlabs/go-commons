package codec

import (
	"encoding/json"
	"go.nandlabs.io/commons/codec/validator"
	"io"
)

var structValidator = validator.NewStructValidator()

type jsonRW struct {
}

func (c *jsonRW) Write(v interface{}, w io.Writer) error {

	return json.NewEncoder(w).Encode(v)

}

func (c *jsonRW) Read(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
