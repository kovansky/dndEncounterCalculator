package webapp

import (
	"net/http"
)

func update(w http.ResponseWriter, r *http.Request) {
	path := "./webapp/public/html/update.html"

	http.ServeFile(w, r, path)
}
