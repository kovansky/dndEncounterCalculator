package tests

import (
	"github.com/kovansky/dndEncounterCalculator/models"
	"testing"
)

func TestPartyUpdate(t *testing.T) {
	party := models.NewPartyModel()
	player1 := models.NewPlayerModel("Alojzy", 5)
	player2 := models.NewPlayerModel("Grzegorz", 8)

	t.Log("--- Add player 1 (lvl 5) ---")
	party.AddPlayer(player1)

	t.Log("Players list:")
	t.Log(party.PartyPlayers)

	if party.PartyAverageLevel != 5.0 {
		t.Errorf("Party Average Level incorrect, got: %f; wanted: %f", party.PartyAverageLevel, 5.0)
	}

	if party.PartyThresholds["easy"] != 250 && party.PartyThresholds["medium"] != 500 && party.PartyThresholds["hard"] != 750 && party.PartyThresholds["deadly"] != 1100 {
		t.Errorf("Party Threshold incorrect.\nEasy, got: %d, wanted: %d\nMedium, got: %d, wanted: %d\nHard, got: %d, wanted: %d\nDeadly, got: %d, wanted: %d",
			party.PartyThresholds["easy"], 250,
			party.PartyThresholds["medium"], 500,
			party.PartyThresholds["hard"], 750,
			party.PartyThresholds["deadly"], 1100,
		)
	}

	if party.PartyMinMax != 0 {
		t.Errorf("Party Min-max difference incorrect, got: %d, wanted: %d", party.PartyMinMax, 0)
	}

	if party.PartyPerLevel[5] != 1 && party.PartyPerLevel[1] != 0 {
		t.Errorf("Party Per-level players incorrect.\nFirst level players, got: %d, wanted: %d\nFifth level players, got: %d, wanted: %d",
			party.PartyPerLevel[1], 0,
			party.PartyPerLevel[5], 1,
		)
	}

	t.Log("--- Add player 2 (lvl 8) ---")
	party.AddPlayer(player2)

	t.Log("Players list:")
	t.Log(party.PartyPlayers)

	if party.PartyAverageLevel != 6.5 {
		t.Errorf("Party Average Level incorrect, got: %f; wanted: %f", party.PartyAverageLevel, 6.5)
	}

	if party.PartyThresholds["easy"] != 700 && party.PartyThresholds["medium"] != 1400 && party.PartyThresholds["hard"] != 2150 && party.PartyThresholds["deadly"] != 3200 {
		t.Errorf("Party Threshold incorrect.\nEasy, got: %d, wanted: %d\nMedium, got: %d, wanted: %d\nHard, got: %d, wanted: %d\nDeadly, got: %d, wanted: %d",
			party.PartyThresholds["easy"], 700,
			party.PartyThresholds["medium"], 1400,
			party.PartyThresholds["hard"], 2150,
			party.PartyThresholds["deadly"], 3200,
		)
	}

	if party.PartyMinMax != 3 {
		t.Errorf("Party Min-max difference incorrect, got: %d, wanted: %d", party.PartyMinMax, 0)
	}

	if party.PartyPerLevel[5] != 1 && party.PartyPerLevel[8] != 1 && party.PartyPerLevel[1] != 1 {
		t.Errorf("Party Per-level players incorrect.\nFirst level players, got: %d, wanted: %d\nFifth level players, got: %d, wanted: %d\nEight level players, got: %d, wanted: %d",
			party.PartyPerLevel[1], 0,
			party.PartyPerLevel[5], 1,
			party.PartyPerLevel[8], 1,
		)
	}

	t.Log("--- Remove player 1 (lvl 5) ---")
	party.RemovePlayer1(player1)

	t.Log("Players list:")
	t.Log(party.PartyPlayers)

	if party.PartyAverageLevel != 8.0 {
		t.Errorf("Party Average Level incorrect, got: %f; wanted: %f", party.PartyAverageLevel, 5.0)
	}

	if party.PartyThresholds["easy"] != 450 && party.PartyThresholds["medium"] != 900 && party.PartyThresholds["hard"] != 1400 && party.PartyThresholds["deadly"] != 2100 {
		t.Errorf("Party Threshold incorrect.\nEasy, got: %d, wanted: %d\nMedium, got: %d, wanted: %d\nHard, got: %d, wanted: %d\nDeadly, got: %d, wanted: %d",
			party.PartyThresholds["easy"], 450,
			party.PartyThresholds["medium"], 900,
			party.PartyThresholds["hard"], 1400,
			party.PartyThresholds["deadly"], 2100,
		)
	}

	if party.PartyMinMax != 0 {
		t.Errorf("Party Min-max difference incorrect, got: %d, wanted: %d", party.PartyMinMax, 0)
	}

	if party.PartyPerLevel[8] != 1 && party.PartyPerLevel[1] != 0 {
		t.Errorf("Party Per-level players incorrect.\nFirst level players, got: %d, wanted: %d\nEight level players, got: %d, wanted: %d",
			party.PartyPerLevel[1], 0,
			party.PartyPerLevel[8], 1,
		)
	}
}
