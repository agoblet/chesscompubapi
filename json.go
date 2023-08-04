package chesscompubapi

import (
	"encoding/json"
	"strings"
	"time"
)

// UnixSecondsTimestamp decodes Unix seconds timestamps in JSON documents.
type UnixSecondsTimestamp time.Time

// DurationInSeconds decodes durations in seconds in JSON documents.
type DurationInSeconds time.Duration

// StringFromPathSuffix decodes url paths in JSON documents where only the part after the last string is needed.
type StringFromPathSuffix string

// UnmarshalJSON unmarshals an integer representing Unix seconds into a time.Time.
func (t *UnixSecondsTimestamp) UnmarshalJSON(data []byte) error {
	var raw int64
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*t = UnixSecondsTimestamp(time.Unix(raw, 0))
	return nil
}

// UnmarshalJSON unmarshals an integer representing a duration in seconds into a time.Duration.
func (d *DurationInSeconds) UnmarshalJSON(data []byte) error {
	var raw int64
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*d = DurationInSeconds(time.Second * time.Duration(raw))
	return nil
}

// UnmarshalJSON extracts the part after the last slash from a path.
// For example, "path/to/some/id" would be unmarshaled as "id".
func (d *StringFromPathSuffix) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	parts := strings.Split(raw, "/")
	*d = StringFromPathSuffix(parts[len(parts)-1])
	return nil
}
