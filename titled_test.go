package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestListTitledPlayers_ShouldListPlayers(t *testing.T) {
	runOutputTestWithTestServer(
		[]testServerRoute{
			{
				pattern: "/pub/titled/GM",
				responseBody: `{
					"players":["hikaru","naroditsky"]
				}`,
				statusCode: 200,
			},
		},
		func(c *chesscompubapi.Client) ([]string, error) { return c.ListTitledPlayers("GM") },
		[]string{"hikaru", "naroditsky"},
		t,
	)
}

func TestListTitledPlayers_ShouldErr(t *testing.T) {
	runErrorTestWithTestServer([]testServerRoute{{
		pattern:      "/pub/titled/GM",
		responseBody: "[]",
		statusCode:   200,
	}}, func(c *chesscompubapi.Client) error {
		_, err := c.ListTitledPlayers("GM")
		return err
	}, t)
}
