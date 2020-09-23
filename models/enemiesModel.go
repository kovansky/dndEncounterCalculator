package models

import "github.com/kovansky/dndEncounterCalculator/models/enum"

//EnemiesModel is a model of whole enemies group
type EnemiesModel struct {
	GroupMonsters     map[string]MonsterModel
	GroupSize         int
	GroupType         enum.GroupType
	GroupXP           int
	GroupModCountType enum.GroupType // type of group with excluded monsters selected to not impede the encounter
}

//NewEnemiesModel returns empty EnemiesModel
func NewEnemiesModel() *EnemiesModel {
	return &EnemiesModel{GroupMonsters: make(map[string]MonsterModel)}
}

//Update calculates and sets all values, that depends on monsters map. It is run every time the monsters map is changed:
// - when monster is added
// - when monster is edited
// - when monster is removed
func (enemies *EnemiesModel) Update() EnemiesModel {
	enemies.GroupSize = enemies.CountSize()
	enemies.GroupType = enemies.GetGroupType()
	enemies.GroupXP = enemies.CalculateGroupXP()
	enemies.GroupModCountType = enemies.CalculateModCountType()

	return *enemies
}

//AddMonster adds monster to the group monsters map and runs Update
func (enemies *EnemiesModel) AddMonster(monster MonsterModel) EnemiesModel {
	enemies.GroupMonsters[monster.MonsterName] = monster

	return enemies.Update()
}

//RemoveMonster removes monster from the group monsters map by monster name and runs Update
func (enemies *EnemiesModel) RemoveMonster(monster string) EnemiesModel {
	delete(enemies.GroupMonsters, monster)

	return enemies.Update()
}

//RemoveMonster1 removes monster from the group monsters map by monster model and runs Update
func (enemies *EnemiesModel) RemoveMonster1(monster MonsterModel) EnemiesModel {
	return enemies.RemoveMonster(monster.MonsterName)
}

//CountSize counts all monsters in group
func (enemies *EnemiesModel) CountSize() int {
	var size = 0

	// Add each monster amount to group size
	for _, monster := range enemies.GroupMonsters {
		size += monster.MonstersAmount
	}

	return size
}

//GetGroupType returns group type based on group size
func (enemies *EnemiesModel) GetGroupType() enum.GroupType {
	return enum.GroupTypeByAmount(enemies.GroupSize)
}

//CalculateGroupXP sums up all monsters XP values
func (enemies *EnemiesModel) CalculateGroupXP() int {
	var xpValue = 0

	// Add each monster XP value to group XP value
	for _, monster := range enemies.GroupMonsters {
		xpValue += monster.GroupXP
	}

	return xpValue
}

//CalculateModCountType returns group type based on amount of monsters, excluding these selected to not impede the encounter
func (enemies *EnemiesModel) CalculateModCountType() enum.GroupType {
	count := 0

	for _, monster := range enemies.GroupMonsters {
		// Check, if monsters should be counted in difficulty
		if monster.CountInCRMod {
			// Add to group size
			count += monster.MonstersAmount
		}
	}

	return enum.GroupTypeByAmount(count)
}
