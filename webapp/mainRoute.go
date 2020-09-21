package webapp

import (
	"net/http"
)

//main is Main View webserver handler
func main(w http.ResponseWriter, r *http.Request) {
	path := "./webapp/public/html/main.html"

	http.ServeFile(w, r, path)
}
