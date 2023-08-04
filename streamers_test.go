package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestListStreamers_ShouldListStreamers(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/streamers",
				responseBody: `{
					"streamers":[
						{
							"username":"ChessKid",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/17491514.a09cc34c.50x50o.f22166f6f601.png",
							"twitch_url":"https://twitch.tv/chesskid",
							"url":"https://www.chess.com/member/ChessKid",
							"is_live":true,
							"is_community_streamer":false
						},
						{
							"username":"AlessiaSanteramo",
							"avatar":"https://images.chesscomfiles.com/uploads/v1/user/60335128.c415baed.50x50o.41f429e8291c.jpg",
							"twitch_url":"https://twitch.tv/alessiasanteramo",
							"url":"https://www.chess.com/member/AlessiaSanteramo",
							"is_live":true,
							"is_community_streamer":false
						}
					]
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) ([]chesscompubapi.Streamer, error) { return c.ListStreamers() },
		[]chesscompubapi.Streamer{
			{
				Username:            "ChessKid",
				Avatar:              "https://images.chesscomfiles.com/uploads/v1/user/17491514.a09cc34c.50x50o.f22166f6f601.png",
				TwitchURL:           "https://twitch.tv/chesskid",
				URL:                 "https://www.chess.com/member/ChessKid",
				IsLive:              true,
				IsCommunityStreamer: false,
			},
			{
				Username:            "AlessiaSanteramo",
				Avatar:              "https://images.chesscomfiles.com/uploads/v1/user/60335128.c415baed.50x50o.41f429e8291c.jpg",
				TwitchURL:           "https://twitch.tv/alessiasanteramo",
				URL:                 "https://www.chess.com/member/AlessiaSanteramo",
				IsLive:              true,
				IsCommunityStreamer: false,
			},
		},
		t,
	)
}

func TestListStreamers_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/streamers",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.ListStreamers()
		return err
	}, t)
}
