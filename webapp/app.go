/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

/*
Package webapp holds Web Application code - the server listener and routes handlers
*/
package webapp

import (
	"github.com/gorilla/mux"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"github.com/markbates/pkger"
	"net/http"
)

//App runs webserver that holds views for application
func App() {
	// Create new router
	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(pkger.Dir("/webapp/public")))) // Static resources (images, stylesheets, script files)

	// Run webserver
	err := http.ListenAndServe(constants.APP_WEBAPP_URL, router)
	misc.Check(err)
}
