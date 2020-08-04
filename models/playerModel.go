package models

type PlayerModel struct {
	PlayerName  string
	PlayerLevel int
}

func NewPlayerModel(playerName string, playerLevel int) *PlayerModel {
	return &PlayerModel{PlayerName: playerName, PlayerLevel: playerLevel}
}
