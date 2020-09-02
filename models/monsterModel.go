package models

import "github.com/kovansky/dndEncounterCalculator/constants"

type MonsterModel struct {
	MonsterName    string
	MonsterCR      float32
	MonsterXP      int
	MonstersAmount int
	GroupXP        int
	CountInCRMod   bool
}

func NewMonsterModel() *MonsterModel {
	return &MonsterModel{}
}

func (monster *MonsterModel) Update() MonsterModel {
	monster.MonsterXP = monster.CalculateMonsterXP()
	monster.GroupXP = monster.CalculateGroupXP()

	return *monster
}

func (monster *MonsterModel) CalculateMonsterXP() int {
	if monster.MonsterCR != 0 && constants.CRXP[monster.MonsterCR] != 0 {
		return constants.CRXP[monster.MonsterCR]
	} else {
		return -1
	}
}

func (monster *MonsterModel) CalculateGroupXP() int {
	return monster.MonsterXP * monster.MonstersAmount
}
