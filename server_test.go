package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {

	t.Run("returns the score for Pepper", func(t *testing.T) {

		request := newGetStoreRequest("Pepper")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "20")

	})

	t.Run("returns the score for Floyd", func(t *testing.T) {

		request := newGetStoreRequest("Floyd")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "10")

	})

}

func newGetStoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
