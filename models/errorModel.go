package models

import "github.com/kovansky/dndEncounterCalculator/models/enum"

type ErrorModel struct {
	ErrorNumber      int            `json:"error_number"`
	ErrorDescription string         `json:"error_description"`
	ErrorType        enum.ErrorType `json:"error_type"`
}
