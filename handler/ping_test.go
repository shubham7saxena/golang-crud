package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	PingHandler(rr, req)

	resp := rr.Result()
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("ping handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	want := `"ping": "pong"`
	body, _ := ioutil.ReadAll(resp.Body)
	got := string(body)
	if got != want {
		t.Errorf("ping handler returned unexpected body: got %v want %v", got, want)
	}
}
