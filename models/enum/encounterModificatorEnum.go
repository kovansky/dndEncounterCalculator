package enum

type EncounterModificator float32

const (
	ModificatorBig    EncounterModificator = 0.5
	ModificatorSingle EncounterModificator = 1
	ModificatorPair   EncounterModificator = 1.5
	ModificatorGroup  EncounterModificator = 2
	ModificatorGang   EncounterModificator = 2.5
	ModificatorMob    EncounterModificator = 3
	ModificatorHorde  EncounterModificator = 4
	ModificatorSmall  EncounterModificator = 5
)

func CalculateEncounterModificator(partySize PartyCategory, groupType GroupType) EncounterModificator {
	switch partySize {
	case PartySmall:
		switch groupType {
		case MonsterSingle:
			return ModificatorPair
		case MonsterPair:
			return ModificatorGroup
		case MonsterGroup:
			return ModificatorGang
		case MonsterGang:
			return ModificatorMob
		case MonsterMob:
			return ModificatorHorde
		case MonsterHorde:
			return ModificatorSmall
		default:
			return ModificatorPair
		}
	case PartyStandard:
		switch groupType {
		case MonsterSingle:
			return ModificatorSingle
		case MonsterPair:
			return ModificatorPair
		case MonsterGroup:
			return ModificatorGroup
		case MonsterGang:
			return ModificatorGang
		case MonsterMob:
			return ModificatorMob
		case MonsterHorde:
			return ModificatorHorde
		default:
			return ModificatorSingle
		}
	case PartyBig:
		switch groupType {
		case MonsterSingle:
			return ModificatorBig
		case MonsterPair:
			return ModificatorSingle
		case MonsterGroup:
			return ModificatorPair
		case MonsterGang:
			return ModificatorGroup
		case MonsterMob:
			return ModificatorGang
		case MonsterHorde:
			return ModificatorMob
		default:
			return ModificatorBig
		}
	default:
		return ModificatorSingle
	}
}
