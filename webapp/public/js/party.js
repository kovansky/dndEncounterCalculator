$(document).ready(function() {
    cloneTemplate("#charInputTmpl", "#content")

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
            char.player_level = $(this).children()[1].value

            characters.push(char)
        })

        // Declared in Go
        readParty(characters).then((ret) => {
            if((ret !== 0) || (ret === 0 && characters.length !== 0)) {
                // Declared in go
                nextWindow()
            }
        })

        // ToDo: invoke go function
    })
})
