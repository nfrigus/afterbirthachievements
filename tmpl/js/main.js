$(document).ready(function () {
    $("#username-input").keypress(function(e) {
        if (e.which == 13) {
            document.location = "/" + $("#username-input").val()
        }
    });
});