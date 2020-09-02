package models

import (
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
)

type PartyModel struct {
	PartyPlayers      map[string]PlayerModel
	PartyAverageLevel float64
	PartyThresholds   map[string]int
	PartyMinMax       int
	PartyPerLevel     map[int]int
	PartySize         int
	PartyCategory     enum.PartyCategory
}

func NewPartyModel() *PartyModel {
	return &PartyModel{PartyPlayers: make(map[string]PlayerModel), PartyThresholds: make(map[string]int), PartyPerLevel: make(map[int]int)}
}

func (party *PartyModel) Update() PartyModel {
	party.PartyAverageLevel = party.CalculateAverageLevel()
	party.PartyThresholds = party.CalculateThresholds()
	party.PartyMinMax = party.CalculateMinMax()
	party.PartyPerLevel = party.CalculatePerLevel()
	party.PartySize = party.CountPlayers()

	return *party
}

func (party *PartyModel) AddPlayer(player PlayerModel) PartyModel {
	party.PartyPlayers[player.PlayerName] = player

	return party.Update()
}

func (party *PartyModel) RemovePlayer(player string) PartyModel {
	delete(party.PartyPlayers, player)

	return party.Update()
}

func (party *PartyModel) RemovePlayer1(player PlayerModel) PartyModel {
	return party.RemovePlayer(player.PlayerName)
}

func (party *PartyModel) CalculateAverageLevel() float64 {
	partySize := float64(len(party.PartyPlayers))
	levels := 0.0

	for _, player := range party.PartyPlayers {
		levels += float64(player.PlayerLevel)
	}

	return levels / partySize
}

func (party *PartyModel) CalculateThresholds() map[string]int {
	thresholds := map[string]int{
		"easy":   0,
		"medium": 0,
		"hard":   0,
		"deadly": 0,
	}

	for _, player := range party.PartyPlayers {
		thresholds["easy"] += constants.EasyThresholds[player.PlayerLevel]
		thresholds["medium"] += constants.MediumThresholds[player.PlayerLevel]
		thresholds["hard"] += constants.HardThresholds[player.PlayerLevel]
		thresholds["deadly"] += constants.DeadlyThresholds[player.PlayerLevel]
	}

	return thresholds
}

func (party *PartyModel) CalculateMinMax() int {
	partyMin := 21
	partyMax := 0

	for _, player := range party.PartyPlayers {
		if player.PlayerLevel < partyMin {
			partyMin = player.PlayerLevel
		}
		if player.PlayerLevel > partyMax {
			partyMax = player.PlayerLevel
		}
	}

	return partyMax - partyMin
}

func (party *PartyModel) CalculatePerLevel() map[int]int {
	perLevel := make(map[int]int)

	for i := 1; i <= 20; i++ {
		perLevel[i] = 0
	}

	for _, player := range party.PartyPlayers {
		perLevel[player.PlayerLevel] += 1
	}

	return perLevel
}

func (party *PartyModel) CountPlayers() int {
	return len(party.PartyPlayers)
}

func (party *PartyModel) GetPartyCategory() enum.PartyCategory {
	return enum.PartyCategoryBySize(party.PartySize)
}
