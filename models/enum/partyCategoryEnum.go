package enum

//PartyCategory specifies possible party categories by size
type PartyCategory string

const (
	PartySmall    PartyCategory = "small"
	PartyStandard PartyCategory = "standard"
	PartyBig      PartyCategory = "big"
)

//PartyCategoryBySize returns party category by given party size
func PartyCategoryBySize(amount int) PartyCategory {
	if amount < 3 {
		return PartySmall
	} else if amount >= 3 && amount < 6 {
		return PartyStandard
	} else {
		return PartyBig
	}
}
