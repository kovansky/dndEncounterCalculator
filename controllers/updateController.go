package controllers

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/webview/webview"
)

func UpdateWindow(currentVersion models.AppVersionModel, remoteVersion models.AppVersionModel) {
	uw := webview.New(false)
	defer uw.Destroy()

	uw.SetTitle("Update avaliable!") // language
	uw.SetSize(600, 200, webview.HintFixed)

	err := uw.Bind("loadWindowState", func() string {
		var data = make(map[string]interface{})

		data["current"] = currentVersion
		data["remote"] = remoteVersion

		jsonData, err := json.Marshal(data)
		misc.Check(err)

		stringed := string(jsonData)

		return stringed
	})
	misc.Check(err)

	uw.Navigate("http://127.0.0.1:12354/update")

	uw.Run()
}
