//go:build e2e
// +build e2e

package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestE2E(t *testing.T) {
	const username = "erik"
	c := chesscompubapi.NewClient()

	archives, err := c.ListArchives(username)
	if err != nil {
		t.Errorf("ListArchives err: %v", err)
		return
	}

	_, err = c.ListGames(archives[0])
	if err != nil {
		t.Errorf("ListGames err: %v", err)
		return
	}

	pgn, err := c.GetPGN(archives[0])
	if err != nil {
		t.Errorf("GetPGN err: %v", err)
		return
	}
	if pgn[0] != '[' {
		t.Errorf("expected PGN, got %s", pgn)
		return
	}

	playerProfile, err := c.GetPlayerProfile(username)
	if err != nil {
		t.Errorf("GetPlayerProfile err: %v", err)
		return
	}

	_, err = c.GetCountryProfile(string(playerProfile.CountryCode))
	if err != nil {
		t.Errorf("GetCountryProfile err: %v", err)
		return
	}

	_, err = c.ListCountryPlayers(string(playerProfile.CountryCode))
	if err != nil {
		t.Errorf("ListCountryPlayers err: %v", err)
		return
	}

	_, err = c.ListPlayerClubs(username)
	if err != nil {
		t.Errorf("ListPlayerClubs err: %v", err)
		return
	}
}
