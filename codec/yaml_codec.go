package codec

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type YamlCodec struct {
	options map[string]interface{}
}

func NewYamlCodec(options map[string]interface{}) Codec {
	return BaseCodec{readerWriter: YamlRW(options)}
}

func YamlRW(options map[string]interface{}) *YamlCodec {
	return &YamlCodec{options: options}
}

func (y *YamlCodec) Write(v interface{}, w io.Writer) error {
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

func (y *YamlCodec) Read(r io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.New(fmt.Sprintf("xml input error: %d", err))
	}
	if errU := yaml.Unmarshal(b, v); err != nil {
		return errors.New(fmt.Sprintf("xml unmarshal error: %d", errU))
	}
	return nil
}
