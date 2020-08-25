package controllers

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/webview/webview"
)

var Party *models.PartyModel

func PartyWindow(wv webview.WebView) {
	wv.SetTitle("Create your party")
	wv.SetSize(600, 550, webview.HintFixed)

	err := wv.Bind("readParty", func(modelString json.RawMessage) int {
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

		for _, player := range model {
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

	err = wv.Bind("nextWindow", func() bool {
		// ToDo: open next window

		wv.Destroy()

		return true
	})

	wv.Navigate("http://127.0.0.1:12345/party")
}
