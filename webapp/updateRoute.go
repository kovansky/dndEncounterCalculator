/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package webapp

import (
	"net/http"
)

//update is Update View webserver handler
func update(w http.ResponseWriter, r *http.Request) {
	path := "./webapp/public/html/update.html"

	http.ServeFile(w, r, path)
}
