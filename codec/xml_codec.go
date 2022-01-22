package codec

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type XMLCodec struct {
	options map[string]interface{}
}

func NewXmlCodec(options map[string]interface{}) Codec {
	return BaseCodec{readerWriter: XmlRW(options)}
}

func XmlRW(options map[string]interface{}) *XMLCodec {
	return &XMLCodec{options: options}
}

func (x *XMLCodec) Write(v interface{}, w io.Writer) error {
	output, err := xml.Marshal(v)
	if err != nil {
		return errors.New(fmt.Sprintf("xml marshal error: %d", err))
	}
	_, errW := w.Write(output)
	if errW != nil {
		return errW
	}
	return nil
}

func (x *XMLCodec) Read(r io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.New(fmt.Sprintf("xml input error: %d", err))
	}
	if errU := xml.Unmarshal(b, v); err != nil {
		return errors.New(fmt.Sprintf("xml unmarshal error: %d", errU))
	}
	return nil
}
