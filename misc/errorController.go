/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

package misc

import (
	"fmt"
	"github.com/kovansky/dndEncounterCalculator/constants"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/webview/webview"
)

//ErrorWindow is a controller function of Error View (dialog). It creates a WebView window
func ErrorWindow(ch chan int, model models.ErrorModel) {
	// Create webview window, and defer destroying it
	ew := webview.New(false)
	defer ew.Destroy()

	// Adjust window data to view
	ew.SetTitle("Error") // language
	ew.SetSize(600, 200, webview.HintFixed)

	// Dialog controls (buttons) logic
	err := ew.Bind("retValue", func(code int) int {
		// Loads dialog code to return channel
		ch <- code
		// Close dialog window
		ew.Terminate()
		return code
	})
	Check(err)

	// Opens Error View in window
	ew.Navigate("data:text/html," + fmt.Sprintf(`<!doctype html>
<html>
<style>
    .btn {
        background:    #C4C4C4;
        border-radius: 7px;
        transition:    color 1s, background-color 1s;
    }

    .btn.icnBtn {
        width:         30px;
        height:        30px;
        position:      absolute;
        border-radius: 100px;
    }

    .btn:hover {
        cursor:     pointer;
        background: #A4A4A4;
    }

    .btnText {
        position:    absolute;
        top:         4px;
        right:       15px;

        font-family: "Georgia", "Calisto MT", "Palatino", sans-serif;
        font-style:  normal;
        font-weight: 300;
        font-size:   14px;
        line-height: 18px;
        text-align:  center;
    }

    .btnIcon {
        position: absolute;
        top:      4px;
        left:     8px;
        width:    17px;
        height:   17px;
    }

    #iconContainer {
        position: absolute;
        width:    170px;
        height:   200px;
        left:     0;
        top:      0;
    }

    #iconContainer img {
        position: absolute;
        width:    110px;
        height:   124px;
        left:     30px;
        top:      38px;
    }

    #errorContainer {
        position: absolute;
        width:    430px;
        height:   200px;
        left:     170px;
        top:      0;
    }

    .winTitle {
        position:    absolute;
        min-width:   244px;
        left:        93px;
        top:         15px;

        font-family: "Georgia", "Calisto MT", "Palatino", sans-serif;
        font-style:  normal;
        font-weight: normal;
        font-size:   22px;
        line-height: 27px;

        color:       #000000;
    }

    .errorNumber {
        position:    absolute;
        min-width:   211px;
        left:        20px;
        top:         55px;

        font-family: "Georgia", "Calisto MT", "Palatino", sans-serif;
        font-style:  normal;
        font-weight: normal;
        font-size:   14px;
        line-height: 17px;

        color:       #000000;
    }

    .errorDescription {
        position:    absolute;
        min-width:   330px;
        left:        50px;
        top:         80px;

        font-family: "Courier New", "Lucida Console", monospace;
        font-style:  normal;
        font-weight: 300;
        font-size:   14px;
        line-height: 16px;

        color:       #000000;
    }

    .okBtn {
        position:   absolute;
        width:      80px;
        height:     25px;
        bottom:     10px;
        right:      15px;
    }
</style>
<body>
<div id="iconContainer">
    <img src="%s/public/icons/errorIcon.svg" style="width: 112px; height: 126px" alt="">
</div>
<div id="errorContainer">
    <p class="winTitle">
        Oops! There was an error!
        <!-- language -->
    </p>

    <p class="errorNumber">
        The error has number %d and says:
        <!-- language -->
    </p>

    <p class="errorDescription">
        %s
        <!-- language -->
    </p>

    <div id="controls">
        <span class="btn okBtn">
            <img class="btnIcon" src="%s/public/icons/checkIcon.svg" alt="">
            <span class="btnText">Ok</span>
            <!-- language -->
        </span>
    </div>
</div>
</body>
<script>
    document.querySelector('.okBtn').addEventListener('click', () => {
        // Declared in go
        retValue(1)
    })
</script>
</html>`, "http://"+constants.APP_WEBAPP_URL, model.ErrorNumber, model.ErrorDescription, "http://"+constants.APP_WEBAPP_URL))

	// Runs window code
	ew.Run()
}
