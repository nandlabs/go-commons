package codec

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type xmlRW struct {
}

func (x *xmlRW) Write(v interface{}, w io.Writer) error {
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

func (x *xmlRW) Read(r io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.New(fmt.Sprintf("xml input error: %d", err))
	}
	if errU := xml.Unmarshal(b, v); err != nil {
		return errors.New(fmt.Sprintf("xml unmarshal error: %d", errU))
	}
	return nil
}
