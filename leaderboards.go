package chesscompubapi

type Leaderboards struct {
	Daily             []LeaderboardEntry `json:"daily"`
	Daily960          []LeaderboardEntry `json:"daily960"`
	LiveRapid         []LeaderboardEntry `json:"live_rapid"`
	LiveBullet        []LeaderboardEntry `json:"live_bullet"`
	LiveBlitz         []LeaderboardEntry `json:"live_blitz"`
	LiveBlitz960      []LeaderboardEntry `json:"live_blitz960"`
	LiveBughouse      []LeaderboardEntry `json:"live_bughouse"`
	LiveCrazyhouse    []LeaderboardEntry `json:"live_crazyhouse"`
	LiveKingOfTheHill []LeaderboardEntry `json:"live_kingofthehill"`
	LiveThreeCheck    []LeaderboardEntry `json:"live_threecheck"`
	Tactics           []LeaderboardEntry `json:"tactics"`
	Rush              []LeaderboardEntry `json:"rush"`
	Battle            []LeaderboardEntry `json:"battle"`
}

type LeaderboardEntry struct {
	PlayerID   int                  `json:"player_id"`
	WinCount   int                  `json:"win_count"`
	LossCount  int                  `json:"loss_count"`
	DrawCount  int                  `json:"draw_count"`
	URL        string               `json:"url"`
	Title      *string              `json:"title"`
	Name       *string              `json:"name"`
	FlairCode  string               `json:"flair_code"`
	Avatar     string               `json:"avatar"`
	TrendScore Trend                `json:"trend_score"`
	TrendRank  Trend                `json:"trend_rank"`
	Status     string               `json:"status"`
	Country    StringFromPathSuffix `json:"country"`
	Username   string               `json:"username"`
	Score      int                  `json:"score"`
	Rank       int                  `json:"rank"`
}

type Trend struct {
	Direction int `json:"direction"`
	Delta     int `json:"delta"`
}

// GetLeaderboards gets information about top 50 player for daily and live games, tactics and lessons.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-leaderboards.
func (c *Client) GetLeaderboards() (Leaderboards, error) {
	const urlTemplate = "leaderboards"
	leaderboards := Leaderboards{}
	err := c.getInto(urlTemplate, &leaderboards)
	return leaderboards, err
}
