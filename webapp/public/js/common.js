function cloneTemplate(templateName, destination) {
    $(templateName).clone().attr('id', '').appendTo(destination)
}
