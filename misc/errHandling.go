/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package misc

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/models"
)

/*
This file declares functions related with error handling
*/

//Check checks if error exists, and runs an Error Dialog
func Check(err error) {
	if err != nil {
		r := ThrowErrorGo(models.ErrorModel{
			ErrorNumber:      1000,
			ErrorDescription: err.Error(),
			ErrorType:        3,
		})
		if r == 1 {
			panic(err)
		}
	}
}

//ThrowError runs an Error Dialog (misc/errorController.go) with error model provided as json
func ThrowError(modelString json.RawMessage) int {
	var model models.ErrorModel

	// Unmarshal json data to go Error Model
	json.Unmarshal(modelString, &model)

	return ThrowErrorGo(model)
}

//ThrowErrorGO runs an Error Dialog (misc/errorController.go) with error model provided as go model
func ThrowErrorGo(model models.ErrorModel) int {
	// Create channel, that reads dialog output
	ch := make(chan int)
	// Open dialog
	go ErrorWindow(ch, model)

	// Return dialog output
	return <-ch
}
