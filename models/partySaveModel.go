package models

//PartySaveModel is to hold party information, that gets saved to the disk when party gets saved
type PartySaveModel struct {
	PartyId      string                 `json:"party_id"`
	PartyName    string                 `json:"party_name"`
	PartyPlayers map[string]PlayerModel `json:"party_players"`
}

//NewPartySaveModel returns empty model
func NewPartySaveModel() *PartySaveModel {
	return &PartySaveModel{PartyPlayers: make(map[string]PlayerModel)}
}
