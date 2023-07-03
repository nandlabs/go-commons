package messaging

import (
	"bytes"
	"io"
	"reflect"

	"go.nandlabs.io/commons/codec"
)

type LocalMessage struct {
	headers     map[string]interface{}
	headerTypes map[string]reflect.Kind
	body        *bytes.Buffer
}

func (lm *LocalMessage) SetBodyStr(input string) (n int, err error) {
	n, err = lm.body.WriteString(input)
	return
}

func (lm *LocalMessage) SetBodyBytes(input []byte) (n int, err error) {
	n, err = lm.body.Write(input)
	return
}

func (lm *LocalMessage) SetFrom(content io.Reader) (n int64, err error) {
	n, err = io.Copy(lm.body, content)
	return
}

func (lm *LocalMessage) WriteJSON(input interface{}) (err error) {
	err = lm.WriteContent(input, codec.JSON)
	return
}

func (lm *LocalMessage) WriteXML(input interface{}) (err error) {
	err = lm.WriteContent(input, codec.XML)
	return
}

func (lm *LocalMessage) WriteContent(input interface{}, contentType string) (err error) {
	var cdc codec.Codec
	// TODO : provide options to customise codec options
	cdc, err = codec.GetDefault(contentType)
	if err == nil {
		err = cdc.Write(input, lm.body)
	}
	return
}

func (lm *LocalMessage) ReadBody() io.Reader {
	return lm.body
}

func (lm *LocalMessage) ReadBytes() []byte {
	return lm.body.Bytes()
}

func (lm *LocalMessage) ReadAsStr() string {
	return lm.body.String()
}

func (lm *LocalMessage) ReadJSON(out interface{}) (err error) {
	err = lm.ReadContent(out, codec.JSON)
	return
}

func (lm *LocalMessage) ReadXML(out interface{}) (err error) {
	err = lm.ReadContent(out, codec.XML)
	return
}

func (lm *LocalMessage) ReadContent(out interface{}, contentType string) (err error) {
	var cdc codec.Codec
	// TODO: provide options to customise codec options
	cdc, err = codec.GetDefault(contentType)
	if err == nil {
		err = cdc.Read(lm.body, out)
	}
	return
}

func (lm *LocalMessage) SetHeader(key string, value []byte) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Array
}

func (lm *LocalMessage) SetStrHeader(key string, value string) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.String
}

func (lm *LocalMessage) SetBoolHeader(key string, value bool) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Bool
}

func (lm *LocalMessage) SetIntHeader(key string, value int) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Int
}

func (lm *LocalMessage) SetInt8Header(key string, value int8) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Int8
}

func (lm *LocalMessage) SetInt16Header(key string, value int16) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Int16
}

func (lm *LocalMessage) SetInt32Header(key string, value int32) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Int32
}

func (lm *LocalMessage) SetInt64Header(key string, value int64) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Int64
}

func (lm *LocalMessage) SetFloatHeader(key string, value float32) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Float32
}

func (lm *LocalMessage) SetFloat64Header(key string, value float64) {
	lm.headers[key] = value
	lm.headerTypes[key] = reflect.Float64
}

func (lm *LocalMessage) GetHeader(key string) (value []byte, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.([]byte)
	}
	return
}

func (lm *LocalMessage) GetStrHeader(key string) (value string, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(string)
	}
	return
}

func (lm *LocalMessage) GetBoolHeader(key string) (value bool, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(bool)
	}
	return
}

func (lm *LocalMessage) GetIntHeader(key string) (value int, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(int)
	}
	return
}

func (lm *LocalMessage) GetInt8Header(key string) (value int8, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(int8)
	}
	return
}

func (lm *LocalMessage) GetInt16Header(key string) (value int16, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(int16)
	}
	return
}

func (lm *LocalMessage) GetInt32Header(key string) (value int32, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(int32)
	}
	return
}

func (lm *LocalMessage) GetInt64Header(key string) (value int64, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(int64)
	}
	return
}

func (lm *LocalMessage) GetFloatHeader(key string) (value float32, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(float32)
	}
	return
}

func (lm *LocalMessage) GetFloat64Header(key string) (value float64, exists bool) {
	var v interface{}
	v, exists = lm.headers[key]
	if exists {
		value = v.(float64)
	}
	return
}
