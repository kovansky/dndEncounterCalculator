$(document).ready(function() {
    cloneTemplate("#charInputTmpl", "#content")

    // Declared in go
    loadWindowState().then((ret) => {
        let jsonData = JSON.parse(ret)

        let partySelect = $("#savedPartySelect")
        $.each(jsonData.partiesSelect, (index, el) => {
            partySelect.append($("<option>", {
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

    $("#savedPartySelect").change(function() {
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
                $(".partyName").val(jsonData.party_name)

                $("#content").html("")

                $.each(jsonData.party_players, (name, value) => {
                    cloneTemplate("#charInputTmpl", "#content")

                    let domElement = $("#content").find(".charInputContainer").last()

                    domElement.find(".charName").val(name)
                    domElement.find(".charLevel").val(value.player_level)
                })
            }
        })
    })

    $(".partyDelete").click(() => {
        let container = $("#partyDeleteConfirmation"),
            select = $("#savedPartySelect"),
            selectText = $("#savedPartySelect option:selected").text()

        if(select.val() !== null) {
            container.addClass("visible")
            container.find(".dataSync[data-field=\"party_name\"]").html(selectText).data("party_id", select.val())
        } else {
            lockWindow()
            runError({
                error_number: 2007,
                error_description: "You have to select the party you want to delete.",
                error_type: 1
            })
                .then((ret) => {
                    if(ret === 1) {
                        unlockWindow()
                    }
                })
        }
    })

    $(".partyDeleteDecline").click(() => {
        let container = $("#partyDeleteConfirmation")

        container.removeClass("visible")
    })

    $(".partyDeleteConfirm").click(() => {
        let container = $("#partyDeleteConfirmation"),
            party_id = container.find(".dataSync[data-field=\"party_name\"]").data("party_id")

        // Declared in go
        removeParty(party_id).then((ret) => {
            if(ret === 1) {
                $("#savedPartySelect option[value=\"" + party_id + "\"]").remove()

                container.removeClass("visible")
            } else if(ret === -2006) {
                lockWindow()
                // Declared in go
                runError({
                    error_number: 2006,
                    error_description: "Party with id " + party_id + " does not exist in saved parties file.",
                    error_type: 2
                }).then((errRet) => {
                    if(errRet === 1) {
                        container.removeClass("visible")

                        unlockWindow()
                    }
                })
            }
        })
    })

    $(".saveParty").click(() => {
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

        let partyName = $(".partyName").val()

        if(partyName === "") {
            lockWindow()
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

            // Declared in go
            writeParty(party).then((ret) => {
                let error = null

                switch(ret) {
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

                if(error != null) {
                    lockWindow()
                    // Declared in go
                    runError(error)
                        .then((ret) => {
                            if(ret === 1) {
                                unlockWindow()
                            }
                        })
                } else if(ret === 0) {
                    let partySelect = $("#savedPartySelect")

                    partySelect.append($("<option>", {
                        value: party.party_id,
                        text: party.party_name
                    }))
                }
            })
        }
    })
})
