package chesscompubapi_test

import (
	"testing"
	"time"

	"github.com/agoblet/chesscompubapi"
)

func TestGetPuzzle_ShouldGetPuzzle(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/puzzle",
				responseBody: `{
					"title":"Zwischenzug Zigzag",
					"url":"https://www.chess.com/forum/view/daily-puzzles/9-5-2023-zwischenzug-zigzag",
					"publish_time":1693897200,
					"fen":"r1b3k1/1p2rqb1/p2pnpp1/2n4p/P2NPB1P/2Q2NPB/1PP2P2/R3R1K1 b - - 0 1",
					"pgn":"[Result \"*\"]\r\n[FEN \"r1b3k1/1p2rqb1/p2pnpp1/2n4p/P2NPB1P/2Q2NPB/1PP2P2/R3R1K1 b - - 0 1\"]\r\n\r\n1... Nxd4 2. Bxc8 Nxf3+ 3. Qxf3 Rxc8 *",
					"image":"https://www.chess.com/dynboard?fen=r1b3k1/1p2rqb1/p2pnpp1/2n4p/P2NPB1P/2Q2NPB/1PP2P2/R3R1K1%20b%20-%20-%200%201&size=2"
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) (chesscompubapi.Puzzle, error) { return c.GetPuzzle() },
		chesscompubapi.Puzzle{
			Title:       "Zwischenzug Zigzag",
			URL:         "https://www.chess.com/forum/view/daily-puzzles/9-5-2023-zwischenzug-zigzag",
			PublishTime: chesscompubapi.UnixSecondsTimestamp(time.Unix(1693897200, 0)),
			FEN:         "r1b3k1/1p2rqb1/p2pnpp1/2n4p/P2NPB1P/2Q2NPB/1PP2P2/R3R1K1 b - - 0 1",
			PGN:         "[Result \"*\"]\r\n[FEN \"r1b3k1/1p2rqb1/p2pnpp1/2n4p/P2NPB1P/2Q2NPB/1PP2P2/R3R1K1 b - - 0 1\"]\r\n\r\n1... Nxd4 2. Bxc8 Nxf3+ 3. Qxf3 Rxc8 *",
			Image:       "https://www.chess.com/dynboard?fen=r1b3k1/1p2rqb1/p2pnpp1/2n4p/P2NPB1P/2Q2NPB/1PP2P2/R3R1K1%20b%20-%20-%200%201&size=2",
		},
		t,
	)
}

func TestGetPuzzle_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/puzzle",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.GetPuzzle()
		return err
	}, t)
}

func TestGetRandomPuzzle_ShouldGetPuzzle(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/puzzle/random",
				responseBody: `{
					"title":"International Games Week",
					"url":"https://www.chess.com/forum/view/daily-puzzles/11-7-2021-international-games-week",
					"publish_time":1636268400,
					"fen":"1r3rk1/4b2p/pq1pb1p1/6P1/2p1PB2/P7/1PPQ4/2KN2RR b - - 1 1",
					"pgn":"[Result \"*\"]\r\n[FEN \"1r3rk1/4b2p/pq1pb1p1/6P1/2p1PB2/P7/1PPQ4/2KN2RR b - - 1 1\"]\r\n\r\n1...c3 2.Qxc3 Rxf4 3.Rxh7 Bf6 4.gxf6 Kxh7 5.Rh1+ Kg8 *",
					"image":"https://www.chess.com/dynboard?fen=1r3rk1/4b2p/pq1pb1p1/6P1/2p1PB2/P7/1PPQ4/2KN2RR%20b%20-%20-%201%201&size=2"
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) (chesscompubapi.Puzzle, error) { return c.GetRandomPuzzle() },
		chesscompubapi.Puzzle{
			Title:       "International Games Week",
			URL:         "https://www.chess.com/forum/view/daily-puzzles/11-7-2021-international-games-week",
			PublishTime: chesscompubapi.UnixSecondsTimestamp(time.Unix(1636268400, 0)),
			FEN:         "1r3rk1/4b2p/pq1pb1p1/6P1/2p1PB2/P7/1PPQ4/2KN2RR b - - 1 1",
			PGN:         "[Result \"*\"]\r\n[FEN \"1r3rk1/4b2p/pq1pb1p1/6P1/2p1PB2/P7/1PPQ4/2KN2RR b - - 1 1\"]\r\n\r\n1...c3 2.Qxc3 Rxf4 3.Rxh7 Bf6 4.gxf6 Kxh7 5.Rh1+ Kg8 *",
			Image:       "https://www.chess.com/dynboard?fen=1r3rk1/4b2p/pq1pb1p1/6P1/2p1PB2/P7/1PPQ4/2KN2RR%20b%20-%20-%201%201&size=2",
		},
		t,
	)
}

func TestGetRandomPuzzle_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/puzzle/random",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.GetRandomPuzzle()
		return err
	}, t)
}
