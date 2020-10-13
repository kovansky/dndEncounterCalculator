/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package models

//PlayerModel is to hold data of single player
type PlayerModel struct {
	PlayerName  string `json:"player_name"`
	PlayerLevel int    `json:"player_level"`
}

//NewPlayerModel creates a new PlayerModel with playerName and playerLevel
func NewPlayerModel(playerName string, playerLevel int) PlayerModel {
	return PlayerModel{PlayerName: playerName, PlayerLevel: playerLevel}
}
