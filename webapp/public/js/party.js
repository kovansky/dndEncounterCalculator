$(document).ready(function() {
    cloneTemplate('#charInputTmpl', '#content')

    $('.charNew').click(() => {
        cloneTemplate('#charInputTmpl', '#content')
    })
})
