package main

import (
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
		isUpdate, rMajor, rMinor, rPath, rChannel := appVersion.CheckForUpdates()
		remoteAvm := models.AppVersionModel{
			Major:   rMajor,
			Minor:   rMinor,
			Path:    rPath,
			Channel: rChannel,
		}

		if isUpdate {
			controllers.UpdateWindow(*appVersion, remoteAvm)
		}
	}()

	wv := webview.New(true)
	defer wv.Destroy()

	err := wv.Bind("runError", misc.ThrowError)
	misc.Check(err)

	controllers.PartyWindow(wv)

	wv.Run()
}
