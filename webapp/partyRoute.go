package webapp

import (
	"net/http"
)

//party is Party View webserver handler
func party(w http.ResponseWriter, r *http.Request) {
	path := "./webapp/public/html/party.html"

	http.ServeFile(w, r, path)
}
