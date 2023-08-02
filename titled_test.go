package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestListTitledPlayers_ShouldListPlayers(t *testing.T) {
	tests := []struct {
		giveTitle, givePattern, giveResponseBody string
		want                                     []string
	}{
		{
			giveTitle:   "GM",
			givePattern: "/pub/titled/GM",
			giveResponseBody: `{
				"players":["hikaru","naroditsky"]
			}`,
			want: []string{"hikaru", "naroditsky"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.giveTitle, func(t *testing.T) {
			runOutputTestWithTestServer(
				[]testServerRoute{
					{
						pattern:      tt.givePattern,
						responseBody: tt.giveResponseBody,
						statusCode:   200,
					},
				},
				func(c *chesscompubapi.Client) (any, error) { return c.ListTitledPlayers(tt.giveTitle) },
				tt.want,
				t,
			)
		})
	}
}

func TestListTitledPlayers_ShouldErr(t *testing.T) {
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
				pattern:      "/pub/titled/GM",
				responseBody: tt.giveResponseBody,
				statusCode:   tt.giveStatusCode,
			}}, func(c *chesscompubapi.Client) error {
				_, err := c.ListTitledPlayers("GM")
				return err
			}, t)
		})
	}
}
