package messaging

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestGetMessagingManager(t *testing.T) {
	tests := []struct {
		name string
		want ManagerMessaging
	}{
		// TODO: Add test cases.
		{
			name: "Test_1",
			want: &messagingSystems{
				messagingSystems: map[string]Messaging{"chan": &LocalMessagingSystem{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessagingManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMessagingManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessagingSystems_IsSupported(t *testing.T) {
	var scheme string
	var output bool

	testManager := GetMessagingManager()

	scheme = "chan"
	output = testManager.IsSupported(scheme)
	if output == false {
		t.Error()
	}
	scheme = "test"
	output = testManager.IsSupported(scheme)
	if output == true {
		t.Error()
	}
}

func TestMessagingSystems_Schemes(t *testing.T) {
	testManager := GetMessagingManager()

	output := testManager.Schemes()
	if output[0] != "chan" {
		t.Errorf("Schemes() default scheme added is file")
	}
}

func TestMessagingSystems_Send(t *testing.T) {
	// Invalid messaging scheme test
	testManager := GetMessagingManager()
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
	uri, _ := url.Parse("test://localhost:8080")
	got := testManager.Send(uri, msg)
	want := errors.New(fmt.Sprintf("unsupported messaging scheme test for in the url %s", uri))
	if got.Error() != want.Error() {
		t.Errorf("Error got :: %v, want :: %v", got, want)
	}
}
