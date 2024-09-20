// file: mypackage_test.go
package task5

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAsyncHTTPRequest(t *testing.T) {
	// Mock the server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := &http.Client{Timeout: 3 * time.Second}

	// Test the asynchronous function
	status, err := AsyncHTTPRequest(server.URL, client)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if status != "200 OK" {
		t.Errorf("expected status 200 OK, got %v", status)
	}
}
