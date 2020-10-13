/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package enum

//GroupType specifies monster group size types. The constants are holding the types low thresholds
type GroupType int

const (
	MonsterGroupError GroupType = -1
	MonsterSingle     GroupType = 1
	MonsterPair       GroupType = 2
	MonsterGroup      GroupType = 3
	MonsterGang       GroupType = 7
	MonsterMob        GroupType = 11
	MonsterHorde      GroupType = 15
)

//GroupTypeByAmount compares given size with possible group types
func GroupTypeByAmount(amount int) GroupType {
	if amount == 1 {
		return MonsterSingle
	} else if amount == 2 {
		return MonsterPair
	} else if amount >= 3 && amount < 7 {
		return MonsterGroup
	} else if amount >= 7 && amount < 11 {
		return MonsterGang
	} else if amount >= 11 && amount < 15 {
		return MonsterMob
	} else if amount >= 15 {
		return MonsterHorde
	} else {
		return MonsterGroupError
	}
}

//GroupTypeName returns group type name
func GroupTypeName(groupType GroupType) string {
	// language
	switch groupType {
	case MonsterGroupError:
		return "error"
	case MonsterSingle:
		return "single"
	case MonsterPair:
		return "pair"
	case MonsterGroup:
		return "group"
	case MonsterGang:
		return "gang"
	case MonsterMob:
		return "mob"
	case MonsterHorde:
		return "horde"
	default:
		return "error"
	}
}
