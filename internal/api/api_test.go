package api_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/murat/go-boilerplate/internal/api"
)

func TestClient_Get(t *testing.T) {
	helloResponse, err := os.ReadFile("./responses/hello.json")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		word    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "success",
			word:    "hello",
			want:    helloResponse,
			wantErr: false,
		},
		{
			name:    "fail",
			word:    "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Add("content-type", "application/json")
				if tt.wantErr {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write(tt.want)
			}))
			defer srv.Close()

			c := api.NewClient(srv.Client(), srv.URL, "xxx")
			got, err := c.Get(tt.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.Get() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
