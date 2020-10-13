/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package enum

//EncounterModifier specifies possible encounter modifiers
type EncounterModifier float32

const (
	ModifierBig    EncounterModifier = 0.5
	ModifierSingle EncounterModifier = 1
	ModifierPair   EncounterModifier = 1.5
	ModifierGroup  EncounterModifier = 2
	ModifierGang   EncounterModifier = 2.5
	ModifierMob    EncounterModifier = 3
	ModifierHorde  EncounterModifier = 4
	ModifierSmall  EncounterModifier = 5
)

//CalculateEncounterModificator compares enemies group size and returns corresponding modifier, considering the party size
func CalculateEncounterModificator(partySize PartyCategory, groupType GroupType) EncounterModifier {
	switch partySize {
	case PartySmall:
		switch groupType {
		case MonsterSingle:
			return ModifierPair
		case MonsterPair:
			return ModifierGroup
		case MonsterGroup:
			return ModifierGang
		case MonsterGang:
			return ModifierMob
		case MonsterMob:
			return ModifierHorde
		case MonsterHorde:
			return ModifierSmall
		default:
			return ModifierPair
		}
	case PartyStandard:
		switch groupType {
		case MonsterSingle:
			return ModifierSingle
		case MonsterPair:
			return ModifierPair
		case MonsterGroup:
			return ModifierGroup
		case MonsterGang:
			return ModifierGang
		case MonsterMob:
			return ModifierMob
		case MonsterHorde:
			return ModifierHorde
		default:
			return ModifierSingle
		}
	case PartyBig:
		switch groupType {
		case MonsterSingle:
			return ModifierBig
		case MonsterPair:
			return ModifierSingle
		case MonsterGroup:
			return ModifierPair
		case MonsterGang:
			return ModifierGroup
		case MonsterMob:
			return ModifierGang
		case MonsterHorde:
			return ModifierMob
		default:
			return ModifierBig
		}
	default:
		return ModifierSingle
	}
}
