package main

import (
	"github.com/kovansky/dndEncounterCalculator/controllers"
	"github.com/kovansky/dndEncounterCalculator/webapp"
	"github.com/webview/webview"
)

func main() {
	go webapp.App()

	wv := webview.New(true)
	defer wv.Destroy()

	controllers.PartyWindow(wv)

	wv.Run()
}
