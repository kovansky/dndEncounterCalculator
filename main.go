package main

import (
	"github.com/webview/webview"
	"os"
	"path/filepath"
)

var dir string

func main() {
	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	wv := webview.New(true)
	defer wv.Destroy()

	wv.SetTitle("Create your party")
	wv.SetSize(500, 400, webview.HintFixed)
	wv.Navigate("http://google.com")

	wv.Run()
}
