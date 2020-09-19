package models

type PartySaveModel struct {
	PartyId      string                 `json:"party_id"`
	PartyName    string                 `json:"party_name"`
	PartyPlayers map[string]PlayerModel `json:"party_players"`
}

func NewPartySaveModel() *PartySaveModel {
	return &PartySaveModel{PartyPlayers: make(map[string]PlayerModel)}
}
