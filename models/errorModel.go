package models

import "github.com/kovansky/dndEncounterCalculator/models/enum"

type ErrorModel struct {
	ErrorNumber      int
	ErrorDescription string
	ErrorType        enum.ErrorType
}

// ToDo: thorw method
