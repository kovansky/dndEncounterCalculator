/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

/*
Package enum is holding all enumerated constants used in application
*/
package enum

//EncounterDifficulty specifies possible encounter difficulty levels
type EncounterDifficulty string

const (
	EncounterTrivial EncounterDifficulty = "trivial"
	EncounterEasy    EncounterDifficulty = "easy"
	EncounterMedium  EncounterDifficulty = "medium"
	EncounterHard    EncounterDifficulty = "hard"
	EncounterDeadly  EncounterDifficulty = "deadly"
)

//CalculateEncounterDifficulty compares adjustedXP to party thresholds and specifies the encounter difficulty
func CalculateEncounterDifficulty(thresholds map[string]int, adjustedXP float32) EncounterDifficulty {
	if adjustedXP < float32(thresholds["easy"]) {
		return EncounterTrivial
	} else if adjustedXP >= float32(thresholds["easy"]) && adjustedXP < float32(thresholds["medium"]) {
		return EncounterEasy
	} else if adjustedXP >= float32(thresholds["medium"]) && adjustedXP < float32(thresholds["hard"]) {
		return EncounterMedium
	} else if adjustedXP >= float32(thresholds["hard"]) && adjustedXP < float32(thresholds["deadly"]) {
		return EncounterHard
	} else {
		return EncounterDeadly
	}
}
