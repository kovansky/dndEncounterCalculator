package models

type PartyModel struct {
	PartyPlayers      []PlayerModel
	PartyAverageLevel int
	PartyThresholds   map[string]int
	PartyMinMax       int
	PartyPerLevel     map[int]int
}
