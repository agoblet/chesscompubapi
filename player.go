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

type PlayerClub struct {
	URL          string    `json:"url"`
	Name         string    `json:"name"`
	Joined       time.Time `json:"-"`
	LastActivity time.Time `json:"-"`
	Icon         string    `json:"icon"`
	ID           string    `json:"@id"`
}

// GetPlayerProfile gets the profile of a player.
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

// ListPlayerClubs lists the clubs that a player is member of.
func (c *Client) ListPlayerClubs(username string) ([]PlayerClub, error) {
	const urlTemplate = "player/%s/clubs"
	body, err := c.get(fmt.Sprintf(urlTemplate, username))
	if err != nil {
		return nil, err
	}

	timestamps := &struct {
		Clubs []struct {
			Joined       int64 `json:"joined"`
			LastActivity int64 `json:"last_activity"`
		} `json:"clubs"`
	}{}
	if err := json.Unmarshal(body, timestamps); err != nil {
		return nil, err
	}

	clubs := &struct {
		Clubs []PlayerClub `json:"clubs"`
	}{}
	if err := json.Unmarshal(body, clubs); err != nil {
		return nil, err
	}
	for i, c := range timestamps.Clubs {
		clubs.Clubs[i].Joined = time.Unix(c.Joined, 0)
		clubs.Clubs[i].LastActivity = time.Unix(c.LastActivity, 0)
	}
	for i, c := range clubs.Clubs {
		clubs.Clubs[i].ID = partAfterLastSlash(c.ID)
	}

	return clubs.Clubs, nil
}

func partAfterLastSlash(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
