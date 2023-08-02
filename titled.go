package chesscompubapi

import "fmt"

// ListTitledPlayers lists usernames for titled players.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-titled.
func (c *Client) ListTitledPlayers(title string) ([]string, error) {
	const urlTemplate = "titled/%s"
	players := &struct {
		Players []string `json:"players"`
	}{}
	err := c.getInto(fmt.Sprintf(urlTemplate, title), players)
	return players.Players, err
}
