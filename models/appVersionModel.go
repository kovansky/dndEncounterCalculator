/*
Package models implements application data models, used to communicate with JavaScript, hold data and write it to the disk.
*/
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

//GetAppVersion returns AppVersionModel from current application version (declared in appInfoConstants.go)
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

//versionFromString splits version string ("MAJOR.MINOR.PATCH-CHANNEL") into separate variables
func versionFromString(stringedVersion string) (int, int, int, enum.VersionChannel) {
	// Create regular expression to split string by dots and hypens
	re := regexp.MustCompile("[.-]+")

	// Split the version string by regular expression into slice
	splitted := re.Split(stringedVersion, -1)

	if len(splitted) == 3 {
		// If version string doesn't contain channel, it is stable
		major, _ := strconv.Atoi(splitted[0])
		minor, _ := strconv.Atoi(splitted[1])
		patch, _ := strconv.Atoi(splitted[2])

		return major, minor, patch, enum.VersionStable
	} else if len(splitted) == 4 {
		// If version string contains channel
		major, _ := strconv.Atoi(splitted[0])
		minor, _ := strconv.Atoi(splitted[1])
		patch, _ := strconv.Atoi(splitted[2])
		ch := enum.VersionChannelByString(splitted[3])

		return major, minor, patch, ch
	} else {
		// If version string is wrong
		return 0, 0, 0, enum.VersionStable
	}
}

//CheckForUpdates compares version with the latest version from version check URL of corresponding channel
func (avm *AppVersionModel) CheckForUpdates() (bool, int, int, int, enum.VersionChannel) {
	// Get version check URL of channel
	url := enum.GetVersionCheckUrlByChannel(avm.Channel)

	// Opens URL, defers closing connection
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	// Creates scanner to read the URL body
	scanner := bufio.NewScanner(resp.Body)

	// Reads first line
	scanner.Scan()
	remoteVerStr := scanner.Text()

	// Split latest (remote) version string into separate variables
	rMajor, rMinor, rPatch, rChannel := versionFromString(remoteVerStr)

	// Compare remote major, minor and patch with current ones
	if (rMajor > avm.Major || rMinor > avm.Minor || rPatch > avm.Patch) && rChannel == avm.Channel {
		// Update exists
		avm.UpdateExists = true
		avm.CheckedForUpdate = true
		return true, rMajor, rMinor, rPatch, rChannel
	} else {
		// This is latest version
		avm.UpdateExists = false
		avm.CheckedForUpdate = true
		return false, 0, 0, 0, ""
	}
}

//String returns version string ("MAJOR.MINOR.PATCH-CHANNEL")
func (avm AppVersionModel) String() string {
	return fmt.Sprintf("%d.%d.%d-%s", avm.Major, avm.Minor, avm.Patch, avm.Channel)
}
