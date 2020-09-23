package webapp

import (
	"net/http"
)

//update is Update View webserver handler
func update(w http.ResponseWriter, r *http.Request) {
	path := "./webapp/public/html/update.html"

	http.ServeFile(w, r, path)
}
