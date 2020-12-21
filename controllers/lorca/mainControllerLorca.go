/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package lorca

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/controllers/functions"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
)
import "github.com/zserge/lorca"

//enemies is a variable that holds all monsters added in the monsters form in the Main View.
var enemies *models.EnemiesModel

//LMainWindow is a controller function of Main View, but for Lorca instead of WebView. It loads a view into existing WebView window
func LMainWindow(ui lorca.UI) {
	currBounds, _ := ui.Bounds()
	err := ui.SetBounds(lorca.Bounds{
		Left:        currBounds.Left,
		Top:         currBounds.Top,
		Width:       1000,
		Height:      725,
		WindowState: "normal",
	})
	misc.Check(err)

	// On view open loads previous state (monsters list), if existed.
	err = ui.Bind("loadWindowState", func() string {
		return functions.LoadMainState(enemies)
	})
	misc.Check(err)

	// Returns Party data as json
	err = ui.Bind("getPartyData", func() string {
		return functions.GetPartyData(Party)
	})
	misc.Check(err)

	// Calculates encounter difficulty and XP values
	err = ui.Bind("calculateResults", func(monstersString json.RawMessage) string {
		return functions.CalculateResults(monstersString, enemies, Party)
	})
	misc.Check(err)

	// Navigates window to Party View
	err = ui.Bind("editParty", func() bool {
		LPartyWindow(ui)

		return true
	})
	misc.Check(err)

	ui.Load("http://" + constants.APP_WEBAPP_URL + "/public/html/main.html")
}
