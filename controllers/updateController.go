package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/pkg/browser"
	"github.com/webview/webview"
)

//UpdateWindow is a controller function of Update View (dialog). It creates a WebView window
func UpdateWindow(currentVersion models.AppVersionModel, remoteVersion models.AppVersionModel) {
	// Create webview window, and defer destroying it
	uw := webview.New(false)
	defer uw.Destroy()

	// Adjust window data to view
	uw.SetTitle("Update avaliable!") // language
	uw.SetSize(600, 200, webview.HintFixed)

	// On view open passes data from backend to view
	err := uw.Bind("loadWindowState", func() string {
		// Holds data to return
		var data = make(map[string]interface{})

		// Adds versions information to return
		data["current"] = currentVersion
		data["remote"] = remoteVersion

		// Marshal return data to json
		jsonData, err := json.Marshal(data)
		misc.Check(err)

		return string(jsonData)
	})
	misc.Check(err)

	// Dialog controls (buttons) logic
	err = uw.Bind("retValue", func(code int) int {
		// If "yes" ("download update") button clicked, open update URL in browser
		if code == 1 {
			url := fmt.Sprintf(constants.APP_UPDATE_URL, remoteVersion.String())

			browser.OpenURL(url)
		}

		// Close dialog window, regardless of button clicked
		uw.Terminate()

		return code
	})
	misc.Check(err)

	// Opens Update View in window
	uw.Navigate("http://127.0.0.1:12354/update")

	// Runs window code
	uw.Run()
}
