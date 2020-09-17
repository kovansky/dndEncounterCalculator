package misc

import (
	"fmt"
	"github.com/kovansky/dndEncounterCalculator/models"
	"github.com/webview/webview"
)

func ErrorWindow(ch chan int, model models.ErrorModel) {
	ew := webview.New(false)
	defer ew.Destroy()

	ew.SetTitle("Error") // language
	ew.SetSize(600, 200, webview.HintFixed)

	err := ew.Bind("retValue", func(status int) int {
		ch <- status
		ew.Terminate()
		return status
	})
	Check(err)

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
</html>`, "http://127.0.0.1:12360", model.ErrorNumber, model.ErrorDescription, "http://127.0.0.1:12360"))

	ew.Run()
}
