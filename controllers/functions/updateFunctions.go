/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package functions

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models"
)

func LoadUpdateState(currentVersion models.AppVersionModel, remoteVersion models.AppVersionModel) string {
	// Holds data to return
	var data = make(map[string]interface{})

	// Adds versions information to return
	data["current"] = currentVersion
	data["remote"] = remoteVersion

	// Marshal return data to json
	jsonData, err := json.Marshal(data)
	misc.Check(err)

	return string(jsonData)
}
