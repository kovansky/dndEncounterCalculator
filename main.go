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

	//controllers.PartyWindow(wv)
	controllers.MainWindow(wv) // ToDo: ONLY FOR TESTING

	wv.Run()
}
