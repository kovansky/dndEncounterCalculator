/*
 * This is script file (local controller) for Update View
 */

$(document).ready(function() {
    // Get initialization data from backend and load them to the view
    // Declared in go
    loadWindowState().then((data) => {
        // Parse string into js object
        let jsonData = JSON.parse(data)

        // Iterate through .dataSync objects in DOM to fill them with values
        $(".dataSync[data-cat=\"current\"]").each(function() {
            // Get the object field (requested data)
            let field = $(this).data("field")

            if(field.includes(".")) {
                let arr = field.split("."),
                    actualField = jsonData.current

                for(let i = 0; i < arr.length; i++) {
                    actualField = actualField[arr[i]]
                }

                // If requested field contains dot (has more than one level), split it
                $(this).html(actualField)
            } else {
                // Set object html to the requested field value
                $(this).html(jsonData.current[field])
            }
        })


        // Iterate through .dataSync objects in DOM to fill them with values
        $(".dataSync[data-cat=\"remote\"]").each(function() {
            // Get the object field (requested data)
            let field = $(this).data("field")

            if(field.includes(".")) {
                let arr = field.split("."),
                    actualField = jsonData.remote

                for(let i = 0; i < arr.length; i++) {
                    actualField = actualField[arr[i]]
                }

                // If requested field contains dot (has more than one level), split it
                $(this).html(actualField)
            } else {
                // Set object html to the requested field value
                $(this).html(jsonData.remote[field])
            }
        })
    })

    // Add ok button handler
    $(".okBtn").click(() => {
        // Return value to backend
        // Declared in go
        retValue(1)
    })

    // Add ignore button handler
    $(".ignoreBtn").click(() => {
        // Return value to backend
        // Declared in go
        retValue(0)
    })
})
