package chesscompubapi_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/agoblet/chesscompubapi"
)

type testServerRoute struct {
	pattern, responseBody string
	statusCode            int
	requestDuration       time.Duration
}

func newTestServer(routes []testServerRoute) *httptest.Server {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.pattern, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(route.statusCode)
			w.Write([]byte(route.responseBody))
			time.Sleep(route.requestDuration)
		})
	}

	server := httptest.NewServer(mux)
	return server
}

func TestClient_ShouldTimeout(t *testing.T) {
	server := newTestServer([]testServerRoute{
		{
			pattern:         "/pub/player/henk/games/archives",
			requestDuration: time.Minute,
		},
	})
	defer server.Close()
	client := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL), chesscompubapi.WithHTTPClient(&http.Client{Timeout: time.Nanosecond}))

	_, err := client.ListArchives("henk")

	urlErr, ok := err.(*url.Error)
	if !ok {
		t.Error("expected *url.Error")
		return
	}
	if !urlErr.Timeout() {
		t.Error("expected timeout")
	}
}

func TestClient_ShouldFailOnStatusCode(t *testing.T) {
	server := newTestServer([]testServerRoute{
		{
			pattern:      "/pub/player/piet/games/2021/12",
			responseBody: "not found",
			statusCode:   404,
		},
	})
	defer server.Close()
	client := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))
	const wantErrMessage = "404 not found"

	_, err := client.ListGames(chesscompubapi.Archive{
		Username: "piet", Year: 2021, Month: 12,
	})

	httpErr, ok := err.(*chesscompubapi.HTTPError)
	if !ok {
		t.Error("expected *chesscompubapi.HTTPError")
		return
	}
	if httpErr.Error() != wantErrMessage {
		t.Errorf("wrong error message, want '%s', got '%s'", wantErrMessage, httpErr.Error())
	}
}
