package chesscompubapi_test

import (
	"reflect"
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
			server := newTestServer([]testServerRoute{
				{
					pattern:      tt.givePattern,
					responseBody: tt.giveResponseBody,
					statusCode:   200,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			got, err := c.GetCountryProfile(tt.giveCountryCode)

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
			server := newTestServer([]testServerRoute{
				{
					pattern:      "/pub/country/MO",
					responseBody: tt.giveResponseBody,
					statusCode:   tt.giveStatusCode,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			_, err := c.GetCountryProfile("MO")
			if err == nil {
				t.Error("expected err")
			}
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
			server := newTestServer([]testServerRoute{
				{
					pattern:      tt.givePattern,
					responseBody: tt.giveResponseBody,
					statusCode:   200,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			got, err := c.ListCountryPlayers(tt.giveCountryCode)

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
			server := newTestServer([]testServerRoute{
				{
					pattern:      "/pub/country/MO",
					responseBody: tt.giveResponseBody,
					statusCode:   tt.giveStatusCode,
				},
			})
			defer server.Close()
			c := chesscompubapi.NewClient(chesscompubapi.WithBaseURL(server.URL))

			_, err := c.GetCountryProfile("MO")
			if err == nil {
				t.Error("expected err")
			}
		})
	}
}
