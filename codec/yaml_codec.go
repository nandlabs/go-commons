package codec

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type yamlRW struct {
}

func (y *yamlRW) Write(v interface{}, w io.Writer) error {
	output, err := yaml.Marshal(v)
	if err != nil {
		return errors.New(fmt.Sprintf("xml marshal error: %d", err))
	}
	_, errW := w.Write(output)
	if errW != nil {
		return errW
	}
	return nil
}

func (y *yamlRW) Read(r io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.New(fmt.Sprintf("xml input error: %d", err))
	}
	if errU := yaml.Unmarshal(b, v); err != nil {
		return errors.New(fmt.Sprintf("xml unmarshal error: %d", errU))
	}
	return nil
}
