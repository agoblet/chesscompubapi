package chesscompubapi

import (
	"fmt"
)

type CountryProfile struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetCountryProfile gets the profile of the country.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-country-profile.
func (c *Client) GetCountryProfile(code string) (CountryProfile, error) {
	const urlTemplate = "country/%s"
	profile := &CountryProfile{}
	err := c.getInto(fmt.Sprintf(urlTemplate, code), profile)
	return *profile, err
}

// ListCountryPlayers lists usernames for players who identify themselves as being in this country.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-country-players.
func (c *Client) ListCountryPlayers(code string) ([]string, error) {
	const urlTemplate = "country/%s/players"
	players := &struct {
		Players []string `json:"players"`
	}{}
	err := c.getInto(fmt.Sprintf(urlTemplate, code), players)
	return players.Players, err
}
