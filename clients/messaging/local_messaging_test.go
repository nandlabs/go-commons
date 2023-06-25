package messaging

import (
	"bytes"
	"net/url"
	"reflect"
	"testing"
)

func TestLocalMessagingSystem_Send(t *testing.T) {
	lms := &LocalMessagingSystem{}
	msg := &LocalMessage{
		headers:     make(map[string]interface{}),
		headerTypes: make(map[string]reflect.Kind),
		body:        &bytes.Buffer{},
	}
	input := "this is a test string"
	_, err := msg.SetBodyStr(input)
	if err != nil {
		t.Errorf("Error SetBodyStr:: %v", err)
	}
	uri, _ := url.Parse("chan://localhost:8080")
	got := lms.Send(uri, msg)
	if got != nil {
		t.Errorf("Error got :: %v", got)
	}
}
