$(document).ready(function() {
    cloneTemplate("#charInputTmpl", "#content")

    // Declared in go
    loadWindowState().then((ret) => {
        let jsonData = JSON.parse(ret)

        let partySelect = $('#savedPartySelect')
        $.each(jsonData.partiesSelect, (index, el) => {
            partySelect.append($('<option>', {
                value: index,
                text: el
            }))
        })

        if(jsonData.party !== "") {
            jsonData.party.forEach((el, index) => {
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

        // Declared in go
        loadParty(id).then((ret) => {
            let jsonData = JSON.parse(ret)

            if(jsonData === "-2006") {
                lockWindow()
                // Declared in go
                runError({
                    error_number: 2006,
                    error_description: "Party with id " + id + " does not exist in saved parties file.",
                    error_type: 2
                }).then((errRet) => {
                    if(errRet === 1) {
                        unlockWindow()
                    }
                })
            } else {
                $('.partyName').val(jsonData.party_name)

                $('#content').html('')

                $.each(jsonData.party_players, (name, value) => {
                    cloneTemplate("#charInputTmpl", "#content")

                    let domElement = $("#content").find(".charInputContainer").last()

                    domElement.find(".charName").val(name)
                    domElement.find(".charLevel").val(value.player_level)
                })
            }
        })

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
