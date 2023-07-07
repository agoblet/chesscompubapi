package chesscompubapi_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/agoblet/chesscompubapi"
)

func TestStringFromPathSuffix(t *testing.T) {
	body := []byte(`"this/is/a/path"`)
	const want = "path"

	var got chesscompubapi.StringFromPathSuffix
	err := json.Unmarshal(body, &got)

	if err != nil {
		t.Errorf("want err nil, got %v", err)
		return
	}
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestDurationInSeconds(t *testing.T) {
	body := []byte("1337")
	const want = chesscompubapi.DurationInSeconds(time.Second * 1337)

	var got chesscompubapi.DurationInSeconds
	err := json.Unmarshal(body, &got)

	if err != nil {
		t.Errorf("want err nil, got %v", err)
		return
	}
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestUnixSecondsTimestamp(t *testing.T) {
	body := []byte("1636636502")
	want := chesscompubapi.UnixSecondsTimestamp(time.Unix(1636636502, 0))

	var got chesscompubapi.UnixSecondsTimestamp
	err := json.Unmarshal(body, &got)

	if err != nil {
		t.Errorf("want err nil, got %v", err)
		return
	}
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
