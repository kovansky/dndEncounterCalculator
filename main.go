package main

import (
	"github.com/kovansky/dndEncounterCalculator/webapp"
	"github.com/webview/webview"
)

func main() {
	go webapp.App()

	wv := webview.New(true)
	defer wv.Destroy()

	wv.SetTitle("Create your party")
	wv.SetSize(500, 400, webview.HintMin)

	wv.Navigate("http://127.0.0.1:12346")

	wv.Run()
}
