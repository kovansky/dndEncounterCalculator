package models

import (
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
)

//PartyModel is to hold the party information - from players list to players per level
type PartyModel struct {
	PartyPlayers      map[string]PlayerModel `json:"party_players"`
	PartyAverageLevel float64                `json:"party_average_level"`
	PartyThresholds   map[string]int         `json:"party_thresholds"`
	PartyMinMax       int                    `json:"party_min_max"` // hold difference between the player with highest level and with lowest level
	PartyPerLevel     map[int]int            `json:"party_per_level"`
	PartySize         int                    `json:"party_size"`
	PartyCategory     enum.PartyCategory     `json:"party_category"`
}

//NewPartyModel inits empty model
func NewPartyModel() *PartyModel {
	return &PartyModel{PartyPlayers: make(map[string]PlayerModel), PartyThresholds: make(map[string]int), PartyPerLevel: make(map[int]int)}
}

//Update calculates and sets all values, that depends on players map. It should be run every time the players map is changed:
// - when player is added
// - when player is edited
// - when player is removed
func (party *PartyModel) Update() PartyModel {
	party.PartyAverageLevel = party.CalculateAverageLevel()
	party.PartyThresholds = party.CalculateThresholds()
	party.PartyMinMax = party.CalculateMinMax()
	party.PartyPerLevel = party.CalculatePerLevel()
	party.PartySize = party.CountPlayers()
	party.PartyCategory = party.GetPartyCategory()

	return *party
}

//AddPlayer adds player to players map and runs Update
func (party *PartyModel) AddPlayer(player PlayerModel) PartyModel {
	party.PartyPlayers[player.PlayerName] = player

	return party.Update()
}

//RemovePlayer removes player by player name and runs Update
func (party *PartyModel) RemovePlayer(player string) PartyModel {
	delete(party.PartyPlayers, player)

	return party.Update()
}

//RemovePlayer1 removes player by passed player model and runs Update
func (party *PartyModel) RemovePlayer1(player PlayerModel) PartyModel {
	return party.RemovePlayer(player.PlayerName)
}

//CalculateAverageLevel calculates average level of all party players
func (party *PartyModel) CalculateAverageLevel() float64 {
	var (
		// Get party size
		partySize = float64(len(party.PartyPlayers))
		// Initialize variable to hold all levels summed up
		levels = 0.0
	)

	// Iterate through players map and add each level to levels
	for _, player := range party.PartyPlayers {
		levels += float64(player.PlayerLevel)
	}

	// Count and return the average
	return levels / partySize
}

//CalculateThresholds calculates all thresholds (for easy, medium, hard and deadly encounter) of party, based on level - threshold relation (see thresholdConstants.go)
func (party *PartyModel) CalculateThresholds() map[string]int {
	// Initialize variable
	thresholds := map[string]int{
		"easy":   0,
		"medium": 0,
		"hard":   0,
		"deadly": 0,
	}

	// Iterate through players in map and add their thresholds by level to party thresholds
	for _, player := range party.PartyPlayers {
		thresholds["easy"] += constants.EasyThresholds[player.PlayerLevel]
		thresholds["medium"] += constants.MediumThresholds[player.PlayerLevel]
		thresholds["hard"] += constants.HardThresholds[player.PlayerLevel]
		thresholds["deadly"] += constants.DeadlyThresholds[player.PlayerLevel]
	}

	return thresholds
}

//CalculateMinMax counts level difference between player with highest level and with the lowest one
func (party *PartyModel) CalculateMinMax() int {
	// Init variables to hold highest/lowest level values
	var (
		partyMin = 21
		partyMax = 0
	)

	for _, player := range party.PartyPlayers {
		// If player level is lower, than current minimum, overwrite
		if player.PlayerLevel < partyMin {
			partyMin = player.PlayerLevel
		}
		// If player level is higher, than current maximum, overwrite
		if player.PlayerLevel > partyMax {
			partyMax = player.PlayerLevel
		}
	}

	// Count and return the difference
	return partyMax - partyMin
}

//CalculatePerLevel calculate amount of players on each level
func (party *PartyModel) CalculatePerLevel() map[int]int {
	// Create empty map (level => amount of players in that level)
	perLevel := make(map[int]int)

	// Set up default value for each level (0)
	for i := 1; i <= 20; i++ {
		perLevel[i] = 0
	}

	// Iterate through players map and add players to count by level
	for _, player := range party.PartyPlayers {
		perLevel[player.PlayerLevel] += 1
	}

	return perLevel
}

//CountPlayers returns amount of players (length of players map)
func (party *PartyModel) CountPlayers() int {
	return len(party.PartyPlayers)
}

//GetPartyCategory returns party category (small, standard, big) by party size
func (party *PartyModel) GetPartyCategory() enum.PartyCategory {
	return enum.PartyCategoryBySize(party.PartySize)
}
