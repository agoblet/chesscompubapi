package chesscompubapi

import (
	"encoding/json"
	"fmt"
)

type CountryProfile struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetCountryProfile lists the profile of the country
func (c *Client) GetCountryProfile(code string) (CountryProfile, error) {
	profile := &CountryProfile{}

	const urlTemplate = "country/%s"
	body, err := c.get(fmt.Sprintf(urlTemplate, code))
	if err != nil {
		return *profile, err
	}

	if err := json.Unmarshal(body, profile); err != nil {
		return *profile, err
	}

	return *profile, nil
}
