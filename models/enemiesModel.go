package models

import "github.com/kovansky/dndEncounterCalculator/models/enum"

type EnemiesModel struct {
	GroupMonsters map[string]MonsterModel
	GroupSize     int
	GroupType     enum.GroupType
	GroupXP       int
}

func NewEnemiesModel() *EnemiesModel {
	return &EnemiesModel{GroupMonsters: make(map[string]MonsterModel)}
}

func (enemies *EnemiesModel) Update() EnemiesModel {

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
