package controllers

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/webview/webview"
)

var Party *models.PartyModel
var SavedParties *misc.DataFile

func PartyWindow(wv webview.WebView) {
	SavedParties = misc.NewDataFile("parties.json").CheckFile()

	wv.SetTitle("Create your party") // language
	wv.SetSize(600, 750, webview.HintFixed)

	err := wv.Bind("loadWindowState", func() string {
		var (
			ret    = make(map[string]interface{})
			saved  map[string]models.PartySaveModel
			idsMap = make(map[string]string)
		)

		if Party != nil {
			var playersAsArray []models.PlayerModel

			for _, player := range Party.PartyPlayers {
				playersAsArray = append(playersAsArray, player)
			}

			ret["party"] = playersAsArray
		} else {
			ret["party"] = ""
		}

		SavedParties.LoadData(&saved)
		for key, val := range saved {
			idsMap[key] = val.PartyName
		}
		ret["partiesSelect"] = idsMap

		retJson, err := json.Marshal(ret)
		misc.Check(err)

		return string(retJson)
	})
	misc.Check(err)

	err = wv.Bind("readParty", func(modelString json.RawMessage) int {
		if Party == nil {
			Party = models.NewPartyModel()
		}

		var (
			model        []models.PlayerModel
			countBefore  = Party.CountPlayers()
			delta        int
			playersNames = map[string]bool{}
		)

		json.Unmarshal(modelString, &model)

		if len(model) == 1 {
			if model[0].PlayerName == "" && model[0].PlayerLevel == 0 {
				return -2001
			}
		}

		for _, player := range model {
			if player.PlayerLevel == 0 {
				Party = models.NewPartyModel()
				return -2002
			} else if player.PlayerLevel < 1 {
				Party = models.NewPartyModel()
				return -2003
			} else if player.PlayerName == "" {
				Party = models.NewPartyModel()
				return -2004
			}

			Party.RemovePlayer(player.PlayerName)
			Party.AddPlayer(player)

			playersNames[player.PlayerName] = true
		}

		for _, player := range Party.PartyPlayers {
			if playersNames[player.PlayerName] != true {
				Party.RemovePlayer1(player)
			}
		}

		if Party.CountPlayers() >= countBefore {
			delta = Party.CountPlayers() - countBefore
		} else {
			delta = countBefore - Party.CountPlayers()
		}

		return delta
	})
	misc.Check(err)

	err = wv.Bind("writeParty", func(modelString json.RawMessage) int {
		var (
			model    = models.NewPartySaveModel()
			oldModel map[string]models.PartySaveModel
		)

		json.Unmarshal(modelString, &model)

		for _, player := range model.PartyPlayers {
			if player.PlayerLevel == 0 {
				Party = models.NewPartyModel()
				return -2002
			} else if player.PlayerLevel < 1 {
				Party = models.NewPartyModel()
				return -2003
			} else if player.PlayerName == "" {
				Party = models.NewPartyModel()
				return -2004
			}
		}

		SavedParties.LoadData(&oldModel)

		oldModel[model.PartyId] = *model

		SavedParties.WriteData(oldModel)

		return 0
	})
	misc.Check(err)

	err = wv.Bind("loadParty", func(partyId string) string {
		var saved map[string]models.PartySaveModel

		SavedParties.LoadData(&saved)

		if value, found := saved[partyId]; found {
			retJson, err := json.Marshal(value)
			misc.Check(err)

			return string(retJson)
		} else {
			return "-2006"
		}
	})

	err = wv.Bind("nextWindow", func() bool {
		MainWindow(wv)

		return true
	})
	misc.Check(err)

	wv.Navigate("http://127.0.0.1:12351/party")
}
