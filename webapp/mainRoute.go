/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package webapp

import (
	"net/http"
)

//main is Main View webserver handler
func main(w http.ResponseWriter, r *http.Request) {
	path := "./webapp/public/html/main.html"

	http.ServeFile(w, r, path)
}
