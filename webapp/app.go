package webapp

import (
	"github.com/kovansky/dndEncounterCalculator/misc"
	"net/http"
)

func App() {
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./webapp/public/"))))
	mux.HandleFunc("/party", party)
	mux.HandleFunc("/main", main)

	server := &http.Server{
		// ToDo: configurable port, saved as const
		Addr:    "127.0.0.1:12360",
		Handler: mux,
	}

	err := server.ListenAndServe()
	misc.Check(err)
}
