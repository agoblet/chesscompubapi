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

func TestGetPGN_ShouldGetPGN(t *testing.T) {
	tests := []struct {
		giveArchive                   chesscompubapi.Archive
		givePattern, giveResponseBody string
	}{
		{
			giveArchive: chesscompubapi.Archive{
				Username: "erik",
				Year:     2020,
				Month:    12,
			},
			givePattern: "/pub/player/erik/games/2020/12/pgn",
			giveResponseBody: `[Event "Hello!"]
[Site "Chess.com"]
[Date "2009.09.17"]
[Round "-"]
[White "jsssuperstar"]
[Black "erik"]
[Result "0-1"]
[CurrentPosition "8/8/8/p7/8/6k1/4n1b1/6K1 w - - 51 90"]
[Timezone "UTC"]
[ECO "A04"]
[ECOUrl "https://www.chess.com/openings/Reti-Opening-Pirc-Invitation"]
[UTCDate "2009.09.17"]
[UTCTime "21:40:19"]
[WhiteElo "1306"]
[BlackElo "2061"]
[TimeControl "1/259200"]
[Termination "erik won by checkmate"]
[StartTime "21:40:19"]
[EndDate "2009.10.31"]
[EndTime "17:53:17"]
[Link "https://www.chess.com/game/daily/28103371"]

1. Nf3 d6 2. Nc3 g6 3. e4 Bg7 4. Bc4 Nf6 5. d4 O-O 6. O-O a6 7. a4 c5 8. dxc5 dxc5 9. Be3 b6 10. Qd2 Qc7 11. Bh6 Rd8 12. Qc1 Bb7 13. e5 Ng4 14. Bxg7 Kxg7 15. Qf4 f5 16. h3 Nh6 17. Ng5 Bc8 18. Be6 Rd4 19. Qh2 Bb7 20. Bxf5 Kh8 21. Be4 Bxe4 22. Ne6 Qb7 23. Qf4 Bf5 24. Nxd4 cxd4 25. Qxh6 Nc6 26. Ne2 Bxc2 27. Qd2 d3 28. Nc1 Rd8 29. f4 a5 30. f5 Nxe5 31. fxg6 Nxg6 32. Rf7 Qd5 33. Rf3 Ne5 34. Rf5 Qe4 35. Qf2 Rg8 36. Rf8 Nc4 37. Rxg8+ Kxg8 38. Qg3+ Kf7 39. b3 Qe3+ 40. Qxe3 Nxe3 41. Nxd3 Bxd3 42. Rc1 Bf5 43. Rc3 Nd5 44. Rf3 Ke6 45. g4 Bg6 46. h4 Be4 47. Rg3 Ke5 48. h5 Kf4 49. Kh2 Ne3 50. Kh3 Bc2 51. g5 Bxb3 52. g6 hxg6 53. Rxg6 Bf7 54. Rxb6 Bxh5 55. Rb5 Bg4+ 56. Kh2 Nc4 57. Kg2 Be2 58. Kf2 Bd3 59. Rd5 Ke4 60. Rb5 Kd4 61. Ke1 Bc2 62. Rg5 e5 63. Ke2 Bxa4 64. Rxe5 Nxe5 65. Kd2 Bb5 66. Kc2 Nf3 67. Kb3 Kc5 68. Ka3 Kc4 69. Kb2 Kb4 70. Ka2 Bc4+ 71. Kb2 Ne5 72. Ka1 Kb3 73. Kb1 Kc3 74. Ka1 Nd3 75. Kb1 Nb4 76. Kc1 Ba2 77. Kd1 Kd3 78. Kc1 Kc3 79. Kd1 Bc4 80. Ke1 Kd3 81. Kf2 Ke4 82. Kg3 Be6 83. Kh2 Kf3 84. Kh1 Nd3 85. Kg1 Bh3 86. Kh2 Nf4 87. Kg1 Kg3 88. Kh1 Bg2+ 89. Kg1 Ne2# 0-1


[Event "1st Chess Aficionados 1601-2000 - Round 1"]
[Site "Chess.com"]
[Date "2009.10.19"]
[Round "-"]
[White "ICMike"]
[Black "erik"]
[Result "0-1"]
[Tournament "https://www.chess.com/tournament/1st-chess-aficionados-1601-2000"]
[CurrentPosition "8/6b1/8/6p1/4p3/6P1/kpK5/1N6 w - - 1 50"]
[Timezone "UTC"]
[ECO "B08"]
[ECOUrl "https://www.chess.com/openings/Pirc-Defense-Classical-Variation-4...Bg7-5.Be3-O-O"]
[UTCDate "2009.10.19"]
[UTCTime "14:52:57"]
[WhiteElo "1812"]
[BlackElo "2061"]
[TimeControl "1/86400"]
[Termination "erik won by resignation"]
[StartTime "14:52:57"]
[EndDate "2009.10.31"]
[EndTime "17:38:04"]
[Link "https://www.chess.com/game/daily/29358099"]

1. e4 d6 2. d4 Nf6 3. Nc3 g6 4. Nf3 Bg7 5. Be3 O-O 6. Bd3 c6 7. h3 b5 8. b4 a5 9. a3 Bb7 10. O-O Nbd7 11. Ne2 c5 12. dxc5 Nxe4 13. Bxe4 Bxe4 14. Rb1 dxc5 15. Ng5 Bc6 16. c3 axb4 17. cxb4 c4 18. Nd4 Qc7 19. Nxc6 Qxc6 20. Qe2 Ne5 21. Bc5 Nd3 22. h4 Nxc5 23. bxc5 Qxc5 24. Qf3 Rxa3 25. Qb7 Rb3 26. Rxb3 cxb3 27. Qe4 e6 28. Nf3 Qc4 29. Qb1 Qc2 30. Qxc2 bxc2 31. Rc1 Rc8 32. Ne1 b4 33. Rxc2 Rxc2 34. Nxc2 b3 35. Na3 Kf8 36. Nb1 Ke7 37. Kf1 Kd6 38. Nd2 b2 39. Nb1 Kc5 40. Ke2 Kc4 41. Kd2 Kb3 42. Kd3 Ka2 43. Kc2 f5 44. g3 h6 45. f3 g5 46. hxg5 hxg5 47. Nd2 e5 48. Nb1 e4 49. fxe4 fxe4 0-1`,
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

			got, err := c.GetPGN(tt.giveArchive)

			if err != nil {
				t.Errorf("expected err to be nil got %v", err)
				return
			}
			if !reflect.DeepEqual(tt.giveResponseBody, got) {
				t.Errorf("got %v, want %v", got, tt.giveResponseBody)
			}
		})
	}
}

func TestGetPGN_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
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
					pattern:      "/pub/player/johndoe/games/2022/01/pgn",
					responseBody: tt.giveResponseBody,
					statusCode:   tt.giveStatusCode,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			_, err := c.GetPGN(chesscompubapi.Archive{
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
