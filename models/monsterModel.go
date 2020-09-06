package models

import "github.com/kovansky/dndEncounterCalculator/constants"

type MonsterModel struct {
	MonsterName    string  `json:"monster_name"`
	MonsterCR      float32 `json:"monster_cr"`
	MonsterXP      int     `json:"monster_xp"`
	MonstersAmount int     `json:"monsters_amount"`
	GroupXP        int     `json:"group_xp"`
	CountInCRMod   bool    `json:"count_in_cr_mod"`
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
