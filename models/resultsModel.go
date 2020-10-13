/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package models

import "github.com/kovansky/dndEncounterCalculator/models/enum"

//ResultsModel is a model of return message, holding encounter difficulty calculations and XP data
type ResultsModel struct {
	MonstersAmount      int                      `json:"monsters_amount"`
	MonstersGroupType   string                   `json:"monsters_group_type"`
	Award               int                      `json:"award"`
	CapoAward           float32                  `json:"capo_award"`
	DifficultyModifier  enum.EncounterModifier   `json:"difficulty_modifier"`
	AdjustedXP          float32                  `json:"adjusted_xp"`
	EncounterDifficulty enum.EncounterDifficulty `json:"encounter_difficulty"`
}
