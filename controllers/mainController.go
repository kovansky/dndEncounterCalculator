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
	"github.com/kovansky/dndEncounterCalculator/controllers/functions"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
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
		return functions.LoadMainState(enemies)
	})
	misc.Check(err)

	// Returns Party data as json
	err = wv.Bind("getPartyData", func() string {
		return functions.GetPartyData(Party)
	})
	misc.Check(err)

	// Calculates encounter difficulty and XP values
	err = wv.Bind("calculateResults", func(monstersString json.RawMessage) string {
		return functions.CalculateResults(monstersString, enemies, Party)
	})
	misc.Check(err)

	// Navigates window to Party View
	err = wv.Bind("editParty", func() bool {
		PartyWindow(wv)

		return true
	})
	misc.Check(err)

	// Opens Main View in window
	wv.Navigate("http://" + constants.APP_WEBAPP_URL + "/public/html/main.html")
}
