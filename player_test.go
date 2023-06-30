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
					LastActivity: time.Unix(1626902692, 0),
					Joined:       time.Unix(1468775412, 0),
					Icon:         "https://images.chesscomfiles.com/uploads/v1/group/3541.45d6eb2c.50x50o.af307b09ebe7.png",
					URL:          "https://www.chess.com/club/open-discussion",
				},
				{
					ID:           "chesscom---tactics-trainer-approvers",
					Name:         "Chess.com - Tactics Trainer Approvers",
					LastActivity: time.Unix(1626902680, 0),
					Joined:       time.Unix(1210299490, 0),
					Icon:         "https://images.chesscomfiles.com/uploads/v1/group/4146.cd1f3309.50x50o.2daf5479afb9.gif",
					URL:          "https://www.chess.com/club/chesscom---tactics-trainer-approvers",
				},
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

			got, err := c.ListPlayerClubs(tt.giveUsername)

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newTestServer([]testServerRoute{
				{
					pattern:      "/pub/player/johndoe/clubs",
					responseBody: tt.giveResponseBody,
					statusCode:   tt.giveStatusCode,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			_, err := c.ListPlayerClubs("johndoe")
			if err == nil {
				t.Error("expected err")
			}
		})
	}
}
