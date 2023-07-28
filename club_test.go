package chesscompubapi_test

import (
	"testing"
	"time"

	"github.com/agoblet/chesscompubapi"
)

func TestGetClub_ShouldGetClub(t *testing.T) {
	tests := []struct {
		giveResponseBody, name string
		want                   chesscompubapi.Club
	}{
		{
			name: "basic",
			giveResponseBody: `{
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
			want: chesscompubapi.Club{
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runOutputTestWithTestServer(
				[]testServerRoute{
					{
						pattern:      "/pub/club/chess-com-developer-community",
						responseBody: tt.giveResponseBody,
						statusCode:   200,
					},
				},
				func(c *chesscompubapi.Client) (any, error) { return c.GetClub("chess-com-developer-community") },
				tt.want,
				t,
			)
		})
	}
}

func TestGetClub_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "corruptJSON",
			giveResponseBody: `[]`,
			giveStatusCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runErrorTestWithTestServer([]testServerRoute{{
				pattern:      "/pub/club/chess-com-developer-community",
				responseBody: tt.giveResponseBody,
				statusCode:   tt.giveStatusCode,
			}}, func(c *chesscompubapi.Client) error {
				_, err := c.ListStreamers()
				return err
			}, t)
		})
	}
}
