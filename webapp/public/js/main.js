/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

/*
 * This is script file (local controller) for Main View
 */

$(document).ready(function() {
    // Clone one monster input template
    cloneTemplate("#monsterInputTmpl", "#monstersForm .fields")

    // Get party data from previous view (party view); provided by Go backend
    // Declared in go
    getPartyData().then((data) => {
        // Parse string into js object
        let jsonData = JSON.parse(data)

        // Iterate through .dataSync objects in DOM to fill them with values
        $(".dataSync[data-cat=\"party\"]").each(function() {
            // Get the object field (requested data)
            let field = $(this).data("field")

            // If requested field contains dot (has more than one level), split it
            if(field.includes(".")) {
                let arr = field.split("."),
                    actualField = jsonData

                for(let i = 0; i < arr.length; i++) {
                    actualField = actualField[arr[i]]
                }

                // Set object html to the requested field value
                $(this).html(actualField)
            } else {
                // Set object html to the requested field value
                $(this).html(jsonData[field])
            }
        })
    })

    // If this isn't window first run, restore previous window state (monsters form state)
    // Declared in go
    loadWindowState().then((ret) => {
        // Check, if return isn't empty. If empty, it is the first run
        if(ret !== "") {
            // Parse string into js object
            let jsonData = JSON.parse(ret)

            // For each monster existing in js object...
            jsonData.forEach((el, index) => { // index - index (int), el - monster (object)
                // ...select last (empty) input...
                let domElement = $(".fields").find(".monsterInputContainer").last()

                // ...and fill it with values...
                domElement.find(".monsterAmount").val(el.monsters_amount)
                domElement.find(".monsterName").val(el.monster_name)
                domElement.find(".monsterCR").val(el.monster_cr)
                domElement.find(".monsterDifficult").prop("checked", el.count_in_cr_mod)

                // ...then clone another input template, if necessary
                if(index !== jsonData.length - 1) {
                    cloneTemplate("#monsterInputTmpl", "#monstersForm .fields")
                }
            })

            // Invoke function to run calculations and return encoutner difficulty
            monstersUpdated()
        }
    })

    // Add monster button click event handler
    $(".addMonster").click(() => {
        // Add empty monster input to view
        cloneTemplate("#monsterInputTmpl", "#monstersForm .fields")
    })

    // Delete monster button click event handler
    $("#content").on("click", ".monsterDelete", function() {
        let count = $(".tmpl").length

        // Check, if it isn't the only one input. If isn't, remove the requested input
        if(count > 2) {
            $(this).closest(".tmpl").remove()

            // Invoke function to run calculations and return encoutner difficulty
            monstersUpdated()
        }
    })

    // Add change event handler, when any monster name/value/CR gets updated, request the calculations
    $(".fields").on("change", "input", monstersUpdated)

    // Edit party button click event handler
    $(".editParty").click(() => {
        // Open previous view - Party View
        // Declared in go
        editParty()
    })
})

// Communicates with Go backend to run calculations and return encoutner difficulty and other related data
function monstersUpdated() {
    // Collect data from form
    let monsters = []

    // Iterate through form
    $(".fields").find(".monsterInputContainer").each(function() {
        // Init empty object
        let monster = {
            monster_name: null,
            monster_cr: null,
            monsters_amount: null,
            count_in_cr_mod: null
        }

        // Read values from inputs into js object
        monster.monsters_amount = parseInt($(this).children(".monsterAmount").val())
        monster.monster_name = $(this).children(".monsterName").val()
        monster.monster_cr = parseFloat($(this).children(".monsterCR").val())
        monster.count_in_cr_mod = $(this).children(".monsterDifficult").is(":checked")

        // Check, if monster name is empty
        if(monster.monster_name.length === 0) {
            $(this).children(".monsterName").addClass("error")
        } else {
            $(this).children(".monsterName").removeClass("error")
        }

        // Add monster object to array
        monsters.push(monster)
    })

    // Make call to Go backend, to calculate results using passed monsters array
    // Declared in go
    calculateResults(monsters).then((ret) => {
        // Parse string into js object
        let jsonData = JSON.parse(ret)

        if(jsonData.monsters_group_type === "error") {
            $("#results").addClass("empty")
        } else {
            // Show results block in view
            $("#results").removeClass("empty")

            // Iterate through .dataSync objects in DOM to fill them with values
            $(".dataSync[data-cat=\"results\"]").each(function() {
                // Get the object field (requested data)
                let field = $(this).data("field")

                // If requested field contains dot (has more than one level), split it
                if(field.includes(".")) {
                    let arr = field.split("."),
                        actualField = jsonData

                    for(let i = 0; i < arr.length; i++) {
                        actualField = actualField[arr[i]]
                    }

                    // Set object html to the requested field value
                    $(this).html(actualField)
                } else {
                    // Set object html to the requested field value
                    $(this).html(jsonData[field])
                }
            })

            // Set the final difficulty class (color)
            $(".encounterDifficultyResult")
                .removeClass("trivial")
                .removeClass("easy")
                .removeClass("medium")
                .removeClass("hard")
                .removeClass("deadly")
                .addClass(jsonData.encounter_difficulty)
        }
    })
}
