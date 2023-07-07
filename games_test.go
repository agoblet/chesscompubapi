package chesscompubapi_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/agoblet/chesscompubapi"
)

func TestListArchives_ShouldListArchives(t *testing.T) {
	tests := []struct {
		giveUsername, givePattern, giveResponseBody string
		want                                        []chesscompubapi.Archive
	}{
		{
			giveUsername:     "johndoe",
			givePattern:      "/pub/player/johndoe/games/archives",
			giveResponseBody: `{"archives":["https://api.chess.com/pub/player/johndoe/games/2020/12"]}`,
			want: []chesscompubapi.Archive{{
				Username: "johndoe",
				Year:     2020,
				Month:    12,
			}},
		},
		{
			giveUsername:     "noarchives",
			givePattern:      "/pub/player/noarchives/games/archives",
			giveResponseBody: `{"archives":[]}`,
			want:             []chesscompubapi.Archive{},
		},
		{
			giveUsername:     "janedoe",
			givePattern:      "/pub/player/janedoe/games/archives",
			giveResponseBody: `{"archives":["https://api.chess.com/pub/player/janedoe/games/2020/09", "https://api.chess.com/pub/player/janedoe/games/2021/11"]}`,
			want: []chesscompubapi.Archive{{
				Username: "janedoe",
				Year:     2020,
				Month:    9,
			}, {
				Username: "janedoe",
				Year:     2021,
				Month:    11,
			}},
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

			got, err := c.ListArchives(tt.giveUsername)

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

func TestListArchives_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "corruptMonth",
			giveResponseBody: `{"archives":["https://api.chess.com/pub/player/johndoe/games/2020/twelve"]}`,
			giveStatusCode:   200,
		},
		{
			name:             "corruptYear",
			giveResponseBody: `{"archives":["https://api.chess.com/pub/player/johndoe/games/LEET/12"]}`,
			giveStatusCode:   200,
		},
		{
			name:             "wrongType",
			giveResponseBody: `{"archives":12345}`,
			giveStatusCode:   200,
		},
		{
			name:             "404",
			giveResponseBody: "not found",
			giveStatusCode:   404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newTestServer([]testServerRoute{
				{
					pattern:      "/pub/player/johndoe/games/archives",
					responseBody: tt.giveResponseBody,
					statusCode:   tt.giveStatusCode,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			_, err := c.ListArchives("johndoe")
			if err == nil {
				t.Error("expected err")
			}
		})
	}
}

func TestListGames_ShouldListGames(t *testing.T) {
	tests := []struct {
		giveArchive                   chesscompubapi.Archive
		givePattern, giveResponseBody string
		want                          []chesscompubapi.Game
	}{
		{
			giveArchive: chesscompubapi.Archive{
				Username: "nogames",
				Year:     2020,
				Month:    12,
			},
			givePattern:      "/pub/player/nogames/games/2020/12",
			giveResponseBody: `{"games":[]}`,
			want:             []chesscompubapi.Game{},
		},
		{
			giveArchive: chesscompubapi.Archive{
				Username: "erik",
				Year:     2009,
				Month:    10,
			},
			givePattern: "/pub/player/erik/games/2009/10",
			giveResponseBody: `{
				"games":[
				    {
						"url":"https://www.chess.com/game/daily/29099782",
						"pgn":"[Event \"Let's Play!\"]\n[Site \"Chess.com\"]\n[Date \"2009.10.01\"]\n[Round \"-\"]\n[White \"Mainline_Novelty\"]\n[Black \"erik\"]\n[Result \"1-0\"]\n[CurrentPosition \"r2q1rk1/3nppbp/2pp1np1/pp4Nb/3PP1P1/PBN1B2P/1PPQ1P2/R3K2R b KQ g3 1 12\"]\n[Timezone \"UTC\"]\n[ECO \"B07\"]\n[ECOUrl \"https://www.chess.com/openings/Pirc-Defense-Main-Line-4.Be3-Bg7-5.Qd2-c6-6.Nf3\"]\n[UTCDate \"2009.10.01\"]\n[UTCTime \"23:14:41\"]\n[WhiteElo \"1633\"]\n[BlackElo \"1920\"]\n[TimeControl \"1/259200\"]\n[Termination \"Mainline_Novelty won by resignation\"]\n[StartTime \"23:14:41\"]\n[EndDate \"2009.10.04\"]\n[EndTime \"15:38:54\"]\n[Link \"https://www.chess.com/game/daily/29099782\"]\n\n1. e4 d6 2. d4 Nf6 3. Nc3 g6 4. Be3 Bg7 5. Qd2 c6 6. Nf3 Bg4 7. Bc4 b5 8. Bb3 a5 9. a3 Nbd7 10. Ng5 O-O 11. h3 Bh5 12. g4 1-0\n",
						"time_control":"1/259200",
						"end_time":1254670734,
						"rated":true,
						"tcn":"mCZRlB!Tbs2Ucu92dlYQgv6EfAXHArWGiq5ZvM8!pxENoE",
						"uuid":"3277772e-aee0-11de-830e-00000001000b",
						"initial_setup":"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
						"fen":"r2q1rk1/3nppbp/2pp1np1/pp4Nb/3PP1P1/PBN1B2P/1PPQ1P2/R3K2R b KQ g3 1 12",
						"start_time":1254438881,
						"time_class":"daily",
						"rules":"chess",
						"white":{
							"rating":1633,
							"result":"win",
							"@id":"https://api.chess.com/pub/player/mainline_novelty",
							"username":"Mainline_Novelty",
							"uuid":"9c535e34-2dd0-11dd-8006-000000000000"
					  	},
					  	"black":{
							"rating":1920,
							"result":"resigned",
							"@id":"https://api.chess.com/pub/player/erik",
							"username":"erik",
							"uuid":"fe696c00-fcba-11db-8029-000000000000"
					  	}
					},
					{
						"url":"https://www.chess.com/game/daily/29150975",
						"pgn":"[Event \"Let's Play!\"]\n[Site \"Chess.com\"]\n[Date \"2009.10.02\"]\n[Round \"-\"]\n[White \"nym\"]\n[Black \"erik\"]\n[Result \"0-1\"]\n[CurrentPosition \"4rk2/1pp2p2/p2b1n1p/3P1bp1/2BP1p2/P1N4P/1PP3P1/R1B3K1 w - - 2 17\"]\n[Timezone \"UTC\"]\n[ECO \"C36\"]\n[ECOUrl \"https://www.chess.com/openings/Kings-Gambit-Accepted-Modern-Defense-4.exd5-Bd6\"]\n[UTCDate \"2009.10.02\"]\n[UTCTime \"17:18:53\"]\n[WhiteElo \"1898\"]\n[BlackElo \"1920\"]\n[TimeControl \"1/432000\"]\n[Termination \"erik won by resignation\"]\n[StartTime \"17:18:53\"]\n[EndDate \"2009.10.04\"]\n[EndTime \"21:15:13\"]\n[Link \"https://www.chess.com/game/daily/29150975\"]\n\n1. e4 e5 2. f4 exf4 3. Nf3 d5 4. exd5 Bd6 5. Bc4 Nf6 6. Qe2+ Qe7 7. Qxe7+ Kxe7 8. d4 Re8 9. O-O h6 10. Ne5 g5 11. Re1 Kf8 12. Nc3 Nbd7 13. Nxd7+ Bxd7 14. Rxe8+ Rxe8 15. h3 a6 16. a3 Bf5 0-1\n",
						"time_control":"1/432000",
						"end_time":1254690913,
						"rated":false,
						"accuracies":{
							"white":38.41369217259554,
							"black":95.88820275766834
						},
						"tcn":"mC0KnDKDgvZJCJ9RfA!Tdm70m080lB?8eg3VvK2Mfe09bs5ZKZ6Ze848pxWOiqZL",
						"uuid":"a87aef2c-af77-11de-83cf-00000001000b",
						"initial_setup":"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
						"fen":"4rk2/1pp2p2/p2b1n1p/3P1bp1/2BP1p2/P1N4P/1PP3P1/R1B3K1 w - - 2 17",
						"start_time":1254503933,
						"time_class":"daily",
						"rules":"chess",
						"white":{
							"rating":1898,
							"result":"resigned",
							"@id":"https://api.chess.com/pub/player/nym",
							"username":"nym",
							"uuid":"d5f6f9fa-c229-11dd-8001-000000000000"
						},
						"black":{
							"rating":1920,
							"result":"win",
							"@id":"https://api.chess.com/pub/player/erik",
							"username":"erik",
							"uuid":"fe696c00-fcba-11db-8029-000000000000"
						}
					}
				]
			}`,
			want: []chesscompubapi.Game{
				{
					White: chesscompubapi.GamePlayer{
						Rating:   1633,
						Result:   "win",
						Username: "Mainline_Novelty",
						UUID:     "9c535e34-2dd0-11dd-8006-000000000000",
					},
					Black: chesscompubapi.GamePlayer{
						Rating:   1920,
						Result:   "resigned",
						Username: "erik",
						UUID:     "fe696c00-fcba-11db-8029-000000000000",
					},
					URL:          "https://www.chess.com/game/daily/29099782",
					EndTime:      chesscompubapi.UnixSecondsTimestamp(time.Unix(1254670734, 0)),
					StartTime:    chesscompubapi.UnixSecondsTimestamp(time.Unix(1254438881, 0)),
					FEN:          "r2q1rk1/3nppbp/2pp1np1/pp4Nb/3PP1P1/PBN1B2P/1PPQ1P2/R3K2R b KQ g3 1 12",
					PGN:          "[Event \"Let's Play!\"]\n[Site \"Chess.com\"]\n[Date \"2009.10.01\"]\n[Round \"-\"]\n[White \"Mainline_Novelty\"]\n[Black \"erik\"]\n[Result \"1-0\"]\n[CurrentPosition \"r2q1rk1/3nppbp/2pp1np1/pp4Nb/3PP1P1/PBN1B2P/1PPQ1P2/R3K2R b KQ g3 1 12\"]\n[Timezone \"UTC\"]\n[ECO \"B07\"]\n[ECOUrl \"https://www.chess.com/openings/Pirc-Defense-Main-Line-4.Be3-Bg7-5.Qd2-c6-6.Nf3\"]\n[UTCDate \"2009.10.01\"]\n[UTCTime \"23:14:41\"]\n[WhiteElo \"1633\"]\n[BlackElo \"1920\"]\n[TimeControl \"1/259200\"]\n[Termination \"Mainline_Novelty won by resignation\"]\n[StartTime \"23:14:41\"]\n[EndDate \"2009.10.04\"]\n[EndTime \"15:38:54\"]\n[Link \"https://www.chess.com/game/daily/29099782\"]\n\n1. e4 d6 2. d4 Nf6 3. Nc3 g6 4. Be3 Bg7 5. Qd2 c6 6. Nf3 Bg4 7. Bc4 b5 8. Bb3 a5 9. a3 Nbd7 10. Ng5 O-O 11. h3 Bh5 12. g4 1-0\n",
					TimeControl:  "1/259200",
					InitialSetup: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
					UUID:         "3277772e-aee0-11de-830e-00000001000b",
					TimeClass:    "daily",
					Rules:        "chess",
					Rated:        true,
					TCN:          "mCZRlB!Tbs2Ucu92dlYQgv6EfAXHArWGiq5ZvM8!pxENoE",
				},
				{
					White: chesscompubapi.GamePlayer{
						Rating:   1898,
						Result:   "resigned",
						Username: "nym",
						UUID:     "d5f6f9fa-c229-11dd-8001-000000000000",
					},
					Black: chesscompubapi.GamePlayer{
						Rating:   1920,
						Result:   "win",
						Username: "erik",
						UUID:     "fe696c00-fcba-11db-8029-000000000000",
					},
					URL:          "https://www.chess.com/game/daily/29150975",
					EndTime:      chesscompubapi.UnixSecondsTimestamp(time.Unix(1254690913, 0)),
					StartTime:    chesscompubapi.UnixSecondsTimestamp(time.Unix(1254503933, 0)),
					FEN:          "4rk2/1pp2p2/p2b1n1p/3P1bp1/2BP1p2/P1N4P/1PP3P1/R1B3K1 w - - 2 17",
					UUID:         "a87aef2c-af77-11de-83cf-00000001000b",
					TimeClass:    "daily",
					Rules:        "chess",
					PGN:          "[Event \"Let's Play!\"]\n[Site \"Chess.com\"]\n[Date \"2009.10.02\"]\n[Round \"-\"]\n[White \"nym\"]\n[Black \"erik\"]\n[Result \"0-1\"]\n[CurrentPosition \"4rk2/1pp2p2/p2b1n1p/3P1bp1/2BP1p2/P1N4P/1PP3P1/R1B3K1 w - - 2 17\"]\n[Timezone \"UTC\"]\n[ECO \"C36\"]\n[ECOUrl \"https://www.chess.com/openings/Kings-Gambit-Accepted-Modern-Defense-4.exd5-Bd6\"]\n[UTCDate \"2009.10.02\"]\n[UTCTime \"17:18:53\"]\n[WhiteElo \"1898\"]\n[BlackElo \"1920\"]\n[TimeControl \"1/432000\"]\n[Termination \"erik won by resignation\"]\n[StartTime \"17:18:53\"]\n[EndDate \"2009.10.04\"]\n[EndTime \"21:15:13\"]\n[Link \"https://www.chess.com/game/daily/29150975\"]\n\n1. e4 e5 2. f4 exf4 3. Nf3 d5 4. exd5 Bd6 5. Bc4 Nf6 6. Qe2+ Qe7 7. Qxe7+ Kxe7 8. d4 Re8 9. O-O h6 10. Ne5 g5 11. Re1 Kf8 12. Nc3 Nbd7 13. Nxd7+ Bxd7 14. Rxe8+ Rxe8 15. h3 a6 16. a3 Bf5 0-1\n",
					TimeControl:  "1/432000",
					Rated:        false,
					InitialSetup: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
					TCN:          "mC0KnDKDgvZJCJ9RfA!Tdm70m080lB?8eg3VvK2Mfe09bs5ZKZ6Ze848pxWOiqZL",
					Accuracies: &chesscompubapi.Accuracies{
						White: 38.41369217259554,
						Black: 95.88820275766834,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveArchive.Username, func(t *testing.T) {
			server := newTestServer([]testServerRoute{
				{
					pattern:      tt.givePattern,
					responseBody: tt.giveResponseBody,
					statusCode:   200,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			got, err := c.ListGames(tt.giveArchive)

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

func TestListGames_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "wrongTimestampType",
			giveResponseBody: `{"games":[{"start_time":"not a number"}]}`,
			giveStatusCode:   200,
		},
		{
			name:             "wrongTimestampType",
			giveResponseBody: `{"games":[{"black":"not an object"}]}`,
			giveStatusCode:   200,
		},
		{
			name:             "500",
			giveResponseBody: "internal server error",
			giveStatusCode:   500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newTestServer([]testServerRoute{
				{
					pattern:      "/pub/player/johndoe/games/2022/01",
					responseBody: tt.giveResponseBody,
					statusCode:   tt.giveStatusCode,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			_, err := c.ListGames(chesscompubapi.Archive{
				Username: "johndoe",
				Year:     2022,
				Month:    1,
			})
			if err == nil {
				t.Error("expected err")
			}
		})
	}
}
