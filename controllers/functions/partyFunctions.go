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
)

func LoadPartyState(party *models.PartyModel, savedParties *misc.DataFile) string {
	var (
		// A variable to hold values to return
		ret = make(map[string]interface{})
		// Map of parties saved to disk
		saved map[string]models.PartySaveModel
		// Map of saved parties IDs and names
		idsMap = make(map[string]string)
	)

	// Check, if Party value was set. If it wasn't, this is the first run (there is no state to be restored)
	if party != nil {
		var playersAsArray []models.PlayerModel

		// Convert players from map to slice
		for _, player := range party.PartyPlayers {
			playersAsArray = append(playersAsArray, player)
		}

		// Add to return
		ret["party"] = playersAsArray
	} else {
		ret["party"] = ""
	}

	// Read saved parties from disk
	savedParties.LoadData(&saved)

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
}

func ReadParty(modelString json.RawMessage, party *models.PartyModel) (int, *models.PartyModel) {
	if party == nil {
		party = models.NewPartyModel()
	}

	var (
		// Slice of players
		model []models.PlayerModel

		// Size of previously held party, to calculate delta between players amount after updating with currently passed values
		countBefore = party.CountPlayers()

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
		return -2001, nil
	}

	// Add every player from list to Party
	for _, player := range model {
		if player.PlayerLevel == 0 {
			// Validate player level (cannot equal 0 nor be empty), error description in misc/errorsIndex.md
			return -2002, nil
		} else if player.PlayerLevel < 1 {
			// Validate player level (cannot be less than 1), error description in misc/errorsIndex.md
			return -2003, nil
		} else if player.PlayerName == "" {
			// Validate player name (cannot be empty), error description in misc/errorsIndex.md
			return -2004, nil
		}

		// Remove old entry from Party (if existed), then add the new one
		party.RemovePlayer(player.PlayerName)
		party.AddPlayer(player)

		// Add to current players map
		playersNames[player.PlayerName] = true
	}

	// Remove from Party players not existing in current players map
	for _, player := range party.PartyPlayers {
		if playersNames[player.PlayerName] != true {
			party.RemovePlayer1(player)
		}
	}

	// Count delta between party size before and after update (read)
	if party.CountPlayers() >= countBefore {
		delta = party.CountPlayers() - countBefore
	} else {
		delta = countBefore - party.CountPlayers()
	}

	return delta, party
}

func WriteParty(modelString json.RawMessage, savedParties *misc.DataFile) (int, *misc.DataFile) {
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
			return -2002, nil
		} else if player.PlayerLevel < 1 {
			// Validate player level (cannot be less than 1), error description in misc/errorsIndex.md
			return -2003, nil
		} else if player.PlayerName == "" {
			// Validate player name (cannot be empty), error description in misc/errorsIndex.md
			return -2004, nil
		}
	}

	// Read currently saved parties from disk
	savedParties.LoadData(&oldModel)

	if oldModel == nil {
		oldModel = make(map[string]models.PartySaveModel)
	}

	// Update party model by id
	oldModel[model.PartyId] = *model

	// Overwrite saved parties to disk
	savedParties.WriteData(oldModel)

	return 0, savedParties
}

func LoadParty(partyId string, savedParties *misc.DataFile) string {
	// Map of parties saved on disk
	var saved map[string]models.PartySaveModel

	// Read saved parties from disk
	savedParties.LoadData(&saved)

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
}

func RemoveParty(partyId string, savedParties *misc.DataFile) (int, *misc.DataFile) {
	// Map of parties saved on disk
	var saved map[string]models.PartySaveModel

	// Read saved parties from disk
	savedParties.LoadData(&saved)

	if saved != nil {
		// Check, if party with requested id exists on disk
		if _, found := saved[partyId]; found {
			// Delete party by id from map
			delete(saved, partyId)

			// Overwrite saved parties to disk
			savedParties.WriteData(saved)

			return 1, savedParties
		} else {
			// Return error, error description in misc/errorsIndex.md
			return -2006, nil
		}
	} else {
		// Return error, error description in misc/errorsIndex.md
		return -2006, nil
	}
}
