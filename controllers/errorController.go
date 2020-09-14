package controllers

import (
	"github.com/webview/webview"
)

func ErrorWindow(ch chan int) {
	ew := webview.New(false)
	defer func() {
		ch <- 1
	}()
	defer ew.Destroy()

	ew.SetTitle("Error") // language
	ew.SetSize(600, 200, webview.HintFixed)

	ew.Navigate("data:text/html," + `
<!doctype html>
<html>
	<body>
        <p>No i giit!</p>
	</body>
</html>`)

	ew.Run()
}
