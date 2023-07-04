package chesscompubapi

import (
	"encoding/json"
	"strings"
	"time"
)

// UnixSecondsTimestamp decodes Unix seconds timestamps in JSON documents.
type UnixSecondsTimestamp struct {
	time.Time
}

// DurationInSeconds decodes durations in seconds in JSON documents.
type DurationInSeconds struct {
	time.Duration
}

func (t *UnixSecondsTimestamp) UnmarshalJSON(data []byte) error {
	var raw int64
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.Time = time.Unix(raw, 0)
	return nil
}

func (d *DurationInSeconds) UnmarshalJSON(data []byte) error {
	var raw int64
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	d.Duration = time.Second * time.Duration(raw)
	return nil
}

func partAfterLastSlash(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
