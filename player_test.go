package chesscompubapi_test

import (
	"testing"
	"time"

	"github.com/agoblet/chesscompubapi"
)

func TestGetPlayerProfile_ShouldGetProfile(t *testing.T) {
	tests := []struct {
		giveUsername, givePattern, giveResponseBody string
		want                                        chesscompubapi.PlayerProfile
	}{
		{
			giveUsername: "hikaru",
			givePattern:  "/pub/player/hikaru",
			giveResponseBody: `{
				"avatar":"https://images.chesscomfiles.com/uploads/v1/user/15448422.da101127.200x200o.64ebfda98bc2.png",
				"player_id":15448422,
				"@id":"https://api.chess.com/pub/player/hikaru",
				"url":"https://www.chess.com/member/Hikaru",
				"name":"Hikaru Nakamura",
				"username":"hikaru",
				"title":"GM",
				"followers":1161658,
				"country":"https://api.chess.com/pub/country/US",
				"location":"Sunrise, Florida",
				"last_online":1687899135,
				"joined":1389043258,
				"status":"premium",
				"is_streamer":true,
				"twitch_url":"https://twitch.tv/gmhikaru",
				"verified":false,
				"league":"Legend"
			}`,
			want: chesscompubapi.PlayerProfile{
				Avatar:      "https://images.chesscomfiles.com/uploads/v1/user/15448422.da101127.200x200o.64ebfda98bc2.png",
				PlayerId:    15448422,
				URL:         "https://www.chess.com/member/Hikaru",
				Name:        "Hikaru Nakamura",
				Username:    "hikaru",
				Title:       "GM",
				Followers:   1161658,
				CountryCode: "US",
				Location:    "Sunrise, Florida",
				LastOnline:  chesscompubapi.UnixSecondsTimestamp(time.Unix(1687899135, 0)),
				Joined:      chesscompubapi.UnixSecondsTimestamp(time.Unix(1389043258, 0)),
				Status:      "premium",
				IsStreamer:  true,
				TwitchURL:   "https://twitch.tv/gmhikaru",
				Verified:    false,
				League:      "Legend",
			},
		},
		{
			giveUsername: "emptycountry",
			givePattern:  "/pub/player/emptycountry",
			giveResponseBody: `{
				"country":"",
				"last_online":1687899135,
				"joined":1389043258
			}`,
			want: chesscompubapi.PlayerProfile{
				LastOnline: chesscompubapi.UnixSecondsTimestamp(time.Unix(1687899135, 0)),
				Joined:     chesscompubapi.UnixSecondsTimestamp(time.Unix(1389043258, 0)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveUsername, func(t *testing.T) {
			runOutputTestWithTestServer(
				[]testServerRoute{
					{
						pattern:      tt.givePattern,
						responseBody: tt.giveResponseBody,
						statusCode:   200,
					},
				},
				func(c *chesscompubapi.Client) (any, error) { return c.GetPlayerProfile(tt.giveUsername) },
				tt.want,
				t,
			)
		})
	}
}

func TestGetPlayerProfile_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "404",
			giveResponseBody: "not found",
			giveStatusCode:   404,
		},
		{
			name:             "corruptTimestamp",
			giveResponseBody: `{"joined": "nan"}`,
			giveStatusCode:   200,
		},
		{
			name:             "corruptPlayerId",
			giveResponseBody: `{"player_id": "nan"}`,
			giveStatusCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runErrorTestWithTestServer([]testServerRoute{{
				pattern:      "/pub/player/johndoe",
				responseBody: tt.giveResponseBody,
				statusCode:   tt.giveStatusCode,
			}}, func(c *chesscompubapi.Client) error {
				_, err := c.GetPlayerProfile("johndoe")
				return err
			}, t)
		})
	}
}

