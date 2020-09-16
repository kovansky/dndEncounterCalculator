package controllers

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
	"github.com/webview/webview"
)

var Party *models.PartyModel

func PartyWindow(wv webview.WebView) {
	wv.SetTitle("Create your party") // language
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
		MainWindow(wv)

		return true
	})
	misc.Check(err)

	err = wv.Bind("loadWindowState", func() string {
		if Party != nil {
			var playersAsArray []models.PlayerModel

			for _, player := range Party.PartyPlayers {
				playersAsArray = append(playersAsArray, player)
			}

			jsonP, err := json.Marshal(playersAsArray)
			misc.Check(err)

			stringed := string(jsonP)

			return stringed
		} else {
			return ""
		}
	})
	misc.Check(err)

	err = wv.Bind("runError", func() int {
		ch := make(chan int)
		go ErrorWindow(ch, models.ErrorModel{
			ErrorNumber:      1,
			ErrorDescription: "Oh no, an error!",
			ErrorType:        enum.ErrorEasy,
		})

		return <-ch
	})

	wv.Navigate("http://127.0.0.1:12342/party")
}
