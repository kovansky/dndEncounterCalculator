package controllers

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/webview/webview"
)

var Party models.PartyModel

func PartyWindow(wv webview.WebView) {
	wv.SetTitle("Create your party")
	wv.SetSize(600, 550, webview.HintFixed)

	err := wv.Bind("readParty", func(modelString json.RawMessage) int {
		//model := []models.PlayerModel{
		//	models.NewPlayerModel("Zbigniew", 1),
		//	models.NewPlayerModel("Mordka", 4),
		//}

		var model []models.PlayerModel

		json.Unmarshal(modelString, &model)

		p := models.NewPartyModel()

		for _, player := range model {
			p.AddPlayer(player)
		}

		return p.CountPlayers()
	})
	misc.Check(err)

	wv.Navigate("http://127.0.0.1:12345/party")
}