func TestGetPlayerStats_ShouldGetStats(t *testing.T) {
	tests := []struct {
		giveUsername, givePattern, giveResponseBody string
		want                                        chesscompubapi.PlayerStats
	}{
		{
			giveUsername: "axelgoblet",
			givePattern:  "/pub/player/axelgoblet/stats",
			giveResponseBody: `{
				"chess_daily":{
					"last":{
						"rating":1016,
						"date":1636636502,
						"rd":129
					},
					"best":{
						"rating":1190,
						"date":1612811146,
						"game":"https://www.chess.com/game/daily/360173289"
					},
					"record":{
						"win":4,
						"loss":8,
						"draw":0,
						"time_per_move":10644,
						"timeout_percent":0
					}
				},
				"chess_rapid":{
					"last":{
						"rating":1302,
						"date":1688405035,
						"rd":46
					},
					"best":{
						"rating":1456,
						"date":1633173887,
						"game":"https://www.chess.com/game/live/27718858417"
					},
					"record":{
						"win":885,
						"loss":783,
						"draw":49
					}
				},
				"fide":0,
				"tactics":{
					"highest":{
						"rating":2225,
						"date":1688441463
					},
					"lowest":{
						"rating":398,
						"date":1609315550
					}
				},
				"puzzle_rush":{
					"best":{
						"total_attempts":30,
						"score":27
					},
					"daily":{
						"total_attempts":30,
						"score":27
					}
				}
			}`,
			want: chesscompubapi.PlayerStats{
				ChessDaily: &chesscompubapi.PlayerGameTypeStats{
					Last: chesscompubapi.LastPlayerGameTypeStats{
						Date:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1636636502, 0)),
						Rating: 1016,
						RD:     129,
					},
					Best: &chesscompubapi.BestPlayerGameTypeStats{
						Date:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1612811146, 0)),
						Rating: 1190,
						Game:   "https://www.chess.com/game/daily/360173289",
					},
					Record: chesscompubapi.RecordPlayerGameTypeStats{
						Win:            4,
						Loss:           8,
						Draw:           0,
						TimePerMove:    pointer(chesscompubapi.DurationInSeconds(time.Second * 10644)),
						TimeoutPercent: pointer(0.),
					},
				},
				ChessRapid: &chesscompubapi.PlayerGameTypeStats{
					Last: chesscompubapi.LastPlayerGameTypeStats{
						Date:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1688405035, 0)),
						Rating: 1302,
						RD:     46,
					},
					Best: &chesscompubapi.BestPlayerGameTypeStats{
						Date:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1633173887, 0)),
						Rating: 1456,
						Game:   "https://www.chess.com/game/live/27718858417",
					},
					Record: chesscompubapi.RecordPlayerGameTypeStats{
						Win:  885,
						Loss: 783,
						Draw: 49,
					},
				},
				Tactics: &chesscompubapi.PlayerHighestLowestStats{
					Highest: chesscompubapi.DateRating{
						Rating: 2225,
						Date:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1688441463, 0)),
					},
					Lowest: chesscompubapi.DateRating{
						Rating: 398,
						Date:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1609315550, 0)),
					},
				},
				PuzzleRush: &chesscompubapi.PlayerPuzzleRushStats{
					Best: chesscompubapi.TotalAttemptsScore{
						TotalAttempts: 30,
						Score:         27,
					},
					Daily: &chesscompubapi.TotalAttemptsScore{
						TotalAttempts: 30,
						Score:         27,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveUsername, func(t *testing.T) {
			runOutputTestWithTestServer(
				[]testServerRoute{
					{
						pattern:      tt.givePattern,
						responseBody: tt.giveResponseBody,
						statusCode:   200,
					},
				},
				func(c *chesscompubapi.Client) (any, error) { return c.GetPlayerStats(tt.giveUsername) },
				tt.want,
				t,
			)
		})
	}
}

func TestGetPlayerStats_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "corruptBody",
			giveResponseBody: `{"chess_daily": "a string instead of an object"}`,
			giveStatusCode:   200,
		},
		{
			name: "corruptDuration",
			giveResponseBody: `{
				"chess_daily":{
					"record":{
						"time_per_move":[10644]
					}
				}
			}`,
			giveStatusCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runErrorTestWithTestServer([]testServerRoute{{
				pattern:      "/pub/player/johndoe/stats",
				responseBody: tt.giveResponseBody,
				statusCode:   tt.giveStatusCode,
			}}, func(c *chesscompubapi.Client) error {
				_, err := c.GetPlayerStats("johndoe")
				return err
			}, t)
		})
	}
}

