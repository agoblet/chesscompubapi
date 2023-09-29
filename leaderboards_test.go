package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestGetLeaderboards_ShouldGetLeaderboards(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/leaderboards",
				responseBody: `{
					"daily":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"daily960":[
						{
							"player_id":2305524,
							"name": "Henk",
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"live_rapid":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"live_bullet":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"live_blitz":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":1,
								"delta":2
							},
							"trend_rank":{
								"direction":3,
								"delta":4
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"live_blitz960":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"live_bughouse":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"live_threecheck":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"live_crazyhouse":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"live_kingofthehill":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"tactics":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"rush":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					],
					"battle":[
						{
							"player_id":2305524,
							"@id":"https://api.chess.com/pub/player/zgorl",
							"url":"https://www.chess.com/member/Zgorl",
							"username":"Zgorl",
							"score":2585,
							"rank":1,
							"country":"NL",
							"title":"FM",
							"status":"premium",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
							"trend_score":{
								"direction":0,
								"delta":0
							},
							"trend_rank":{
								"direction":0,
								"delta":0
							},
							"flair_code":"white_bishop",
							"win_count":454,
							"loss_count":204,
							"draw_count":57
						}
					]
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) (chesscompubapi.Leaderboards, error) {
			return c.GetLeaderboards()
		},
		chesscompubapi.Leaderboards{
			Daily: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			Daily960: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					Name:     pointer("Henk"),
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			LiveRapid: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			LiveBullet: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			LiveThreeCheck: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			LiveKingOfTheHill: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			LiveCrazyhouse: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			Rush: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			Tactics: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			Battle: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			LiveBughouse: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			LiveBlitz960: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 0,
						Delta:     0,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
			LiveBlitz: []chesscompubapi.LeaderboardEntry{
				{
					PlayerID: 2305524,
					URL:      "https://www.chess.com/member/Zgorl",
					Username: "Zgorl",
					Score:    2585,
					Rank:     1,
					Country:  "NL",
					Title:    pointer("FM"),
					Status:   "premium",
					Avatar:   "https://images.chesscomfiles.com/uploads/v1/user/2305524.5341b605.200x200o.67a89f6c51b4.jpeg",
					TrendScore: chesscompubapi.Trend{
						Direction: 1,
						Delta:     2,
					},
					TrendRank: chesscompubapi.Trend{
						Direction: 3,
						Delta:     4,
					},
					FlairCode: "white_bishop",
					WinCount:  454,
					LossCount: 204,
					DrawCount: 57,
				},
			},
		},
		t,
	)
}

func TestGetLeaderboards_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/leaderboards",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.GetLeaderboards()
		return err
	}, t)
}
