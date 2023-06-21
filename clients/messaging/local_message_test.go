package messaging

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestLocalMessage_SetBodyBytes(t *testing.T) {
	message := LocalMessage{
		headers:     make(map[string]interface{}),
		headerTypes: make(map[string]reflect.Kind),
		body:        &bytes.Buffer{},
	}
	input := []byte("this is a test string")
	_, err := message.SetBodyBytes(input)
	if err != nil {
		t.Errorf("Error SetBodyBytes: %v", err)
	}
	res := message.ReadAsByte()
	if string(res) != string(input) {
		t.Errorf("Error ReadAsStr, want= %v, got= %v", input, res)
	}
}

func TestLocalMessage_SetBodyStr(t *testing.T) {
	message := LocalMessage{
		headers:     make(map[string]interface{}),
		headerTypes: make(map[string]reflect.Kind),
		body:        &bytes.Buffer{},
	}
	input := "this is a test string"
	_, err := message.SetBodyStr(input)
	if err != nil {
		t.Errorf("Error SetBodyBytes: %v", err)
	}
	res := message.ReadAsStr()
	if res != input {
		t.Errorf("Error ReadAsStr: %v", err)
	}
}

func TestLocalMessage_SetFrom(t *testing.T) {
	message := LocalMessage{
		headers:     make(map[string]interface{}),
		headerTypes: make(map[string]reflect.Kind),
		body:        &bytes.Buffer{},
	}
	input := strings.NewReader("some io.Reader stream to be read")
	_, err := message.SetFrom(input)
	if err != nil {
		t.Errorf("Error SetFrom: %v", err)
	}
	res := message.ReadBody()
	b1 := make([]byte, 2)
	r1, err := res.Read(b1)
	if r1 != len(b1) {
		t.Error("Error SetFrom")
	}
}

func TestLocalMessage_WriteJSON(t *testing.T) {
	message := LocalMessage{
		headers:     make(map[string]interface{}),
		headerTypes: make(map[string]reflect.Kind),
		body:        &bytes.Buffer{},
	}
	input := interface{}(`{"name":"Test","body":"Hello","time":123124124}`)
	err := message.WriteJSON(input)
	if err != nil {
		t.Errorf("Error WriteJSON: %v", err)
	}
	var got interface{}
	err = message.ReadJSON(&got)
	if !reflect.DeepEqual(input, got) {
		t.Errorf("Error ReadJSON: %v", err)
	}
}

//func TestLocalMessage_WriteXML(t *testing.T) {
//	message := LocalMessage{
//		headers:     make(map[string]interface{}),
//		headerTypes: make(map[string]reflect.Kind),
//		body:        &bytes.Buffer{},
//	}
//	input := interface{}(`<XMLMessage><name>Test</name><body>Hello</body><time>123124124</time></XMLMessage>`)
//	err := message.WriteXML(input)
//	if err != nil {
//		t.Errorf("Error WriteXML: %v", err)
//	}
//	var got interface{}
//	err = message.ReadXML(&got)
//	fmt.Println(got)
//	if !reflect.DeepEqual(input, got) {
//		t.Errorf("Error ReadXML: %v", err)
//	}
//}
