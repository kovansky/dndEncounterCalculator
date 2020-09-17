function cloneTemplate(templateName, destination) {
    $(templateName).clone().attr('id', '').appendTo(destination)
}

function lockWindow() {
    let el = $("#lock")

    if(!el.length) {
        el = $('body').append($('<div></div>')
            .attr('id', 'lock')
            .addClass('lockVisible')
        )
    } else {
        el.addClass('lockVisible')
    }
}

function unlockWindow() {
    let el = $('#lock')

    if(el.length) {
        el.removeClass('lockVisible')
    }
}

function idFromString(str) {
    return str.replace(/\W/, '')
}
