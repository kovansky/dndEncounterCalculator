/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package functions

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
)

func LoadMainState(enemies *models.EnemiesModel) string {
	// Check, if enemies value was set. If it wasn't, this is the first run (there is no state to be restored)
	if enemies != nil {
		var monstersAsArray []models.MonsterModel

		// Convert monsters from map to slice
		for _, monster := range enemies.GroupMonsters {
			monstersAsArray = append(monstersAsArray, monster)
		}

		// Marshal monsters array into json
		jsonE, err := json.Marshal(monstersAsArray)
		misc.Check(err)

		stringed := string(jsonE)

		return stringed
	} else {
		return ""
	}
}

func GetPartyData(party *models.PartyModel) string {
	// Marshal Party data to json
	jsonParty, err := json.Marshal(party)
	misc.Check(err)

	return string(jsonParty)
}

func CalculateResults(monstersString json.RawMessage, enemies *models.EnemiesModel, party *models.PartyModel) string {
	// Reset enemies list
	enemies = models.NewEnemiesModel()

	var (
		// Slice of monsters
		monsters []models.MonsterModel

		// Difficulty modifier, depends on monsters amount and party size
		modifier enum.EncounterModifier

		// XP value after applying difficulty modifier
		adjustedXP float32

		// Encounter difficulty
		difficulty enum.EncounterDifficulty

		// Results of calculation
		results models.ResultsModel
	)

	// Load monsters list from view into monsters slice
	json.Unmarshal(monstersString, &monsters)

	// Add monsters to enemies variable
	for _, monster := range monsters {
		monster.Update()

		// If monster name is null, skip it
		if len(monster.MonsterName) == 0 {
			continue
		}

		// Add monster to struct
		enemies.AddMonster(monster)
	}

	// Calculate things
	modifier = enum.CalculateEncounterModificator(party.PartyCategory, enemies.GroupModCountType)
	adjustedXP = float32(enemies.GroupXP) * float32(modifier)
	difficulty = enum.CalculateEncounterDifficulty(party.PartyThresholds, adjustedXP)

	// Pack everything to results model
	results = models.ResultsModel{
		MonstersAmount:      enemies.GroupSize,
		MonstersGroupType:   enum.GroupTypeName(enemies.GroupType),
		Award:               enemies.GroupXP,
		CapoAward:           float32(enemies.GroupXP) / float32(party.PartySize),
		DifficultyModifier:  modifier,
		AdjustedXP:          adjustedXP,
		EncounterDifficulty: difficulty,
	}

	ret, err := json.Marshal(results)
	misc.Check(err)

	return string(ret)
}
