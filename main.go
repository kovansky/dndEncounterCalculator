/*
Package main holds application startup code
*/
package main

// ToDo: add copyright

import (
	"github.com/kovansky/dndEncounterCalculator/controllers"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/kovansky/dndEncounterCalculator/webapp"
	"github.com/webview/webview"
)

//main is the heart of application, the runner of the rest of the program
func main() {
	// Run webapp in separate goroutine
	go webapp.App()
	// Check for updates in separate goroutine
	go func() {
		// Get current application version
		appVersion := models.GetAppVersion()
		// Check for updates from remote
		isUpdate, rMajor, rMinor, rPatch, rChannel := appVersion.CheckForUpdates()
		// Create AppVersionModel from remote data
		remoteAvm := models.AppVersionModel{
			Major:   rMajor,
			Minor:   rMinor,
			Patch:   rPatch,
			Channel: rChannel,
		}

		// If update avaliable, open Update Dialog
		if isUpdate {
			controllers.UpdateWindow(*appVersion, remoteAvm)
		}
	}()

	// Create new webview instance and defer destroying it
	wv := webview.New(true)
	defer wv.Destroy()

	// Bind runError function for JS to be avaliable in all views using this webview instance
	err := wv.Bind("runError", misc.ThrowError)
	misc.Check(err)

	// Run first window - party window
	controllers.PartyWindow(wv)

	// Runs window code
	wv.Run()
}
