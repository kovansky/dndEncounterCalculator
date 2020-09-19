package models

import (
	"bufio"
	"fmt"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
	"net/http"
	"regexp"
	"strconv"
)

type AppVersionModel struct {
	CodeName         string              `json:"code_name"`
	Major            int                 `json:"major"`
	Minor            int                 `json:"minor"`
	Patch            int                 `json:"patch"`
	Channel          enum.VersionChannel `json:"channel"`
	UpdateExists     bool                `json:"update_exists"`
	CheckedForUpdate bool                `json:"checked_for_update"`
}

func GetAppVersion() *AppVersionModel {
	major, minor, patch, channel := versionFromString(constants.APP_VERSION_CURRENT)

	return &AppVersionModel{
		CodeName:         constants.APP_VERSION_CODENAME,
		Major:            major,
		Minor:            minor,
		Patch:            patch,
		Channel:          channel,
		UpdateExists:     false,
		CheckedForUpdate: false,
	}
}

func versionFromString(stringedVersion string) (int, int, int, enum.VersionChannel) {
	re := regexp.MustCompile("[.-]+")

	splitted := re.Split(stringedVersion, -1)

	if len(splitted) == 3 {
		major, _ := strconv.Atoi(splitted[0])
		minor, _ := strconv.Atoi(splitted[1])
		patch, _ := strconv.Atoi(splitted[2])

		return major, minor, patch, enum.VersionStable
	} else if len(splitted) == 4 {
		major, _ := strconv.Atoi(splitted[0])
		minor, _ := strconv.Atoi(splitted[1])
		patch, _ := strconv.Atoi(splitted[2])
		ch := enum.VersionChannelByString(splitted[3])

		return major, minor, patch, ch
	} else {
		return 0, 0, 0, enum.VersionStable
	}
}

func (avm *AppVersionModel) CheckForUpdates() (bool, int, int, int, enum.VersionChannel) {
	url := enum.GetVersionCheckUrlByChannel(avm.Channel)

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	scanner.Scan()
	remoteVerStr := scanner.Text()

	rMajor, rMinor, rPatch, rChannel := versionFromString(remoteVerStr)

	if (rMajor > avm.Major || rMinor > avm.Minor || rPatch > avm.Patch) && rChannel == avm.Channel {
		avm.UpdateExists = true
		avm.CheckedForUpdate = true
		return true, rMajor, rMinor, rPatch, rChannel
	} else {
		avm.UpdateExists = false
		avm.CheckedForUpdate = true
		return false, 0, 0, 0, ""
	}
}

func (avm AppVersionModel) ToString() string {
	return fmt.Sprintf("%d.%d.%d-%s", avm.Major, avm.Minor, avm.Patch, avm.Channel)
}
