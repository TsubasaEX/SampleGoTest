package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleSlothfulMessage(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/sloth", nil)

	handleSlothfulMessage(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Stay slothful!") {
		t.Errorf(
			`response body "%s" does not contain "Stay slothful!"`,
			wr.Body.String(),
		)
	}
}

func TestGetSlothfulMessage(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/sloth", handleSlothfulMessage)

	svr := httptest.NewServer(router)
	defer svr.Close()

	c := NewClient(http.DefaultClient, svr.URL)
	m, err := c.GetSlothfulMessage()
	if err != nil {
		t.Fatalf("error in GetSlothfulMessage: %v", err)
	}
	if m.Message != "Stay slothful!" {
		t.Errorf(
			`message %s should contain string "Sloth"`,
			m.Message,
		)
	}
}
