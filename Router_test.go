package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewRouter(t *testing.T) {
	test := NewRouter()

	if len(test.rules) != 0 {
		t.Errorf("expected  struct")
	}
}

func FindHandler(t *testing.T) {
	test := NewRouter()

	if len(test.rules) != 0 {
		t.Errorf("expected  struct")
	}
}

func TestServeHTTP(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("asdf"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var testo Router
	testo.ServeHTTP(w, r)
}
