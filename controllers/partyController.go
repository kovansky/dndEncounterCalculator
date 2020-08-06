package controllers

import "github.com/webview/webview"

func PartyWindow(wv webview.WebView) {
	wv.SetTitle("Create your party")
	wv.SetSize(600, 550, webview.HintFixed)

	wv.Navigate("http://127.0.0.1:12346/party")
}
