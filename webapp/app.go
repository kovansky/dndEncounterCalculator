/*
Package webapp holds Web Application code - the server listener and routes handlers
*/
package webapp

import (
	"fmt"
	"github.com/kovansky/dndEncounterCalculator/misc"
	"net/http"
)

//App runs webserver that holds views for application
func App() {
	// Create new Mux
	mux := http.NewServeMux()
	// Register handlers
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./webapp/public/")))) // Static resources (images, stylesheets, script files)

	mux.HandleFunc("/party", party)   // Party view
	mux.HandleFunc("/main", main)     // Main view
	mux.HandleFunc("/update", update) // Update view

	// Create server configuration
	server := &http.Server{
		// ToDo: configurable port, saved as const
		Addr:    "127.0.0.1:12354",
		Handler: mux,
	}

	fmt.Println(server.Addr)

	// Run webserver
	err := server.ListenAndServe()
	misc.Check(err)
}
