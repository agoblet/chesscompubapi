//go:build e2e
// +build e2e

package chesscompubapi_test

import (
	"testing"

	"github.com/agoblet/chesscompubapi"
)

func TestE2E(t *testing.T) {
	c := chesscompubapi.NewClient()

	leaderboards, err := c.GetLeaderboards()
	if err != nil {
		t.Errorf("GetLeaderboards err: %v", err)
		return
	}

	archives, err := c.ListArchives(leaderboards.LiveBlitz[0].Username)
	if err != nil {
		t.Errorf("ListArchives err: %v", err)
		return
	}

	games, err := c.ListGames(archives[0])
	if err != nil {
		t.Errorf("ListGames err: %v", err)
		return
	}
	if len(games) == 0 {
		t.Errorf("ListGames expected output")
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

	_, err = c.ListGamesToMove(archives[0].Username)
	if err != nil {
		t.Errorf("ListGamesToMove err: %v", err)
		return
	}

	playerProfile, err := c.GetPlayerProfile(leaderboards.LiveBullet[2].Username)
	if err != nil {
		t.Errorf("GetPlayerProfile err: %v", err)
		return
	}

	titleTwins, err := c.ListTitledPlayers(*playerProfile.Title)
	if err != nil {
		t.Errorf("ListTitledPlayers err: %v", err)
		return
	}
	if len(titleTwins) == 0 {
		t.Errorf("ListTitledPlayers expected output")
		return
	}

	country, err := c.GetCountryProfile(string(playerProfile.CountryCode))
	if err != nil {
		t.Errorf("GetCountryProfile err: %v", err)
		return
	}

	countryPlayers, err := c.ListCountryPlayers(country.Code)
	if err != nil {
		t.Errorf("ListCountryPlayers err: %v", err)
		return
	}
	if len(countryPlayers) == 0 {
		t.Errorf("ListCountryPlayers expected output")
		return
	}

	playerClubs, err := c.ListPlayerClubs(games[0].White.Username)
	if err != nil {
		t.Errorf("ListPlayerClubs err: %v", err)
		return
	}

	club, err := c.GetClub(string(playerClubs[0].ID))
	if err != nil {
		t.Errorf("GetClub err: %v", err)
		return
	}

	activity, err := c.GetClubMemberActivity(string(club.ID))
	if err != nil {
		t.Errorf("GetClubMemberActivity err: %v", err)
		return
	}
	if len(activity.AllTime) == 0 {
		t.Errorf("GetClubMemberActivity expected output")
		return
	}

	streamers, err := c.ListStreamers()
	if err != nil {
		t.Errorf("ListStreamers err: %v", err)
		return
	}
	if len(streamers) == 0 {
		t.Errorf("ListStreamers expected output")
		return
	}

	_, err = c.GetRandomPuzzle()
	if err != nil {
		t.Errorf("GetRandomPuzzle err: %v", err)
		return
	}

	_, err = c.GetDailyPuzzle()
	if err != nil {
		t.Errorf("GetDailyPuzzle err: %v", err)
		return
	}
}
