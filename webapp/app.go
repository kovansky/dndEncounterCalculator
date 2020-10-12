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
	mux := mux.NewRouter()
	// Register handlers
	mux.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./webapp/public/")))) // Static resources (images, stylesheets, script files)

	mux.HandleFunc("/", party)
	mux.HandleFunc("/party", party)   // Party view
	mux.HandleFunc("/main", main)     // Main view
	mux.HandleFunc("/update", update) // Update view

	// Run webserver
	// ToDo: changeable addr
	err := http.ListenAndServe("127.0.0.1:12354", mux)
	misc.Check(err)
}
