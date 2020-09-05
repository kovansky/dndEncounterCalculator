package controllers

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/webview/webview"
)

func MainWindow(wv webview.WebView) {
	wv.SetTitle("D&D Encounter Calculator") // language
	wv.SetSize(1000, 800, webview.HintFixed)

	jsonParty, err := json.Marshal(Party)
	misc.Check(err)

	err = wv.Bind("getPartyData", func() string {
		return string(jsonParty)
	})
	misc.Check(err)

	wv.Navigate("http://127.0.0.1:12348/main")
}
