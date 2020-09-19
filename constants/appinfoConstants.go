package constants

const (
	// WebApp data
	APP_WEBAPP_PORT = 12354
	APP_WEBAPP_HOST = "127.0.0.1"
	APP_WEBAPP_URL  = APP_WEBAPP_HOST + ":" + string(rune(APP_WEBAPP_PORT))

	// Version data
	APP_VERSION_CHECK_URL_BETA   = "https://github.com/kovansky/dndEncounterCalculator/blob/develop/versioning/latest-beta.txt"
	APP_VERSION_CHECK_URL_STABLE = "https://github.com/kovansky/dndEncounterCalculator/blob/master/versioning/latest-stable.txt"
	APP_VERSION_CURRENT          = "0.1.0-DEV"
	APP_VERSION_CODENAME         = "YMIR"

	// Application data
	APP_GITHUB_REPO = "https://github.com/kovansky/dndEncounterCalculator"
	APP_AUTHOR      = "F4 Developer (Stanisław Kowański)"
	APP_AUTHOR_URL  = "https://f4dev.me"
)
