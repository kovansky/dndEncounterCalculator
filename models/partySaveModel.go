/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

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
