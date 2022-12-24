package rest

import (
	"go.nandlabs.io/commons/codec"
	"net/http"
	"reflect"
	"testing"
)

var client = NewClient()

func TestNewClient(t *testing.T) {
	if reflect.TypeOf(client) != reflect.TypeOf(NewClient()) {
		t.Errorf("NewClient() = %v, want %v", client, NewClient())
	}
}

func TestClientOptions(t *testing.T) {
	gotReq := client.ReqTimeout(10)
	if reflect.TypeOf(client) != reflect.TypeOf(gotReq) {
		t.Errorf("NewClient() = %v, want %v", gotReq, client)
	}

	gotCodec := client.AddCodecOption(codec.PrettyPrint, true)
	if reflect.TypeOf(client) != reflect.TypeOf(gotCodec) {
		t.Errorf("NewClient() = %v, want %v", gotCodec, client)
	}

	gotIdle := client.IdleTimeout(2)
	if reflect.TypeOf(client) != reflect.TypeOf(gotIdle) {
		t.Errorf("NewClient() = %v, want %v", gotIdle, client)
	}

	gotHttpEmpty := client.ErrorOnHttpStatus()
	if reflect.TypeOf(client) != reflect.TypeOf(gotHttpEmpty) {
		t.Errorf("NewClient() = %v, want %v", gotHttpEmpty, client)
	}

	gotHttp := client.ErrorOnHttpStatus(200, 300, 404)
	if reflect.TypeOf(client) != reflect.TypeOf(gotHttp) {
		t.Errorf("NewClient() = %v, want %v", gotHttp, client)
	}

	gotMaxIdle := client.MaxIdle(3)
	if reflect.TypeOf(client) != reflect.TypeOf(gotMaxIdle) {
		t.Errorf("NewClient() = %v, want %v", gotMaxIdle, client)
	}

	gotMaxIdlePerHost := client.MaxIdlePerHost(4)
	if reflect.TypeOf(client) != reflect.TypeOf(gotMaxIdlePerHost) {
		t.Errorf("NewClient() = %v, want %v", gotMaxIdlePerHost, client)
	}

	gotSSlVerify := client.SSlVerify(false)
	if reflect.TypeOf(client) != reflect.TypeOf(gotSSlVerify) {
		t.Errorf("NewClient() = %v, want %v", gotSSlVerify, client)
	}

	gotEndProxy := client.UseEnvProxy("test.com", "test", "test")
	if gotEndProxy != nil {
		t.Errorf("NewClient() = %v, want %v", gotEndProxy, client)
	}

	gotRetry := client.Retry(3, 5)
	if reflect.TypeOf(client) != reflect.TypeOf(gotRetry) {
		t.Errorf("NewClient() = %v, want %v", gotRetry, client)
	}

	gotCircuitBreaker := client.UseCircuitBreaker(1, 2, 1, 3)
	if reflect.TypeOf(client) != reflect.TypeOf(gotCircuitBreaker) {
		t.Errorf("NewClient() = %v, want %v", gotCircuitBreaker, client)
	}
}

func TestClient_NewRequest(t *testing.T) {
	req := client.NewRequest("http://localhost:8080", http.MethodGet)
	want := &Request{
		url:    "http://localhost:8080",
		method: http.MethodGet,
	}
	if reflect.TypeOf(req) != reflect.TypeOf(want) {
		t.Errorf("NewRequest() = %v, want %v", req, want)
	}
}

func TestClient_Execute(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		method string
		input  interface{}
		want   interface{}
	}{
		{
			name:   "TestClient_1",
			url:    "localhost",
			method: "",
			input:  "",
			want:   "Get \"localhost\": unsupported protocol scheme \"\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := client.NewRequest(tt.url, tt.method)
			_, err := client.Execute(req)
			if tt.want != err.Error() {
				t.Errorf("Got: %s, want: %s", err, tt.want)
			}
		})
	}
}
