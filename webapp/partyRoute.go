package webapp

import (
	"net/http"
)

func party(w http.ResponseWriter, r *http.Request) {
	path := "./webapp/public/html/party.html"

	http.ServeFile(w, r, path)
}
