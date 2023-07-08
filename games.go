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
	Black        GamePlayer           `json:"black"`
	FEN          string               `json:"fen"`
	Accuracies   *Accuracies          `json:"accuracies"`
	White        GamePlayer           `json:"white"`
	UUID         string               `json:"uuid"`
	PGN          string               `json:"pgn"`
	TCN          string               `json:"tcn"`
	URL          string               `json:"url"`
	TimeControl  string               `json:"time_control"`
	Rules        string               `json:"rules"`
	InitialSetup string               `json:"initial_setup"`
	TimeClass    string               `json:"time_class"`
	ECO          *string              `json:"eco"`
	Match        *string              `json:"match"`
	Tournament   *string              `json:"tournament"`
	Rated        bool                 `json:"rated"`
	StartTime    UnixSecondsTimestamp `json:"start_time"`
	EndTime      UnixSecondsTimestamp `json:"end_time"`
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
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-games-archive-list.
func (c *Client) ListArchives(username string) ([]Archive, error) {
	const urlTemplate = "player/%s/games/archives"
	archives := &struct {
		Archives []Archive `json:"archives"`
	}{}
	err := c.getInto(fmt.Sprintf(urlTemplate, username), archives)
	return archives.Archives, err
}

// ListGames lists all Games available in an archive.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-games-archive.
func (c *Client) ListGames(archive Archive) ([]Game, error) {
	const urlTemplate = "player/%s/games/%d/%02d"
	games := &struct {
		Games []Game `json:"games"`
	}{}
	err := c.getInto(fmt.Sprintf(urlTemplate, archive.Username, archive.Year, archive.Month), games)
	return games.Games, err
}

// GetPGN lists all Games available in an archive.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-games-pgn.
func (c *Client) GetPGN(archive Archive) (string, error) {
	const urlTemplate = "player/%s/games/%d/%02d/pgn"
	pgn, err := c.get(fmt.Sprintf(urlTemplate, archive.Username, archive.Year, archive.Month))
	if err != nil {
		return "", err
	}
	return string(pgn), nil
}

// UnmarshalJSON unmarshals an archive URL from a JSON document into an Archive.
func (a *Archive) UnmarshalJSON(data []byte) error {
	var url string
	if err := json.Unmarshal(data, &url); err != nil {
		return err
	}

	parts := strings.Split(url, "/")
	year, err := strconv.Atoi(parts[len(parts)-2])
	if err != nil {
		return err
	}
	month, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return err
	}

	a.Username = parts[len(parts)-4]
	a.Year = year
	a.Month = time.Month(month)

	return nil
}
