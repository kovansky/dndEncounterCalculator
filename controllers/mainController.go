/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

/*
Package controllers implements Controllers of application Views, their logic and communication between them and Views.
*/
package controllers

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
	"github.com/webview/webview"
)

//enemies is a variable that holds all monsters added in the monsters form in the Main View.
var enemies *models.EnemiesModel

//MainWindow is a controller function of Main View. It loads a view into existing WebView window
func MainWindow(wv webview.WebView) {
	// Adjust window data to view
	wv.SetTitle("D&D Encounter Calculator") // language
	wv.SetSize(1000, 675, webview.HintFixed)

	// On view open loads previous state (monsters list), if existed.
	err := wv.Bind("loadWindowState", func() string {
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
	})
	misc.Check(err)

	// Returns Party data as json
	err = wv.Bind("getPartyData", func() string {
		// Marshal Party data to json
		jsonParty, err := json.Marshal(Party)
		misc.Check(err)

		return string(jsonParty)
	})
	misc.Check(err)

	// Calculates encounter difficulty and XP values
	err = wv.Bind("calculateResults", func(monstersString json.RawMessage) string {
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
		modifier = enum.CalculateEncounterModificator(Party.PartyCategory, enemies.GroupModCountType)
		adjustedXP = float32(enemies.GroupXP) * float32(modifier)
		difficulty = enum.CalculateEncounterDifficulty(Party.PartyThresholds, adjustedXP)

		// Pack everything to results model
		results = models.ResultsModel{
			MonstersAmount:      enemies.GroupSize,
			MonstersGroupType:   enum.GroupTypeName(enemies.GroupType),
			Award:               enemies.GroupXP,
			CapoAward:           float32(enemies.GroupXP) / float32(Party.PartySize),
			DifficultyModifier:  modifier,
			AdjustedXP:          adjustedXP,
			EncounterDifficulty: difficulty,
		}

		ret, err := json.Marshal(results)
		misc.Check(err)

		return string(ret)
	})
	misc.Check(err)

	// Navigates window to Party View
	err = wv.Bind("editParty", func() bool {
		PartyWindow(wv)

		return true
	})
	misc.Check(err)

	// Opens Main View in window
	wv.Navigate("http://" + constants.APP_WEBAPP_URL + "/main")
}
