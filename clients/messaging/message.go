package messaging

import (
	"bytes"
	"io"
)

type Message struct {
	Headers     map[string]interface{}
	headerTypes map[string]string
	Body        *bytes.Buffer
}

func (m *Message) SetBody(content io.Reader) (err error) {
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(content)
	if err != nil {
		return
	}
	m.Body = buffer
	return
}

func (m *Message) GetBody() io.Reader {
	return m.Body
}

func (m *Message) SetHeader(key string, value interface{}) {
	m.Headers[key] = value
}

func (m *Message) GetHeader(key string) (value interface{}, exists bool) {
	value, exists = m.Headers[key]
	return
}
