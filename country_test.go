package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestGetCountryProfile_ShouldGetProfile(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/country/MO",
				responseBody: `{
					"code":"MO",
					"name":"Mordor"
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) (chesscompubapi.CountryProfile, error) {
			return c.GetCountryProfile("MO")
		},
		chesscompubapi.CountryProfile{
			Code: "MO",
			Name: "Mordor",
		},
		t,
	)
}

func TestGetCountryProfile_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/country/MO",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.GetCountryProfile("MO")
		return err
	}, t)
}

func TestListCountryPlayers_ShouldListPlayers(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/country/MO/players",
				responseBody: `{
					"players":["hikaru","naroditsky"]
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) ([]string, error) { return c.ListCountryPlayers("MO") },
		[]string{"hikaru", "naroditsky"},
		t,
	)
}

func TestListCountryPlayers_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/country/MO/players",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.ListCountryPlayers("MO")
		return err
	}, t)
}

func TestListCountryClubs_ShouldListClubs(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/country/MO/clubs",
				responseBody: `{
					"clubs":["https://api.chess.com/pub/club/winners","https://api.chess.com/pub/club/losers"]
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) ([]chesscompubapi.StringFromPathSuffix, error) {
			return c.ListCountryClubs("MO")
		},
		[]chesscompubapi.StringFromPathSuffix{"winners", "losers"},
		t,
	)
}

func TestListCountryClubs_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/country/MO/clubs",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.ListCountryClubs("MO")
		return err
	}, t)
}
