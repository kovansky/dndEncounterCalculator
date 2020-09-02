package controllers

import "github.com/webview/webview"

func MainWindow(wv webview.WebView) {
	wv.SetTitle("D&D Encounter Calculator") // language
	wv.SetSize(1000, 800, webview.HintFixed)

	wv.Navigate("http://127.0.0.1:12345/main")
}
