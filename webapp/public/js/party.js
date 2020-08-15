$(document).ready(function() {
    cloneTemplate('#charInputTmpl', '#content')

    $('.charNew').click(() => {
        cloneTemplate('#charInputTmpl', '#content')
    })

    $('#content').on('click', '.charDelete', function() {
        let count = $('.tmpl').length

        if(count > 2) {
            $(this).closest('.tmpl').remove()
        }
    })
})
