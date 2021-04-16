// +build lorca

/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

/*
Package main holds application startup code
*/
package main

import (
	controllers "github.com/kovansky/dndEncounterCalculator/controllers/lorca"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/kovansky/dndEncounterCalculator/webapp"
	"github.com/zserge/lorca"
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
			controllers.LUpdateWindow(*appVersion, remoteAvm)
		}
	}()

	// Create lorca instance and defer closing it
	ui, _ := lorca.New("", "", 100, 100)
	defer ui.Close()

	// Bind runError function for JS to be avaliable in all views using this lorca instance
	err := ui.Bind("runError", misc.ThrowError)
	misc.Check(err)

	// Run first window
	controllers.LPartyWindow(ui)

	// Wait until done
	<-ui.Done()
}
