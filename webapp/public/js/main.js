$(document).ready(function() {
    cloneTemplate('#monsterInputTmpl', '#monstersForm .fields')

    $('.addMonster').click(() => {
        cloneTemplate('#monsterInputTmpl', '#monstersForm .fields')
    })
})
