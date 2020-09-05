$(document).ready(function() {
    cloneTemplate("#monsterInputTmpl", "#monstersForm .fields")

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
