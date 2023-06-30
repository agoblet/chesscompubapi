package chesscompubapi

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type PlayerProfile struct {
	URL         string    `json:"url"`
	Username    string    `json:"username"`
	PlayerId    int       `json:"player_id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Location    string    `json:"location"`
	CountryCode string    `json:"country"`
	Joined      time.Time `json:"-"`
	LastOnline  time.Time `json:"-"`
	Followers   int       `json:"followers"`
	IsStreamer  bool      `json:"is_streamer"`
	Verified    bool      `json:"verified"`
	League      string    `json:"league"`
	TwitchURL   string    `json:"twitch_url"`
	FIDE        int       `json:"fide"`
}

// GetPlayerProfile lists the profile of a player.
func (c *Client) GetPlayerProfile(username string) (PlayerProfile, error) {
	profile := &PlayerProfile{}

	const urlTemplate = "player/%s"
	body, err := c.get(fmt.Sprintf(urlTemplate, username))
	if err != nil {
		return *profile, err
	}

	timestamps := &struct {
		Joined     int64 `json:"joined"`
		LastOnline int64 `json:"last_online"`
	}{}
	if err := json.Unmarshal(body, timestamps); err != nil {
		return *profile, err
	}

	if err := json.Unmarshal(body, profile); err != nil {
		return *profile, err
	}
	profile.Joined = time.Unix(timestamps.Joined, 0)
	profile.LastOnline = time.Unix(timestamps.LastOnline, 0)
	profile.CountryCode = partAfterLastSlash(profile.CountryCode)

	return *profile, nil
}

func partAfterLastSlash(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
