package enum

type PartyCategory string

const (
	PartySmall    PartyCategory = "small"
	PartyStandard PartyCategory = "standard"
	PartyBig      PartyCategory = "big"
)

func PartyCategoryBySize(amount int) PartyCategory {
	if amount < 3 {
		return PartySmall
	} else if amount >= 3 && amount < 6 {
		return PartyStandard
	} else {
		return PartyBig
	}
}
