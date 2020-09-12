$(document).ready(function() {
    cloneTemplate("#monsterInputTmpl", "#monstersForm .fields")

    // Declared in go
    getPartyData().then((data) => {
        let jsonData = JSON.parse(data)

        $(".dataSync[data-cat=\"party\"]").each(function() {
            let field = $(this).data("field")

            if(field.includes(".")) {
                let arr = field.split("."),
                    actualField = jsonData

                for(let i = 0; i < arr.length; i++) {
                    actualField = actualField[arr[i]]
                }

                $(this).html(actualField)
            } else {
                $(this).html(jsonData[field])
            }
        })
    })

    $(".addMonster").click(() => {
        cloneTemplate("#monsterInputTmpl", "#monstersForm .fields")
    })

    $("#content").on("click", ".monsterDelete", function() {
        let count = $(".tmpl").length

        if(count > 2) {
            $(this).closest(".tmpl").remove()
            monstersUpdated()
        }
    })

    $(".fields").on("change", "input", monstersUpdated)

    $('.editParty').click(() => {
        // Declared in go
        editParty()
    })
})

function monstersUpdated() {
    // Collect data
    let monsters = []

    $(".fields").find(".monsterInputContainer").each(function() {
        let monster = {
            monster_name: null,
            monster_cr: null,
            monsters_amount: null,
            count_in_cr_mod: null
        }

        monster.monsters_amount = parseInt($(this).children(".monsterAmount").val())
        monster.monster_name = $(this).children(".monsterName").val()
        monster.monster_cr = parseFloat($(this).children(".monsterCR").val())
        monster.count_in_cr_mod = $(this).children(".monsterDifficult").is(":checked")

        monsters.push(monster)
    })

    // Declared in go
    calculateResults(monsters).then((ret) => {
        let jsonData = JSON.parse(ret)

        $('#results').removeClass('empty')

        $(".dataSync[data-cat=\"results\"]").each(function() {
            let field = $(this).data("field")

            if(field.includes(".")) {
                let arr = field.split("."),
                    actualField = jsonData

                for(let i = 0; i < arr.length; i++) {
                    actualField = actualField[arr[i]]
                }

                $(this).html(actualField)
            } else {
                $(this).html(jsonData[field])
            }
        })

        $('.encounterDifficultyResult')
            .removeClass('trivial')
            .removeClass('easy')
            .removeClass('medium')
            .removeClass('hard')
            .removeClass('deadly')
            .addClass(jsonData.encounter_difficulty)
    })
}
