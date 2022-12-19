package codec

import (
	"encoding/json"
	"go.nandlabs.io/commons/codec/validator"
	"io"
)

var structValidator = validator.NewStructValidator()

type JsonCodec struct {
	//Json JsonCodec specific code will come here.
	options map[string]interface{}
}

func NewJsonCodec(options map[string]interface{}) Codec {
	return BaseCodec{readerWriter: JsonRW(options)}
}

func JsonRW(options map[string]interface{}) *JsonCodec {
	//Case the defn here and return the codec
	// base codec's reader writer
	return &JsonCodec{options: options}
}

func (c *JsonCodec) Write(v interface{}, w io.Writer) (err error) {
	// marshal wrapper
	// if the validation is successful then use the core json marshal to generate the json-codec from the struct and write it back to the buffer
	if err = structValidator.Validate(v); err == nil {
		err = json.NewEncoder(w).Encode(v)
	}
	return

}

func (c *JsonCodec) Read(r io.Reader, v interface{}) error {
	// unmarshal wrapper
	// read the data from reader and map it to the interface
	return json.NewDecoder(r).Decode(v)
}

// Commenting for now, to be used later for the info during caching, if required!!
/*func (c *JsonCodec) Load(v interface{}) error {
	// check if the input interface constraints are present in the cache or not
	// caching the struct constraints
	// map all the constraints
	// parse the constraints to save to the cache while the struct comes in
	// make sure the map is synchronised
	return errors.New("register is not implemented in base codec")
}*/
