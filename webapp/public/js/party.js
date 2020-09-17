$(document).ready(function() {
    cloneTemplate("#charInputTmpl", "#content")

    // Declared in go
    // ToDo: add populating saved party select box to load window state
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
            if(((ret !== 0) || (ret === 0 && characters.length !== 0)) && (ret > -2000)) {
                // Declared in go
                nextWindow()
            } else {
                let error = null

                switch(ret) {
                    case -2001:
                        error = {
                            error_number: 2001,
                            error_description: "Party cannot be empty",
                            error_type: 1
                        }
                        break
                    case -2002:
                        error = {
                            error_number: 2002,
                            error_description: "Player level cannot be null",
                            error_type: 1
                        }
                        break
                    case -2003:
                        error = {
                            error_number: 2003,
                            error_description: "Player level cannot be less than 1",
                            error_type: 1
                        }
                        break
                    case -2004:
                        error = {
                            error_number: 2004,
                            error_description: "Player name cannot be null",
                            error_type: 1
                        }
                        break
                }

                lockWindow()
                // Declared in go
                runError(error)
                    .then((ret) => {
                        if(ret === 1) {
                            unlockWindow()
                        }
                    })
            }
        })
    })

    $('#savedPartySelect').change(function() {
        let id = $(this).val()

        // ToDo: load party by id
    })

    $('.partyDelete').click(() => {
        // ToDo: show "are you sure?" and delete by id logic
    })

    $('.saveParty').click(() => {
        let party = {},
            characters = {}

        $("#content").find(".charInputContainer").each(function() {
            let char = {
                player_name: null,
                player_level: null
            }

            char.player_name = $(this).children()[0].value
            char.player_level = parseInt($(this).children()[1].value)

            characters[char.player_name] = char
        })

        party.party_players = characters

        let partyName = $('.partyName').val()

        if(partyName === "") {
            runError({
                error_number: 2005,
                error_description: "Party name cannot be empty",
                error_type: 1
            })
                .then((ret) => {
                    if(ret === 1) {
                        unlockWindow()
                    }
                })
        } else {
            party.party_name = partyName
            party.party_id = idFromString(partyName)

            // ToDo: get return code. On 0 (success) add to listing
            // Declared in go
            writeParty(party)
        }
    })
})
