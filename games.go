package chesscompubapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Archive represents a monthly archive of games of a player.
type Archive struct {
	Username string
	Year     int
	Month    time.Month
}

// Represents a finished game played by 2 players.
type Game struct {
	Black        GamePlayer  `json:"black"`
	FEN          string      `json:"fen"`
	Accuracies   *Accuracies `json:"accuracies"`
	White        GamePlayer  `json:"white"`
	UUID         string      `json:"uuid"`
	PGN          string      `json:"pgn"`
	TCN          string      `json:"tcn"`
	URL          string      `json:"url"`
	TimeControl  string      `json:"time_control"`
	Rules        string      `json:"rules"`
	InitialSetup string      `json:"initial_setup"`
	TimeClass    string      `json:"time_class"`
	ECO          *string     `json:"eco"`
	Match        *string     `json:"match"`
	Tournament   *string     `json:"tournament"`
	Rated        bool        `json:"rated"`
	StartTime    time.Time   `json:"-"`
	EndTime      time.Time   `json:"-"`
}

// GamePlayer represents one of the 2 players of a chess game.
type GamePlayer struct {
	Username string `json:"username"`
	Rating   int    `json:"rating"`
	UUID     string `json:"uuid"`
	Result   string `json:"result"`
}

// Accuracies represents the accuracies of the 2 players of a chess game.
type Accuracies struct {
	Black float64 `json:"black"`
	White float64 `json:"white"`
}

// ListArchives lists all Archives available for a player.
func (c *Client) ListArchives(username string) ([]Archive, error) {
	const urlTemplate = "player/%s/games/archives"
	body, err := c.get(fmt.Sprintf(urlTemplate, username))
	if err != nil {
		return nil, err
	}

	archiveURLs := &struct {
		URLs []string `json:"archives"`
	}{}
	if err := json.Unmarshal(body, archiveURLs); err != nil {
		return nil, err
	}

	archives := make([]Archive, len(archiveURLs.URLs))
	for i, url := range archiveURLs.URLs {
		archives[i], err = url2Archive(url, username)
		if err != nil {
			return nil, err
		}
	}

	return archives, nil
}

// ListGames lists all Games available in an archive.
func (c *Client) ListGames(archive Archive) ([]Game, error) {
	const urlTemplate = "player/%s/games/%d/%02d"
	body, err := c.get(fmt.Sprintf(urlTemplate, archive.Username, archive.Year, archive.Month))
	if err != nil {
		return nil, err
	}

	timestamps := &struct {
		Games []struct {
			StartTime int64 `json:"start_time"`
			EndTime   int64 `json:"end_time"`
		} `json:"games"`
	}{}
	if err := json.Unmarshal(body, timestamps); err != nil {
		return nil, err
	}

	games := &struct {
		Games []Game `json:"games"`
	}{}
	if err := json.Unmarshal(body, games); err != nil {
		return nil, err
	}
	for i, g := range timestamps.Games {
		games.Games[i].StartTime = time.Unix(g.StartTime, 0)
		games.Games[i].EndTime = time.Unix(g.EndTime, 0)
	}

	return games.Games, nil
}

func url2Archive(url, username string) (Archive, error) {
	parts := strings.Split(url, "/")
	year, err := strconv.Atoi(parts[len(parts)-2])
	if err != nil {
		return Archive{}, err
	}
	month, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return Archive{}, err
	}

	return Archive{
		Username: username,
		Year:     year,
		Month:    time.Month(month),
	}, nil
}
