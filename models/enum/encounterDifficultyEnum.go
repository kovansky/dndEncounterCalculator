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
