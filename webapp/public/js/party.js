/*
 * This is script file (local controller) for Party View
 */

$(document).ready(function() {
    // Clone one player input template
    cloneTemplate("#charInputTmpl", "#content")

    // If this isn't window first run, restore previous window state (players form state, party data)
    // Declared in go
    loadWindowState().then((ret) => {
        // Parse string into js object
        let jsonData = JSON.parse(ret)

        // Add avaliable saved parties to select box
        let partySelect = $("#savedPartySelect")
        $.each(jsonData.partiesSelect, (index, el) => { // index - id, el - name
            partySelect.append($("<option>", {
                value: index,
                text: el
            }))
        })

        // Check, if party data isn't empty. If empty, it is the first run
        if(jsonData.party !== "") {
            // For each player existing in js object...
            jsonData.party.forEach((el, index) => { // index - index (int), el - player (object)
                // ...select last (empty) input...
                let domElement = $("#content").find(".charInputContainer").last()

                // ...and fill it with values...
                domElement.find(".charName").val(el.player_name)
                domElement.find(".charLevel").val(el.player_level)

                // ...then clone another input template, if necessary
                if(index !== jsonData.party.length - 1) {
                    cloneTemplate("#charInputTmpl", "#content")
                }
            })
        }
    })

    // Add player button click event handler
    $(".charNew").click(() => {
        // Add empty player input to view
        cloneTemplate("#charInputTmpl", "#content")
    })

    // Delete player button click event handler
    $("#content").on("click", ".charDelete", function() {
        let count = $(".tmpl").length

        // Check, if it isn't the only one input. If isn't, remove the requested input
        if(count > 2) {
            $(this).closest(".tmpl").remove()
        }
    })

    // Go to the next window (Main View)
    $(".nextStep").click(function() {
        let characters = []

        // Collect party data (party players)
        $("#content").find(".charInputContainer").each(function() {
            // Init empty player object
            let char = {
                player_name: null,
                player_level: null
            }

            // Read values from input into js object
            char.player_name = $(this).children()[0].value
            char.player_level = parseInt($(this).children()[1].value)

            // Add player to characters array
            characters.push(char)
        })

        // Pass data to Go backend, to save in RAM
        // Declared in go
        readParty(characters).then((ret) => {
            // If return is correct (ok, no errors), go to next window (Main View). Else go to error return
            if(((ret !== 0) || (ret === 0 && characters.length !== 0)) && (ret > -2000)) {
                // Declared in go
                nextWindow()
            } else {
                let error = null

                // Declare an error according to the error number
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

                // Lock window to block interactions while dialog is open
                lockWindow()
                // Open error dialog
                // Declared in go
                runError(error)
                    .then((ret) => {
                        if(ret === 1) {
                            // Unlock window lock on dialog close
                            unlockWindow()
                        }
                    })
            }
        })
    })

    // On party select from select element
    $("#savedPartySelect").change(function() {
        // Get party ID
        let id = $(this).val()

        // Request the backend to return the party data corresponding to the party ID
        // Declared in go
        loadParty(id).then((ret) => {
            // Parse string into js object
            let jsonData = JSON.parse(ret)

            // Check, if there is an error
            if(jsonData === "-2006") {
                // Lock window to block interactions while dialog is open
                lockWindow()
                // Open error dialog
                // Declared in go
                runError({
                    error_number: 2006,
                    error_description: "Party with id " + id + " does not exist in saved parties file.",
                    error_type: 2
                }).then((errRet) => {
                    if(errRet === 1) {
                        // Unlock window lock on dialog close
                        unlockWindow()
                    }
                })
            } else {
                // If no error detected, proceed
                // Fill party name text input with choosen party name
                $(".partyName").val(jsonData.party_name)

                // Delete existing players
                $("#content").html("")

                // For each player existing in js object...
                $.each(jsonData.party_players, (name, value) => { // name - player name (string), value - player object
                    // ...clone input for that player...
                    cloneTemplate("#charInputTmpl", "#content")

                    // ...choose that input...
                    let domElement = $("#content").find(".charInputContainer").last()

                    // ...and fill it with values.
                    domElement.find(".charName").val(name)
                    domElement.find(".charLevel").val(value.player_level)
                })
            }
        })
    })

    // Deleting parties from saved parties
    $(".partyDelete").click(() => {
        // Setup variables
        let container = $("#partyDeleteConfirmation"), // Confirmation container (dialog)
            select = $("#savedPartySelect"), // Parties select box
            selectText = $("#savedPartySelect option:selected").text() // Selected party name

        // Check, if the ID of selected option isn't null. If null, it is placeholder (disabled option - no option was choosen)
        if(select.val() !== null) {
            // Show and fill confirmation dialog
            container.addClass("visible")
            container.find(".dataSync[data-field=\"party_name\"]").html(selectText).data("party_id", select.val())
        } else {
            // Lock window to block interactions while dialog is open
            lockWindow()
            // Open error dialog
            // Declared in go
            runError({
                error_number: 2007,
                error_description: "You have to select the party you want to delete.",
                error_type: 1
            })
                .then((ret) => {
                    if(ret === 1) {
                        // Unlock window lock on dialog close
                        unlockWindow()
                    }
                })
        }
    })

    // Add handler for party delete declining button (in dialog)
    $(".partyDeleteDecline").click(() => {
        let container = $("#partyDeleteConfirmation") // Dialog

        // Hide dialog
        container.removeClass("visible")
    })

    // Add handler for party delete confirmation button (in dialog)
    $(".partyDeleteConfirm").click(() => {
        // Setup variables
        let container = $("#partyDeleteConfirmation"), // Dialog
            party_id = container.find(".dataSync[data-field=\"party_name\"]").data("party_id") // Desired party ID

        // Request backend to remove party with given id from disk
        // Declared in go
        removeParty(party_id).then((ret) => {
            // Check return for errors
            if(ret === 1) {
                // Operation successful
                // Remove option from select box
                $("#savedPartySelect option[value=\"" + party_id + "\"]").remove()

                // Hide dialog
                container.removeClass("visible")
            } else if(ret === -2006) {
                // Lock window to block interactions while dialog is open
                lockWindow()
                // Open error dialog
                // Declared in go
                runError({
                    error_number: 2006,
                    error_description: "Party with id " + party_id + " does not exist in saved parties file.",
                    error_type: 2
                }).then((errRet) => {
                    if(errRet === 1) {
                        // Hide dialog
                        container.removeClass("visible")

                        // Unlock window lock on dialog close
                        unlockWindow()
                    }
                })
            }
        })
    })

    // Add handler for party saving
    $(".saveParty").click(() => {
        // Empty objects for data holding
        let party = {},
            characters = {}

        // Iterate through inputs (players) to add them to the object
        $("#content").find(".charInputContainer").each(function() {
            // Create empty character object
            let char = {
                player_name: null,
                player_level: null
            }

            // Fill it with value
            char.player_name = $(this).children()[0].value
            char.player_level = parseInt($(this).children()[1].value)

            // Add to characters object
            characters[char.player_name] = char
        })

        // Add characters object to the party object
        party.party_players = characters

        // Get party name (from input)
        let partyName = $(".partyName").val()

        // Validate party name
        if(partyName === "") {
            // Lock window to block interactions while dialog is open
            lockWindow()
            // Open error dialog
            runError({
                error_number: 2005,
                error_description: "Party name cannot be empty",
                error_type: 1
            })
                .then((ret) => {
                    if(ret === 1) {
                        // Unlock window lock on dialog close
                        unlockWindow()
                    }
                })
        } else {
            // Save party name to party object
            party.party_name = partyName
            // Generate party ID from name and save it to object
            party.party_id = idFromString(partyName)

            // Request backend to save the given party to disk
            // Declared in go
            writeParty(party).then((ret) => {
                let error = null

                // Check for errors
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

                // Check, if error occurred
                if(error != null) {
                    // Lock window to block interactions while dialog is open
                    lockWindow()
                    // Open error dialog
                    // Declared in go
                    runError(error)
                        .then((ret) => {
                            if(ret === 1) {
                                // Unlock window lock on dialog close
                                unlockWindow()
                            }
                        })
                } else if(ret === 0) {
                    // Operation successful
                    // Get party select box from DOM
                    let partySelect = $("#savedPartySelect")

                    // Append a new option to the select
                    // FixMe: if already exists, DO NOT ADD
                    partySelect.append($("<option>", {
                        value: party.party_id,
                        text: party.party_name
                    }))
                }
            })
        }
    })
})
