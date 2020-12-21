/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
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

var (
	//Party is a variable holding list of players and their levels
	Party *models.PartyModel
	//SavedParties is a variable to read the parties saved by the user to a file
	SavedParties *misc.DataFile
)

//PartyWindow is a controller function of Party View. It loads a view into existing WebView window
func PartyWindow(wv webview.WebView) {
	// Point to a file with saved parties data
	SavedParties = misc.NewDataFile("parties.json").CheckFile()

	// Adjust window data to view
	wv.SetTitle("Create your party") // language
	wv.SetSize(600, 750, webview.HintFixed)

	// On view open loads previous state (players list), if existed
	err := wv.Bind("loadWindowState", func() string {
		return functions.LoadPartyState(Party, SavedParties)
	})
	misc.Check(err)

	// Reads party from players form from Party View
	err = wv.Bind("readParty", func(modelString json.RawMessage) int {
		_return, _party := functions.ReadParty(modelString, Party)

		if _party != nil {
			Party = _party
		}

		return _return
	})
	misc.Check(err)

	// Writes party to disk (adds to saved parties)
	err = wv.Bind("writeParty", func(modelString json.RawMessage) int {
		_return, _savedParties := functions.WriteParty(modelString, SavedParties)

		if _savedParties != nil {
			SavedParties = _savedParties
		}

		return _return
	})
	misc.Check(err)

	// Loads party from disk to view by id
	err = wv.Bind("loadParty", func(partyId string) string {
		return functions.LoadParty(partyId, SavedParties)
	})
	misc.Check(err)

	// Removes party from saved parties file on disk by id
	err = wv.Bind("removeParty", func(partyId string) int {
		_return, _savedParties := functions.RemoveParty(partyId, SavedParties)

		if _savedParties != nil {
			SavedParties = _savedParties
		}

		return _return
	})
	misc.Check(err)

	// Navigates window to Main View
	err = wv.Bind("nextWindow", func() bool {
		MainWindow(wv)

		return true
	})
	misc.Check(err)

	// Opens Party View in window
	wv.Navigate("http://" + constants.APP_WEBAPP_URL + "/public/html/party.html")
}
