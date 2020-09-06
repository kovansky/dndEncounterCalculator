package models

import "github.com/kovansky/dndEncounterCalculator/models/enum"

type ResultsModel struct {
	MonstersAmount      int                      `json:"monsters_amount"`
	MonstersGroupType   string                   `json:"monsters_group_type"`
	Award               int                      `json:"award"`
	CapoAward           float32                  `json:"capo_award"`
	DifficultyModifier  enum.EncounterModifier   `json:"difficulty_modifier"`
	AdjustedXP          float32                  `json:"adjusted_xp"`
	EncounterDifficulty enum.EncounterDifficulty `json:"encounter_difficulty"`
}
