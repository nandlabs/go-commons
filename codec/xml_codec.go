package codec

import (
	"encoding/xml"
	"io"
)

type xmlRW struct {
	options map[string]interface{}
}

func (x *xmlRW) Write(v interface{}, w io.Writer) error {
	encoder := xml.NewEncoder(w)
	return encoder.Encode(v)

}

func (x *xmlRW) Read(r io.Reader, v interface{}) error {
	decoder := xml.NewDecoder(r)
	return decoder.Decode(v)
}
