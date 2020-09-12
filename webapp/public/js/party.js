$(document).ready(function() {
    cloneTemplate("#charInputTmpl", "#content")

    // Declared in go
    loadWindowState().then((ret) => {
        if(ret !== "") {
            let jsonData = JSON.parse(ret)

            jsonData.forEach((el, index) => {
                let domElement = $("#content").find(".charInputContainer").last()

                domElement.find(".charName").val(el.player_name)
                domElement.find(".charLevel").val(el.player_level)

                if(index !== jsonData.length - 1) {
                    cloneTemplate("#charInputTmpl", "#content")
                }
            })
        }
    })

    $(".charNew").click(() => {
        cloneTemplate("#charInputTmpl", "#content")
    })

    $("#content").on("click", ".charDelete", function() {
        let count = $(".tmpl").length

        if(count > 2) {
            $(this).closest(".tmpl").remove()
        }
    })

    $(".nextStep").click(function() {
        let characters = []

        $("#content").find(".charInputContainer").each(function() {
            let char = {
                player_name: null,
                player_level: null
            }

            char.player_name = $(this).children()[0].value
            char.player_level = parseInt($(this).children()[1].value)

            characters.push(char)
        })

        // Declared in go
        readParty(characters).then((ret) => {
            if((ret !== 0) || (ret === 0 && characters.length !== 0)) {
                // Declared in go
                nextWindow()
            }
        })
    })
})
