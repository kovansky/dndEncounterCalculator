/*
Package webapp holds Web Application code - the server listener and routes handlers
*/
package webapp

import (
	"github.com/gorilla/mux"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"net/http"
)

//App runs webserver that holds views for application
func App() {
	// Create new Mux
	router := mux.NewRouter()
	// Register handlers
	router.HandleFunc("/party", party)   // Party view
	router.HandleFunc("/main", main)     // Main view
	router.HandleFunc("/update", update) // Update view
	router.HandleFunc("/", party)

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./webapp/public/")))) // Static resources (images, stylesheets, script files)

	// Run webserver
	// ToDo: changeable addr
	err := http.ListenAndServe("127.0.0.1:12356", router)
	misc.Check(err)
}
