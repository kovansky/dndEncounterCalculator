/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package models

import "github.com/kovansky/dndEncounterCalculator/models/enum"

//ErrorModel is a model of error to be passed into error dialog throw
type ErrorModel struct {
	ErrorNumber      int            `json:"error_number"`
	ErrorDescription string         `json:"error_description"`
	ErrorType        enum.ErrorType `json:"error_type"`
}
