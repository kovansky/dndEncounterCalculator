package main

import (
	"fmt"
	"github.com/kovansky/dndEncounterCalculator/controllers"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/kovansky/dndEncounterCalculator/webapp"
	"github.com/webview/webview"
)

func main() {
	go webapp.App()
	go func() {
		appVersion := models.GetAppVersion()
		isUpdate, _, _, _, _ := appVersion.CheckForUpdates()

		if isUpdate {
			fmt.Println("There is an update avaliable")
		}
	}()

	wv := webview.New(true)
	defer wv.Destroy()

	err := wv.Bind("runError", misc.ThrowError)
	misc.Check(err)

	controllers.PartyWindow(wv)

	wv.Run()
}
