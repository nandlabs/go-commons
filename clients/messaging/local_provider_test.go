package messaging

import (
	"bytes"
	"net/url"
	"reflect"
	"testing"
)

func TestLocalProvider_Send(t *testing.T) {
	lms := NewMessaging()
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

	uriErr, _ := url.Parse("http://localhost:8080")
	got = lms.Send(uriErr, msg)
	if got.Error() != "unsupported provider with url http://localhost:8080" {
		t.Errorf("Error got :: %v", got)
	}
}
func TestLocalProvider_SendBatch(t *testing.T) {
	lms := NewMessaging()
	msg1 := &LocalMessage{
		headers:     make(map[string]interface{}),
		headerTypes: make(map[string]reflect.Kind),
		body:        &bytes.Buffer{},
	}
	input1 := "this is a test string 1"
	_, err := msg1.SetBodyStr(input1)
	if err != nil {
		t.Errorf("Error SetBodyStr:: %v", err)
	}
	msg2 := &LocalMessage{
		headers:     make(map[string]interface{}),
		headerTypes: make(map[string]reflect.Kind),
		body:        &bytes.Buffer{},
	}
	input2 := "this is a test string 2"
	_, err = msg2.SetBodyStr(input2)
	if err != nil {
		t.Errorf("Error SetBodyStr:: %v", err)
	}
	uri, _ := url.Parse("chan://localhost:8080")
	got := lms.SendBatch(uri, msg1, msg2)
	if got != nil {
		t.Errorf("Error got :: %v", got)
	}
}
