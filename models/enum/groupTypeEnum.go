package enum

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
