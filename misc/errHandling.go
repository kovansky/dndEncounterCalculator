package misc

import (
	"encoding/json"
	"github.com/kovansky/dndEncounterCalculator/models"
)

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

func ThrowError(modelString json.RawMessage) int {
	var model models.ErrorModel

	json.Unmarshal(modelString, &model)

	return ThrowErrorGo(model)
}

func ThrowErrorGo(model models.ErrorModel) int {
	ch := make(chan int)
	go ErrorWindow(ch, model)

	return <-ch
}
