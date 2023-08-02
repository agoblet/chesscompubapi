package chesscompubapi

import (
	"fmt"
)

type PlayerProfile struct {
	URL         string               `json:"url"`
	Username    string               `json:"username"`
	PlayerId    int                  `json:"player_id"`
	Title       *string              `json:"title"`
	Status      string               `json:"status"`
	Name        string               `json:"name"`
	Avatar      string               `json:"avatar"`
	Location    string               `json:"location"`
	CountryCode StringFromPathSuffix `json:"country"`
	Joined      UnixSecondsTimestamp `json:"joined"`
	LastOnline  UnixSecondsTimestamp `json:"last_online"`
	Followers   int                  `json:"followers"`
	IsStreamer  bool                 `json:"is_streamer"`
	Verified    bool                 `json:"verified"`
	League      string               `json:"league"`
	TwitchURL   string               `json:"twitch_url"`
	FIDE        int                  `json:"fide"`
}

type PlayerClub struct {
	URL          string               `json:"url"`
	Name         string               `json:"name"`
	Joined       UnixSecondsTimestamp `json:"joined"`
	LastActivity UnixSecondsTimestamp `json:"last_activity"`
	Icon         string               `json:"icon"`
	ID           StringFromPathSuffix `json:"@id"`
}

type PlayerStats struct {
	ChessDaily    *PlayerGameTypeStats      `json:"chess_daily"`
	Chess960Daily *PlayerGameTypeStats      `json:"chess960_daily"`
	ChessBlitz    *PlayerGameTypeStats      `json:"chess_blitz"`
	ChessBullet   *PlayerGameTypeStats      `json:"chess_bullet"`
	ChessRapid    *PlayerGameTypeStats      `json:"chess_rapid"`
	Tactics       *PlayerHighestLowestStats `json:"tactics"`
	Lessons       *PlayerHighestLowestStats `json:"lessons"`
	PuzzleRush    *PlayerPuzzleRushStats    `json:"puzzle_rush"`
	FIDE          int                       `json:"fide"`
}

type PlayerGameTypeStats struct {
	Last       LastPlayerGameTypeStats       `json:"last"`
	Best       *BestPlayerGameTypeStats      `json:"best"`
	Record     RecordPlayerGameTypeStats     `json:"record"`
	Tournament TournamentPlayerGameTypeStats `json:"tournament"`
}

type LastPlayerGameTypeStats struct {
	Date   UnixSecondsTimestamp `json:"date"`
	Rating int                  `json:"rating"`
	RD     int                  `json:"rd"`
}

type BestPlayerGameTypeStats struct {
	Date   UnixSecondsTimestamp `json:"date"`
	Rating int                  `json:"rating"`
	Game   string               `json:"game"`
}

type RecordPlayerGameTypeStats struct {
	Win            int                `json:"win"`
	Loss           int                `json:"loss"`
	Draw           int                `json:"draw"`
	TimePerMove    *DurationInSeconds `json:"time_per_move"`
	TimeoutPercent *float64           `json:"timeout_percent"`
}

type TournamentPlayerGameTypeStats struct {
	Count         int `json:"count"`
	Withdraw      int `json:"withdraw"`
	Points        int `json:"points"`
	HighestFinish int `json:"highest_finish"`
}

type PlayerHighestLowestStats struct {
	Highest DateRating `json:"highest"`
	Lowest  DateRating `json:"lowest"`
}

type DateRating struct {
	Date   UnixSecondsTimestamp `json:"date"`
	Rating int                  `json:"rating"`
}

type PlayerPuzzleRushStats struct {
	Daily *TotalAttemptsScore `json:"daily"`
	Best  TotalAttemptsScore  `json:"best"`
}

type TotalAttemptsScore struct {
	TotalAttempts int `json:"total_attempts"`
	Score         int `json:"score"`
}

// GetPlayerProfile gets the profile of a player.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-player.
func (c *Client) GetPlayerProfile(username string) (PlayerProfile, error) {
	const urlTemplate = "player/%s"
	profile := PlayerProfile{}
	err := c.getInto(fmt.Sprintf(urlTemplate, username), &profile)
	return profile, err
}

// ListPlayerClubs lists the clubs that a player is member of.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-player-clubs.
func (c *Client) ListPlayerClubs(username string) ([]PlayerClub, error) {
	const urlTemplate = "player/%s/clubs"
	clubs := &struct {
		Clubs []PlayerClub `json:"clubs"`
	}{}
	err := c.getInto(fmt.Sprintf(urlTemplate, username), clubs)
	return clubs.Clubs, err
}

// GetPlayerStats gets the profile of a player.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-player-stats.
func (c *Client) GetPlayerStats(username string) (PlayerStats, error) {
	const urlTemplate = "player/%s/stats"
	stats := PlayerStats{}
	err := c.getInto(fmt.Sprintf(urlTemplate, username), &stats)
	return stats, err
}
