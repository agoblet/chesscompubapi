package chesscompubapi

type Puzzle struct {
	URL         string               `json:"url"`
	FEN         string               `json:"fen"`
	PGN         string               `json:"pgn"`
	Image       string               `json:"image"`
	Title       string               `json:"title"`
	Comments    string               `json:"comments"`
	PublishTime UnixSecondsTimestamp `json:"publish_time"`
}

// GetPuzzle gets information about the daily puzzle found in www.chess.com.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-daily-puzzle.
func (c *Client) GetPuzzle() (Puzzle, error) {
	const urlTemplate = "puzzle"
	puzzle := Puzzle{}
	err := c.getInto(urlTemplate, &puzzle)
	return puzzle, err
}

// GetRandomPuzzle gets information about a randomly picked daily puzzle.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-random-daily-puzzle.
func (c *Client) GetRandomPuzzle() (Puzzle, error) {
	const urlTemplate = "puzzle/random"
	puzzle := Puzzle{}
	err := c.getInto(urlTemplate, &puzzle)
	return puzzle, err
}
