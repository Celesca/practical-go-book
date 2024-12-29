package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
}

func TestFetchRemoteResource(t *testing.T) {
	server := startTestHTTPServer()
	defer server.Close()

	body, err := fetchRemoteResource(server.URL)
	if err != nil {
		t.Fatalf("Error fetching resource: %v", err)
	}

	if string(body) != "Hello, client\n" {
		t.Fatalf("Unexpected response: %s", body)
	}
}
