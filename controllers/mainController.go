package controllers

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
	"github.com/webview/webview"
)

var enemies *models.EnemiesModel

func MainWindow(wv webview.WebView) {
	wv.SetTitle("D&D Encounter Calculator") // language
	wv.SetSize(1000, 675, webview.HintFixed)

	jsonParty, err := json.Marshal(Party)
	misc.Check(err)

	err = wv.Bind("getPartyData", func() string {
		return string(jsonParty)
	})
	misc.Check(err)

	err = wv.Bind("calculateResults", func(monstersString json.RawMessage) string {
		if enemies == nil {
			enemies = models.NewEnemiesModel()
		}

		var (
			monsters   []models.MonsterModel
			modifier   enum.EncounterModifier
			adjustedXP float32
			difficulty enum.EncounterDifficulty
			results    models.ResultsModel
		)

		json.Unmarshal(monstersString, &monsters)

		for _, monster := range monsters {
			monster.Update()
			enemies.AddMonster(monster)
		}

		modifier = enum.CalculateEncounterModificator(Party.PartyCategory, enemies.GroupModCountType)
		adjustedXP = float32(enemies.GroupXP) * float32(modifier)
		difficulty = enum.CalculateEncounterDifficulty(Party.PartyThresholds, adjustedXP)

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

	err = wv.Bind("loadWindowState", func() string {
		if enemies != nil {
			var monstersAsArray []models.MonsterModel

			for _, monster := range enemies.GroupMonsters {
				monstersAsArray = append(monstersAsArray, monster)
			}

			jsonE, err := json.Marshal(monstersAsArray)
			misc.Check(err)

			stringed := string(jsonE)

			return stringed
		} else {
			return ""
		}
	})
	misc.Check(err)

	err = wv.Bind("editParty", func() bool {
		PartyWindow(wv)

		return true
	})
	misc.Check(err)

	wv.Navigate("http://127.0.0.1:12360/main")
}
