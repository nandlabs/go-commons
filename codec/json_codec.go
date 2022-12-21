package codec

import (
	"encoding/json"
	"io"

	"go.nandlabs.io/commons/codec/validator"
)

var structValidator = validator.NewStructValidator()

type jsonRW struct {
	options map[string]interface{}
}

func (c *jsonRW) Write(v interface{}, w io.Writer) error {
	//only utf-8 charset is supported
	var escapeHtml = false
	if c.options != nil {
		if v, ok := c.options[JsonEscapeHTML]; ok {
			escapeHtml = v.(bool)
		}

	}
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(escapeHtml)
	return encoder.Encode(v)

}

func (c *jsonRW) Read(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
