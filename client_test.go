package chesscompubapi_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
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
			if _, err := w.Write([]byte(route.responseBody)); err != nil {
				w.WriteHeader(500)
				return
			}
			time.Sleep(route.requestDuration)
		})
	}

	server := httptest.NewServer(mux)
	return server
}

func runErrorTestWithTestServer(routes []testServerRoute, f func(c *chesscompubapi.Client) error, t *testing.T) {
	server := newTestServer(routes)
	defer server.Close()
	c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

	err := f(c)

	if err == nil {
		t.Error("expected err")
	}
}

func runOutputTestWithTestServer[T any](routes []testServerRoute, f func(c *chesscompubapi.Client) (T, error), want T, t *testing.T) {
	server := newTestServer(routes)
	defer server.Close()
	c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

	got, err := f(c)

	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
		return
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWithTimeout_ShouldTimeout(t *testing.T) {
	server := newTestServer([]testServerRoute{
		{
			pattern:         "/pub/player/henk/games/archives",
			requestDuration: time.Minute,
		},
	})
	defer server.Close()
	client := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL), chesscompubapi.WithTimeout(time.Nanosecond))
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
