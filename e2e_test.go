//go:build e2e
// +build e2e

package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestE2E(t *testing.T) {
	c := chesscompubapi.NewClient()

	archives, err := c.ListArchives("erik")
	if err != nil {
		t.Errorf("ListArchives err: %v", err)
		return
	}

	_, err = c.ListGames(archives[0])
	if err != nil {
		t.Errorf("ListGames err: %v", err)
		return
	}

	playerProfile, err := c.GetPlayerProfile("erik")
	if err != nil {
		t.Errorf("GetPlayerProfile err: %v", err)
		return
	}

	_, err = c.GetCountryProfile(playerProfile.CountryCode)
	if err != nil {
		t.Errorf("GetCountryProfile err: %v", err)
		return
	}
}
