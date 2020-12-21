/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package lorca

import (
	"fmt"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/controllers/functions"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/pkg/browser"
	"github.com/zserge/lorca"
)

//LUpdateWindow is a controller function of Update View (dialog), but for Lorca instead of WebView. It creates a Lorca window
func LUpdateWindow(currentVersion models.AppVersionModel, remoteVersion models.AppVersionModel) {
	ui, _ := lorca.New("", "", 600, 250)
	defer ui.Close()

	// On view open passes data from backend to view
	err := ui.Bind("loadWindowState", func() string {
		return functions.LoadUpdateState(currentVersion, remoteVersion)
	})
	misc.Check(err)

	// Dialog controls (buttons) logic
	err = ui.Bind("retValue", func(code int) int {
		// If "yes" ("download update") button clicked, open update URL in browser
		if code == 1 {
			url := fmt.Sprintf(constants.APP_UPDATE_URL, remoteVersion.StringNoChannel())

			browser.OpenURL(url)
		}

		// Close dialog window, regardless of button clicked
		ui.Close()

		return code
	})
	misc.Check(err)

	// Opens Update View in window
	ui.Load("http://" + constants.APP_WEBAPP_URL + "/public/html/update.html")

	// Wait until window is closed
	<-ui.Done()
}
