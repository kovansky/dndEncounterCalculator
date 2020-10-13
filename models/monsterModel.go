/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package models

import "github.com/kovansky/dndEncounterCalculator/constants"

//MonsterModel represents a monster, or monsters of same type (but may be more than one)
type MonsterModel struct {
	MonsterName    string  `json:"monster_name"`
	MonsterCR      float32 `json:"monster_cr"`
	MonsterXP      int     `json:"monster_xp"`
	MonstersAmount int     `json:"monsters_amount"`
	GroupXP        int     `json:"group_xp"`
	CountInCRMod   bool    `json:"count_in_cr_mod"` // declares, if this monster(s) should be counted as difficult in terms of encounter difficulty
}

//NewMonsterModel returns empty model
func NewMonsterModel() *MonsterModel {
	return &MonsterModel{}
}

//Update calculates and sets all values, that depends on monster CR. It should be run every time the model is changed:
func (monster *MonsterModel) Update() MonsterModel {
	monster.MonsterXP = monster.CalculateMonsterXP()
	monster.GroupXP = monster.CalculateGroupXP()

	return *monster
}

//CalculateMonsterXP calculate single monster XP value from CR - XP relation
func (monster *MonsterModel) CalculateMonsterXP() int {
	if monster.MonsterCR != 0 && constants.CRXP[monster.MonsterCR] != 0 {
		return constants.CRXP[monster.MonsterCR]
	} else {
		return -1
	}
}

//CalculateGroupXP calculates XP value of whole group (single monster XP multiplied by monsters amount)
func (monster *MonsterModel) CalculateGroupXP() int {
	return monster.MonsterXP * monster.MonstersAmount
}
