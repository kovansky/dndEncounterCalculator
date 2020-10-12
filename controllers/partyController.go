package controllers

import (
	"encoding/json"
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
		var (
			// A variable to hold values to return
			ret = make(map[string]interface{})
			// Map of parties saved to disk
			saved map[string]models.PartySaveModel
			// Map of saved parties IDs and names
			idsMap = make(map[string]string)
		)

		// Check, if Party value was set. If it wasn't, this is the first run (there is no state to be restored)
		if Party != nil {
			var playersAsArray []models.PlayerModel

			// Convert players from map to slice
			for _, player := range Party.PartyPlayers {
				playersAsArray = append(playersAsArray, player)
			}

			// Add to return
			ret["party"] = playersAsArray
		} else {
			ret["party"] = ""
		}

		// Read saved parties from disk
		SavedParties.LoadData(&saved)

		if saved != nil {
			// Populate idsMap with saved parties
			for key, val := range saved {
				idsMap[key] = val.PartyName
			}
			// Add to return
			ret["partiesSelect"] = idsMap
		} else {
			ret["partiesSelect"] = map[string]string{}
		}

		// Marshal return value to json
		retJson, err := json.Marshal(ret)
		misc.Check(err)

		return string(retJson)
	})
	misc.Check(err)

	// Reads party from players form from Party View
	err = wv.Bind("readParty", func(modelString json.RawMessage) int {
		if Party == nil {
			Party = models.NewPartyModel()
		}

		var (
			// Slice of players
			model []models.PlayerModel

			// Size of previously held party, to calculate delta between players amount after updating with currently passed values
			countBefore = Party.CountPlayers()

			// Delta between previously held party size and party size after updating with currently passed value
			delta int

			// Map of "current players", players passed from form, to delete players from Party variable, that were deleted in view
			playersNames = map[string]bool{}
		)

		// Load players list from view into players slice
		json.Unmarshal(modelString, &model)

		// Check, if passed players list isn't empty
		if len(model) == 1 && model[0].PlayerName == "" && model[0].PlayerLevel == 0 {
			// If it is, return an error (error description in misc/errorsIndex.md)
			return -2001
		}

		// Add every player from list to Party
		for _, player := range model {
			if player.PlayerLevel == 0 {
				// Validate player level (cannot equal 0 nor be empty), error description in misc/errorsIndex.md
				return -2002
			} else if player.PlayerLevel < 1 {
				// Validate player level (cannot be less than 1), error description in misc/errorsIndex.md
				return -2003
			} else if player.PlayerName == "" {
				// Validate player name (cannot be empty), error description in misc/errorsIndex.md
				return -2004
			}

			// Remove old entry from Party (if existed), then add the new one
			Party.RemovePlayer(player.PlayerName)
			Party.AddPlayer(player)

			// Add to current players map
			playersNames[player.PlayerName] = true
		}

		// Remove from Party players not existing in current players map
		for _, player := range Party.PartyPlayers {
			if playersNames[player.PlayerName] != true {
				Party.RemovePlayer1(player)
			}
		}

		// Count delta between party size before and after update (read)
		if Party.CountPlayers() >= countBefore {
			delta = Party.CountPlayers() - countBefore
		} else {
			delta = countBefore - Party.CountPlayers()
		}

		return delta
	})
	misc.Check(err)

	// Writes party to disk (adds to saved parties)
	err = wv.Bind("writeParty", func(modelString json.RawMessage) int {
		var (
			// Data to save
			model = models.NewPartySaveModel()
			// Map of parties currently saved on disk
			oldModel map[string]models.PartySaveModel
		)

		// Load party data from view to save model
		json.Unmarshal(modelString, &model)

		for _, player := range model.PartyPlayers {
			if player.PlayerLevel == 0 {
				// Validate player level (cannot equal 0 nor be empty), error description in misc/errorsIndex.md
				return -2002
			} else if player.PlayerLevel < 1 {
				// Validate player level (cannot be less than 1), error description in misc/errorsIndex.md
				return -2003
			} else if player.PlayerName == "" {
				// Validate player name (cannot be empty), error description in misc/errorsIndex.md
				return -2004
			}
		}

		// Read currently saved parties from disk
		SavedParties.LoadData(&oldModel)

		if oldModel == nil {
			oldModel = make(map[string]models.PartySaveModel)
		}

		// Update party model by id
		oldModel[model.PartyId] = *model

		// Overwrite saved parties to disk
		SavedParties.WriteData(oldModel)

		return 0
	})
	misc.Check(err)

	// Loads party from disk to view by id
	err = wv.Bind("loadParty", func(partyId string) string {
		// Map of parties saved on disk
		var saved map[string]models.PartySaveModel

		// Read saved parties from disk
		SavedParties.LoadData(&saved)

		if saved != nil {
			// Check, if party with requested id exists on disk
			if value, found := saved[partyId]; found {
				// Marshal party data to json
				retJson, err := json.Marshal(value)
				misc.Check(err)

				return string(retJson)
			} else {
				// Return error, error description in misc/errorsIndex.md
				return "-2006"
			}
		} else {
			// Return error, error description in misc/errorsIndex.md
			return "-2006"
		}
	})
	misc.Check(err)

	// Removes party from saved parties file on disk by id
	err = wv.Bind("removeParty", func(partyId string) int {
		// Map of parties saved on disk
		var saved map[string]models.PartySaveModel

		// Read saved parties from disk
		SavedParties.LoadData(&saved)

		if saved != nil {
			// Check, if party with requested id exists on disk
			if _, found := saved[partyId]; found {
				// Delete party by id from map
				delete(saved, partyId)

				// Overwrite saved parties to disk
				SavedParties.WriteData(saved)

				return 1
			} else {
				// Return error, error description in misc/errorsIndex.md
				return -2006
			}
		} else {
			// Return error, error description in misc/errorsIndex.md
			return -2006
		}
	})
	misc.Check(err)

	// Navigates window to Main View
	err = wv.Bind("nextWindow", func() bool {
		MainWindow(wv)

		return true
	})
	misc.Check(err)

	// Opens Party View in window
	wv.Navigate("http://127.0.0.1:12356/party")
}
