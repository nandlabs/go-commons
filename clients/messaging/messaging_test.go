package messaging

import (
	"bytes"
	"testing"
)

func TestMessage_SetBody(t *testing.T) {
	message := Message{
		Headers: map[string]interface{}{
			"Content-Type": "text/plain",
		},
		headerTypes: map[string]string{
			"Content-Type": "string",
		},
	}
	err := message.SetBody(bytes.NewBufferString("This is the message body"))
	if err != nil {
		t.Errorf("error setting body: %v", err)
	}
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(message.GetBody())
	if err != nil {
		t.Errorf("error reading body: %v", err)
		return
	}
	want := bytes.NewBufferString("This is the message body")
	if want.String() != buffer.String() {
		t.Errorf("invalid body, Got:: %v, Want:: %v", buffer.String(), want.String())
	}
}

func TestMessage_GetHeader(t *testing.T) {
	// Create a new Message instance for testing
	message := Message{
		Headers: map[string]interface{}{
			"Content-Type":   "text/plain",
			"X-Request-ID":   "123456789",
			"Custom-Header":  []byte{0x61, 0x62, 0x63},
			"Custom-Header2": 42,
		},
	}

	message.SetHeader("Custom-Header3", true)

	tests := []struct {
		key          string
		expectedVal  interface{}
		expectedBool bool
	}{
		{"Content-Type", "text/plain", true},
		{"X-Request-ID", "123456789", true},
		{"Custom-Header", []byte{0x61, 0x62, 0x63}, true},
		{"Custom-Header2", 42, true},
		{"Custom-Header3", true, true},
		{"Non-Existent-Header", nil, false},
	}

	for _, test := range tests {
		value, exists := message.GetHeader(test.key)

		// Check if the returned value matches the expected value
		switch expectedVal := test.expectedVal.(type) {
		case string:
			if value.(string) != expectedVal {
				t.Errorf("GetHeader(%s) returned value %v, expected %v", test.key, value, expectedVal)
			}
		case []byte:
			if string(value.([]byte)) != string(expectedVal) {
				t.Errorf("GetHeader(%s) returned value %v, expected %v", test.key, value, expectedVal)
			}
		default:
			if value != expectedVal {
				t.Errorf("GetHeader(%s) returned value %v, expected %v", test.key, value, expectedVal)
			}
		}

		// Check if the returned existence flag matches the expected flag
		if exists != test.expectedBool {
			t.Errorf("GetHeader(%s) returned exists %v, expected %v", test.key, exists, test.expectedBool)
		}
	}
}
