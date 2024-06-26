package chesscompubapi_test

import (
	"testing"
	"time"

	"github.com/agoblet/chesscompubapi"
)

func TestGetClub_ShouldGetClub(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/club/chess-com-developer-community",
				responseBody: `{
					"@id":"https://api.chess.com/pub/club/chess-com-developer-community",
					"name":"Chess.com Developer Community",
					"club_id":57796,
					"country":"https://api.chess.com/pub/country/XX",
					"average_daily_rating":974,
					"members_count":10071,
					"created":1500301978,
					"last_activity":1682347892,
					"admin":[
						"https://api.chess.com/pub/player/erik",
						"https://api.chess.com/pub/player/andreamorandini",
						"https://api.chess.com/pub/player/bcurtis",
						"https://api.chess.com/pub/player/jdcannon"
					],
					"visibility":"public",
					"join_request":"https://www.chess.com/club/join/57796",
					"icon":"https://images.chesscomfiles.com/uploads/v1/group/57796.67ee0038.50x50o.585842f67281.png",
					"description":"<p>Chess.com's official club for APIs, data, and code for developers. Get involved!</p>\n<p>\u00a0</p>\n<p>Apply for OAuth access / Chess.com login here: <a href=\"https://forms.gle/RwGLuZkwDysCj2GV7\" target=\"_blank\" rel=\"noreferrer noopener\">https://forms.gle/RwGLuZkwDysCj2GV7</a></p>\n<p>\u00a0</p>\n<p>More resources:</p>\n<ul>\n<li><a href=\"https://chesscom.notion.site/Getting-started-with-Chess-com-OAuth-2-0-Server-5958e57c8c934a3aa7abda2d670969e8\">https://chesscom.notion.site/Getting-started-with-Chess-com-OAuth-2-0-Server-5958e57c8c934a3aa7abda2d670969e8</a></li>\n<li><a href=\"https://www.chess.com/news/view/published-data-api\">https://www.chess.com/news/view/published-data-api</a></li>\n<li><a href=\"https://www.npmjs.com/package/chess-web-api\" target=\"_blank\" rel=\"noreferrer noopener\">https://www.npmjs.com/package/chess-web-api</a></li>\n<li><a href=\"https://github.com/ChessCom/ios-chessclock\" target=\"_blank\" rel=\"noreferrer noopener\">https://github.com/ChessCom/ios-chessclock</a></li>\n<li><a href=\"https://github.com/ChessCom/android-chessclock\" target=\"_blank\" rel=\"noreferrer noopener\">https://github.com/ChessCom/android-chessclock</a></li>\n</ul>\n<p>\u00a0</p>\n<p>Our <a href=\"https://www.chess.com/news/view/chess-com-bug-bounty-policy\">Bug Bounty Policy</a> is available <a href=\"https://www.chess.com/news/view/chess-com-bug-bounty-policy\">here</a>.\u00a0</p>",
					"url":"https://www.chess.com/club/chess-com-developer-community"
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) (chesscompubapi.Club, error) {
			return c.GetClub("chess-com-developer-community")
		},
		chesscompubapi.Club{
			ID:                 "chess-com-developer-community",
			Name:               "Chess.com Developer Community",
			ClubID:             57796,
			Country:            "XX",
			AverageDailyRating: 974,
			MembersCount:       10071,
			Created:            chesscompubapi.UnixSecondsTimestamp(time.Unix(1500301978, 0)),
			LastActivity:       chesscompubapi.UnixSecondsTimestamp(time.Unix(1682347892, 0)),
			Admins: []chesscompubapi.StringFromPathSuffix{
				"erik",
				"andreamorandini",
				"bcurtis",
				"jdcannon",
			},
			Visibility:  "public",
			JoinRequest: "https://www.chess.com/club/join/57796",
			Icon:        "https://images.chesscomfiles.com/uploads/v1/group/57796.67ee0038.50x50o.585842f67281.png",
			Description: "<p>Chess.com's official club for APIs, data, and code for developers. Get involved!</p>\n<p>\u00a0</p>\n<p>Apply for OAuth access / Chess.com login here: <a href=\"https://forms.gle/RwGLuZkwDysCj2GV7\" target=\"_blank\" rel=\"noreferrer noopener\">https://forms.gle/RwGLuZkwDysCj2GV7</a></p>\n<p>\u00a0</p>\n<p>More resources:</p>\n<ul>\n<li><a href=\"https://chesscom.notion.site/Getting-started-with-Chess-com-OAuth-2-0-Server-5958e57c8c934a3aa7abda2d670969e8\">https://chesscom.notion.site/Getting-started-with-Chess-com-OAuth-2-0-Server-5958e57c8c934a3aa7abda2d670969e8</a></li>\n<li><a href=\"https://www.chess.com/news/view/published-data-api\">https://www.chess.com/news/view/published-data-api</a></li>\n<li><a href=\"https://www.npmjs.com/package/chess-web-api\" target=\"_blank\" rel=\"noreferrer noopener\">https://www.npmjs.com/package/chess-web-api</a></li>\n<li><a href=\"https://github.com/ChessCom/ios-chessclock\" target=\"_blank\" rel=\"noreferrer noopener\">https://github.com/ChessCom/ios-chessclock</a></li>\n<li><a href=\"https://github.com/ChessCom/android-chessclock\" target=\"_blank\" rel=\"noreferrer noopener\">https://github.com/ChessCom/android-chessclock</a></li>\n</ul>\n<p>\u00a0</p>\n<p>Our <a href=\"https://www.chess.com/news/view/chess-com-bug-bounty-policy\">Bug Bounty Policy</a> is available <a href=\"https://www.chess.com/news/view/chess-com-bug-bounty-policy\">here</a>.\u00a0</p>",
			URL:         "https://www.chess.com/club/chess-com-developer-community",
		},
		t,
	)
}

func TestGetClub_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/club/chess-com-developer-community",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.GetClub("chess-com-developer-community")
		return err
	}, t)
}

func TestGetClubMemberActivity_ShouldGetActivity(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/club/chess-com-developer-community/members",
				responseBody: `{
					"weekly":[
						{
							"username":"0cean",
							"joined":1675110044
						}
					],
					"monthly":[
						{
							"username":"0nepunchpawn",
							"joined":1626915696
						},
						{
							"username":"101anj101",
							"joined":1689197757
						}
					],
					"all_time":[
						{
							"username":"000_elite_warrior_000",
							"joined":1636391975
						}
					]
				 }`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) (chesscompubapi.ClubMemberActivity, error) {
			return c.GetClubMemberActivity("chess-com-developer-community")
		},
		chesscompubapi.ClubMemberActivity{
			Weekly: []chesscompubapi.ClubMember{
				{
					Username: "0cean",
					Joined:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1675110044, 0)),
				},
			},
			Monthly: []chesscompubapi.ClubMember{
				{
					Username: "0nepunchpawn",
					Joined:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1626915696, 0)),
				},
				{
					Username: "101anj101",
					Joined:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1689197757, 0)),
				},
			},
			AllTime: []chesscompubapi.ClubMember{
				{
					Username: "000_elite_warrior_000",
					Joined:   chesscompubapi.UnixSecondsTimestamp(time.Unix(1636391975, 0)),
				},
			},
		},
		t,
	)
}

func TestGetClubMemberActivity_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/club/chess-com-developer-community/members",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.GetClubMemberActivity("chess-com-developer-community")
		return err
	}, t)
}

func TestGetClubMatches_ShouldGetClubMatches(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/club/chess-com-developer-community/matches",
				responseBody: `{
					"finished":[
						{
							"name":"0cean",
							"start_time":1675110044,
							"opponent":"https://api.chess.com/pub/club/team-argentina-b",
							"time_class":"daily",
							"result":"win",
							"@id":"https://api.chess.com/pub/match/796340"
						}
					],
					"in_progress":[
						{
							"name":"0cean",
							"start_time":1675110044,
							"opponent":"https://api.chess.com/pub/club/team-argentina-b",
							"time_class":"daily",
							"@id":"https://api.chess.com/pub/match/796340"
						}
					],
					"registered":[
						{
							"name":"0cean",
							"opponent":"https://api.chess.com/pub/club/team-argentina-b",
							"time_class":"daily",
							"@id":"https://api.chess.com/pub/match/796340"
						}
					]
				 }`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) (chesscompubapi.ClubMatches, error) {
			return c.GetClubMatches("chess-com-developer-community")
		},
		chesscompubapi.ClubMatches{
			Finished: []chesscompubapi.FinishedClubMatch{
				{
					Name:      "0cean",
					ID:        "796340",
					Opponent:  "team-argentina-b",
					Result:    "win",
					StartTime: chesscompubapi.UnixSecondsTimestamp(time.Unix(1675110044, 0)),
					TimeClass: "daily",
				},
			},
			Registered: []chesscompubapi.RegisteredClubMatch{
				{
					Name:      "0cean",
					ID:        "796340",
					Opponent:  "team-argentina-b",
					TimeClass: "daily",
				},
			},
			InProgress: []chesscompubapi.InProgressClubMatch{
				{
					Name:      "0cean",
					ID:        "796340",
					Opponent:  "team-argentina-b",
					StartTime: chesscompubapi.UnixSecondsTimestamp(time.Unix(1675110044, 0)),
					TimeClass: "daily",
				},
			},
		},
		t,
	)
}

func TestGetClubMatches_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/club/chess-com-developer-community/matches",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.GetClubMatches("chess-com-developer-community")
		return err
	}, t)
}
