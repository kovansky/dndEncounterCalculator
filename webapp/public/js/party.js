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
                name: null,
                level: null
            }

            char.name = $(this).children()[0].value
            char.level = $(this).children()[1].value

            characters.push(char)
        })

        console.log(characters)

        // ToDo: invoke go function
    })
})
