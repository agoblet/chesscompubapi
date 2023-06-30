package chesscompubapi_test

import (
	"reflect"
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
				LastOnline:  time.Unix(1687899135, 0),
				Joined:      time.Unix(1389043258, 0),
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
				"country":""
			}`,
			want: chesscompubapi.PlayerProfile{
				Joined:     time.Unix(0, 0),
				LastOnline: time.Unix(0, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveUsername, func(t *testing.T) {
			server := newTestServer([]testServerRoute{
				{
					pattern:      tt.givePattern,
					responseBody: tt.giveResponseBody,
					statusCode:   200,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			got, err := c.GetPlayerProfile(tt.giveUsername)

			if err != nil {
				t.Errorf("expected err to be nil got %v", err)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
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
			server := newTestServer([]testServerRoute{
				{
					pattern:      "/pub/player/johndoe",
					responseBody: tt.giveResponseBody,
					statusCode:   tt.giveStatusCode,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			_, err := c.GetPlayerProfile("johndoe")
			if err == nil {
				t.Error("expected err")
			}
		})
	}
}
