package models

import "github.com/kovansky/dndEncounterCalculator/models/enum"

type EnemiesModel struct {
	GroupMonsters     map[string]MonsterModel
	GroupSize         int
	GroupType         enum.GroupType
	GroupXP           int
	GroupModCountType enum.GroupType
}

func NewEnemiesModel() *EnemiesModel {
	return &EnemiesModel{GroupMonsters: make(map[string]MonsterModel)}
}

func (enemies *EnemiesModel) Update() EnemiesModel {
	enemies.GroupSize = enemies.CountSize()
	enemies.GroupType = enemies.GetGroupType()
	enemies.GroupXP = enemies.CalculateGroupXP()
	enemies.GroupModCountType = enemies.CalculateModCountType()

	return *enemies
}

func (enemies *EnemiesModel) AddMonster(monster MonsterModel) EnemiesModel {
	enemies.GroupMonsters[monster.MonsterName] = monster

	return enemies.Update()
}

func (enemies *EnemiesModel) RemoveMonster(monster string) EnemiesModel {
	delete(enemies.GroupMonsters, monster)

	return enemies.Update()
}

func (enemies *EnemiesModel) RemoveMonster1(monster MonsterModel) EnemiesModel {
	return enemies.RemoveMonster(monster.MonsterName)
}

func (enemies *EnemiesModel) CountSize() int {
	var size = 0

	for _, monster := range enemies.GroupMonsters {
		size += monster.MonstersAmount
	}

	return size
}

func (enemies *EnemiesModel) GetGroupType() enum.GroupType {
	return enum.GroupTypeByAmount(enemies.GroupSize)
}

func (enemies *EnemiesModel) CalculateGroupXP() int {
	var xpValue = 0

	for _, monster := range enemies.GroupMonsters {
		xpValue += monster.GroupXP
	}

	return xpValue
}

func (enemies *EnemiesModel) CalculateModCountType() enum.GroupType {
	count := 0

	for _, monster := range enemies.GroupMonsters {
		if monster.CountInCRMod {
			count += monster.MonstersAmount
		}
	}

	return enum.GroupTypeByAmount(count)
}
