$(document).ready(function() {
    cloneTemplate("#monsterInputTmpl", "#monstersForm .fields")

    // Declared in go
    getPartyData().then((data) => {
        let jsonData = JSON.parse(data)

        $('.dataSync[data-cat="party"]').each(function() {
            let field = $(this).data('field')

            if(field.includes('.')) {
                let arr = field.split('.'),
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
        }
    })
})