func TestListPlayerClubs_ShouldListClubs(t *testing.T) {
	tests := []struct {
		giveUsername, givePattern, giveResponseBody string
		want                                        []chesscompubapi.PlayerClub
	}{
		{
			giveUsername: "erik",
			givePattern:  "/pub/player/erik/clubs",
			giveResponseBody: `{
				"clubs":[
					{
						"@id":"https://www.chess.com/club/open-discussion",
						"name":"Open Discussion",
						"last_activity":1626902692,
						"icon":"https://images.chesscomfiles.com/uploads/v1/group/3541.45d6eb2c.50x50o.af307b09ebe7.png",
						"url":"https://www.chess.com/club/open-discussion",
						"joined":1468775412
					},
					{
						"@id":"https://www.chess.com/club/chesscom---tactics-trainer-approvers",
						"name":"Chess.com - Tactics Trainer Approvers",
						"last_activity":1626902680,
						"icon":"https://images.chesscomfiles.com/uploads/v1/group/4146.cd1f3309.50x50o.2daf5479afb9.gif",
						"url":"https://www.chess.com/club/chesscom---tactics-trainer-approvers",
						"joined":1210299490
					}
				]
			}`,
			want: []chesscompubapi.PlayerClub{
				{
					ID:           "open-discussion",
					Name:         "Open Discussion",
					LastActivity: chesscompubapi.UnixSecondsTimestamp(time.Unix(1626902692, 0)),
					Joined:       chesscompubapi.UnixSecondsTimestamp(time.Unix(1468775412, 0)),
					Icon:         "https://images.chesscomfiles.com/uploads/v1/group/3541.45d6eb2c.50x50o.af307b09ebe7.png",
					URL:          "https://www.chess.com/club/open-discussion",
				},
				{
					ID:           "chesscom---tactics-trainer-approvers",
					Name:         "Chess.com - Tactics Trainer Approvers",
					LastActivity: chesscompubapi.UnixSecondsTimestamp(time.Unix(1626902680, 0)),
					Joined:       chesscompubapi.UnixSecondsTimestamp(time.Unix(1210299490, 0)),
					Icon:         "https://images.chesscomfiles.com/uploads/v1/group/4146.cd1f3309.50x50o.2daf5479afb9.gif",
					URL:          "https://www.chess.com/club/chesscom---tactics-trainer-approvers",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveUsername, func(t *testing.T) {
			runOutputTestWithTestServer(
				[]testServerRoute{
					{
						pattern:      tt.givePattern,
						responseBody: tt.giveResponseBody,
						statusCode:   200,
					},
				},
				func(c *chesscompubapi.Client) (any, error) { return c.ListPlayerClubs(tt.giveUsername) },
				tt.want,
				t,
			)
		})
	}
}

func TestListPlayerClubs_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "404",
			giveResponseBody: "not found",
			giveStatusCode:   404,
		},
		{
			name:             "corruptTimestamp",
			giveResponseBody: `{"clubs":[{"joined": "nan"}]}`,
			giveStatusCode:   200,
		},
		{
			name:             "corruptURL",
			giveResponseBody: `{"clubs":[{"url": {"i": "am an object"}}]}`,
			giveStatusCode:   200,
		},
		{
			name:             "corruptID",
			giveResponseBody: `{"clubs":[{"@id": {"i": "am an object"}}]}`,
			giveStatusCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runErrorTestWithTestServer([]testServerRoute{{
				pattern:      "/pub/player/johndoe/clubs",
				responseBody: tt.giveResponseBody,
				statusCode:   tt.giveStatusCode,
			}}, func(c *chesscompubapi.Client) error {
				_, err := c.ListPlayerClubs("johndoe")
				return err
			}, t)
		})
	}
}

func pointer[T any](v T) *T {
	return &v
}
