package models

type MonsterModel struct {
	MonsterName    string
	MonsterCR      int
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
	// ToDo: calculate monster XP value, based on CR-XP relation (CRXP constant)
	return 0
}

func (monster *MonsterModel) CalculateGroupXP() int {
	return monster.MonsterXP * monster.MonstersAmount
}
