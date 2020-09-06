package webapp

import (
	"net/http"
)

func main(w http.ResponseWriter, r *http.Request) {
	path := "./webapp/public/html/main.html"

	http.ServeFile(w, r, path)
}
