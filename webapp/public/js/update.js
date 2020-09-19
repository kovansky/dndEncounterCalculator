$(document).ready(function() {
    // Declared in go
    loadWindowState().then((data) => {
        let jsonData = JSON.parse(data)

        $(".dataSync[data-cat=\"current\"]").each(function() {
            let field = $(this).data("field")

            if(field.includes(".")) {
                let arr = field.split("."),
                    actualField = jsonData.current

                for(let i = 0; i < arr.length; i++) {
                    actualField = actualField[arr[i]]
                }

                $(this).html(actualField)
            } else {
                $(this).html(jsonData.current[field])
            }
        })

        $(".dataSync[data-cat=\"remote\"]").each(function() {
            let field = $(this).data("field")

            if(field.includes(".")) {
                let arr = field.split("."),
                    actualField = jsonData.remote

                for(let i = 0; i < arr.length; i++) {
                    actualField = actualField[arr[i]]
                }

                $(this).html(actualField)
            } else {
                $(this).html(jsonData.remote[field])
            }
        })
    })

    $('.okBtn').click(() => {
        // Declared in go
        retValue(1)
    })

    $('.ignoreBtn').click(() => {
        // Declared in go
        retValue(0)
    })
})
