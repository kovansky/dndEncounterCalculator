package enum

import "github.com/kovansky/dndEncounterCalculator/constants"

type VersionChannel string

const (
	VersionStable VersionChannel = "STABLE"
	VersionBeta   VersionChannel = "BETA"
	VersionDev    VersionChannel = "DEV"
)

func VersionChannelByString(str string) VersionChannel {
	switch str {
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

func GetVersionCheckUrlByChannel(ch VersionChannel) string {
	switch ch {
	case VersionStable:
		return constants.APP_VERSION_CHECK_URL_STABLE
	case VersionBeta:
	case VersionDev:
		return constants.APP_VERSION_CHECK_URL_BETA
	}

	return constants.APP_VERSION_CHECK_URL_STABLE
}
