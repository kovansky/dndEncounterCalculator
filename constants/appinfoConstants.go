/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

/*
Package constants holds all constant information used in application
*/
package constants

/*
This file creates and sets values of application information constants, like app version data or application author.
*/

const (
	// WebApp data

	APP_WEBAPP_PORT = "12356"
	APP_WEBAPP_HOST = "127.0.0.1"
	APP_WEBAPP_URL  = APP_WEBAPP_HOST + ":" + APP_WEBAPP_PORT

	// Version data

	APP_VERSION_CHECK_URL_BETA   = "https://github.com/kovansky/dndEncounterCalculator/raw/develop/versioning/latest-beta.txt"
	APP_VERSION_CHECK_URL_STABLE = "https://github.com/kovansky/dndEncounterCalculator/raw/develop/versioning/latest-stable.txt"
	APP_VERSION_CURRENT          = "1.0.0-STABLE"
	APP_VERSION_CODENAME         = "YMIR"
	APP_UPDATE_URL               = "https://github.com/kovansky/dndEncounterCalculator/releases/tag/v%s"

	// Application data

	APP_GITHUB_REPO = "https://github.com/kovansky/dndEncounterCalculator"
	APP_AUTHOR      = "F4 Developer (Stanisław Kowański)"
	APP_AUTHOR_URL  = "https://f4dev.me"
)
