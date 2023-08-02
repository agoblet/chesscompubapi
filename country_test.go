package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestGetCountryProfile_ShouldGetProfile(t *testing.T) {
	tests := []struct {
		giveCountryCode, givePattern, giveResponseBody string
		want                                           chesscompubapi.CountryProfile
	}{
		{
			giveCountryCode: "MO",
			givePattern:     "/pub/country/MO",
			giveResponseBody: `{
				"code":"MO",
				"name":"Mordor"
			}`,
			want: chesscompubapi.CountryProfile{
				Code: "MO",
				Name: "Mordor",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveCountryCode, func(t *testing.T) {
			runOutputTestWithTestServer(
				[]testServerRoute{
					{
						pattern:      tt.givePattern,
						responseBody: tt.giveResponseBody,
						statusCode:   200,
					},
				},
				func(c *chesscompubapi.Client) (any, error) { return c.GetCountryProfile(tt.giveCountryCode) },
				tt.want,
				t,
			)
		})
	}
}

func TestGetCountryProfile_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "corruptJSON",
			giveResponseBody: `[]`,
			giveStatusCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runErrorTestWithTestServer([]testServerRoute{{
				pattern:      "/pub/country/MO",
				responseBody: tt.giveResponseBody,
				statusCode:   tt.giveStatusCode,
			}}, func(c *chesscompubapi.Client) error {
				_, err := c.GetCountryProfile("MO")
				return err
			}, t)
		})
	}
}

func TestListCountryPlayers_ShouldListPlayers(t *testing.T) {
	tests := []struct {
		giveCountryCode, givePattern, giveResponseBody string
		want                                           []string
	}{
		{
			giveCountryCode: "MO",
			givePattern:     "/pub/country/MO/players",
			giveResponseBody: `{
				"players":["hikaru","naroditsky"]
			}`,
			want: []string{"hikaru", "naroditsky"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveCountryCode, func(t *testing.T) {
			runOutputTestWithTestServer(
				[]testServerRoute{
					{
						pattern:      tt.givePattern,
						responseBody: tt.giveResponseBody,
						statusCode:   200,
					},
				},
				func(c *chesscompubapi.Client) (any, error) { return c.ListCountryPlayers(tt.giveCountryCode) },
				tt.want,
				t,
			)
		})
	}
}

func TestListCountryPlayers_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "corruptJSON",
			giveResponseBody: `[]`,
			giveStatusCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runErrorTestWithTestServer([]testServerRoute{{
				pattern:      "/pub/country/MO/players",
				responseBody: tt.giveResponseBody,
				statusCode:   tt.giveStatusCode,
			}}, func(c *chesscompubapi.Client) error {
				_, err := c.ListCountryPlayers("MO")
				return err
			}, t)
		})
	}
}

func TestListCountryClubs_ShouldListClubs(t *testing.T) {
	tests := []struct {
		giveCountryCode, givePattern, giveResponseBody string
		want                                           []chesscompubapi.StringFromPathSuffix
	}{
		{
			giveCountryCode: "MO",
			givePattern:     "/pub/country/MO/clubs",
			giveResponseBody: `{
				"clubs":["https://api.chess.com/pub/club/winners","https://api.chess.com/pub/club/losers"]
			}`,
			want: []chesscompubapi.StringFromPathSuffix{"winners", "losers"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveCountryCode, func(t *testing.T) {
			runOutputTestWithTestServer(
				[]testServerRoute{
					{
						pattern:      tt.givePattern,
						responseBody: tt.giveResponseBody,
						statusCode:   200,
					},
				},
				func(c *chesscompubapi.Client) (any, error) { return c.ListCountryClubs(tt.giveCountryCode) },
				tt.want,
				t,
			)
		})
	}
}

func TestListCountryClubs_ShouldErr(t *testing.T) {
	tests := []struct {
		name, giveResponseBody string
		giveStatusCode         int
	}{
		{
			name:             "corruptJSON",
			giveResponseBody: `[]`,
			giveStatusCode:   200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runErrorTestWithTestServer([]testServerRoute{{
				pattern:      "/pub/country/MO/clubs",
				responseBody: tt.giveResponseBody,
				statusCode:   tt.giveStatusCode,
			}}, func(c *chesscompubapi.Client) error {
				_, err := c.ListCountryClubs("MO")
				return err
			}, t)
		})
	}
}
