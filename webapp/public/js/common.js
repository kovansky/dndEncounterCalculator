/*
 * This file contains global-wide functions, applied to all views
 */

// This function is used to clone templates. It clones the templateName (jQuery selector) into desired destination,
// and removes the ID attribute to omit ID-related problems
function cloneTemplate(templateName, destination) {
    $(templateName).clone().attr("id", "").appendTo(destination)
}

// This function shows program lock (used to block interaction, when dialog is open), or creates it, if doesn't exist
function lockWindow() {
    let el = $("#lock")

    if(!el.length) {
        el = $("body").append($("<div></div>")
            .attr("id", "lock")
            .addClass("lockVisible")
        )
    } else {
        el.addClass("lockVisible")
    }
}

// This function hides program lock
function unlockWindow() {
    let el = $("#lock")

    if(el.length) {
        el.removeClass("lockVisible")
    }
}

// Converts string to id (removes all "bad" characters, like spaces, and transforms to lower case)
function idFromString(str) {
    return str.replace(/\W/, "").toLowerCase()
}
