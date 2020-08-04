package models

import "github.com/kovansky/dndEncounterCalculator/constants"

type PartyModel struct {
	PartyPlayers      []PlayerModel
	PartyAverageLevel float32
	PartyThresholds   map[string]int
	PartyMinMax       int
	PartyPerLevel     map[int]int
}

func (party PartyModel) Update() {
	partySize := len(party.PartyPlayers)
	average := 0
	thresholds := map[string]int{
		"easy":   0,
		"medium": 0,
		"hard":   0,
		"deadly": 0,
	}
	partyMin := 21
	partyMax := 0
	perLevel := make(map[int]int)

	for i := 1; i <= 20; i++ {
		perLevel[i] = 0
	}

	for _, player := range party.PartyPlayers {
		average += player.PlayerLevel

		thresholds["easy"] += constants.EasyThresholds[player.PlayerLevel]
		thresholds["medium"] += constants.MediumThresholds[player.PlayerLevel]
		thresholds["hard"] += constants.HardThresholds[player.PlayerLevel]
		thresholds["deadly"] += constants.DeadlyThresholds[player.PlayerLevel]

		if player.PlayerLevel < partyMin {
			partyMin = player.PlayerLevel
		}
		if player.PlayerLevel > partyMax {
			partyMax = player.PlayerLevel
		}

		perLevel[player.PlayerLevel] += 1
	}

	party.PartyAverageLevel = float32(average / partySize)
	party.PartyThresholds = thresholds
	party.PartyMinMax = partyMax - partyMin
	party.PartyPerLevel = perLevel
}
