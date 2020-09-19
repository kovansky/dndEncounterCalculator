package models

import (
	"bufio"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
	"net/http"
	"regexp"
	"strconv"
)

type AppVersionModel struct {
	CodeName         string
	Major            int
	Minor            int
	Path             int
	Channel          enum.VersionChannel
	UpdateExists     bool
	CheckedForUpdate bool
}

func GetAppVersion() *AppVersionModel {
	major, minor, path, channel := versionFromString(constants.APP_VERSION_CURRENT)

	return &AppVersionModel{
		CodeName:         constants.APP_VERSION_CODENAME,
		Major:            major,
		Minor:            minor,
		Path:             path,
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
		path, _ := strconv.Atoi(splitted[2])

		return major, minor, path, enum.VersionStable
	} else if len(splitted) == 4 {
		major, _ := strconv.Atoi(splitted[0])
		minor, _ := strconv.Atoi(splitted[1])
		path, _ := strconv.Atoi(splitted[2])
		ch := enum.VersionChannelByString(splitted[3])

		return major, minor, path, ch
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

	rMajor, rMinor, rPath, rChannel := versionFromString(remoteVerStr)

	if (rMajor > avm.Major || rMinor > avm.Minor || rPath > avm.Path) && rChannel == avm.Channel {
		avm.UpdateExists = true
		avm.CheckedForUpdate = true
		return true, rMajor, rMinor, rPath, rChannel
	} else {
		avm.UpdateExists = false
		avm.CheckedForUpdate = true
		return false, 0, 0, 0, ""
	}
}
