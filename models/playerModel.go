package models

type PlayerModel struct {
	PlayerName  string `json:"player_name"`
	PlayerLevel int    `json:"player_level"`
}

func NewPlayerModel(playerName string, playerLevel int) PlayerModel {
	return PlayerModel{PlayerName: playerName, PlayerLevel: playerLevel}
}
