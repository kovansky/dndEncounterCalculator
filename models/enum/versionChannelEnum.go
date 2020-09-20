package enum

import (
	"github.com/kovansky/dndEncounterCalculator/constants"
	"strings"
)

//VersionChannel specifies possible version channels
type VersionChannel string

const (
	VersionStable VersionChannel = "STABLE"
	VersionBeta   VersionChannel = "BETA"
	VersionDev    VersionChannel = "DEV"
)

//VersionChannelByString returns version channel by string, or Stable if string doesn't relate to any existing channel
func VersionChannelByString(str string) VersionChannel {
	switch strings.ToUpper(str) {
	case "STABLE":
		return VersionStable
	case "BETA":
		return VersionBeta
	case "DEV":
		return VersionDev
	default:
		return VersionStable
	}
}

//GetVersionCheckUrlByChannel returns an version check URL corresponding to the version channel
func GetVersionCheckUrlByChannel(ch VersionChannel) string {
	switch ch {
	case VersionStable:
		return constants.APP_VERSION_CHECK_URL_STABLE
	case VersionBeta:
		return constants.APP_VERSION_CHECK_URL_BETA
	case VersionDev:
		return constants.APP_VERSION_CHECK_URL_BETA
	}

	return constants.APP_VERSION_CHECK_URL_STABLE
}
