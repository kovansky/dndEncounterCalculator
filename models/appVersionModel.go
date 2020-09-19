package models

import (
	"bufio"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/kovansky/dndEncounterCalculator/models/enum"
	"net/http"
	"regexp"
	"strconv"
)

type AppVersionModel struct {
	CodeName     string
	Major        int
	Minor        int
	Path         int
	Channel      enum.VersionChannel
	UpdateExists bool
}

func GetAppVersion() *AppVersionModel {
	major, minor, path, channel := versionFromString(constants.APP_VERSION_CURRENT)

	return &AppVersionModel{
		CodeName:     constants.APP_VERSION_CODENAME,
		Major:        major,
		Minor:        minor,
		Path:         path,
		Channel:      channel,
		UpdateExists: nil,
	}
}

func versionFromString(stringedVersion string) (int, int, int, enum.VersionChannel) {
	re := regexp.MustCompile("[.-]+")

	splitted := re.Split(constants.APP_VERSION_CURRENT, -1)

	if len(splitted) == 3 {
		major, _ := strconv.Atoi(splitted[0])
		minor, _ := strconv.Atoi(splitted[1])
		path, _ := strconv.Atoi(splitted[2])

		return major, minor, path, enum.VersionStable
	} else {
		major, _ := strconv.Atoi(splitted[0])
		minor, _ := strconv.Atoi(splitted[1])
		path, _ := strconv.Atoi(splitted[2])
		ch := enum.VersionChannelByString(splitted[3])

		return major, minor, path, ch
	}
}

func (avm *AppVersionModel) CheckForUpdates() (bool, int, int, int, enum.VersionChannel) {
	url := enum.GetVersionCheckUrlByChannel(avm.Channel)

	resp, err := http.Get(url)
	misc.Check(err)
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	scanner.Scan()
	remoteVerStr := scanner.Text()

	rMajor, rMinor, rPath, rChannel := versionFromString(remoteVerStr)

	if (rMajor > avm.Major || rMinor > avm.Minor || rPath > avm.Path) && rChannel == avm.Channel {
		avm.UpdateExists = true
		return true, rMajor, rMinor, rPath, rChannel
	} else {
		avm.UpdateExists = false
		return false, 0, 0, 0, ""
	}
}
