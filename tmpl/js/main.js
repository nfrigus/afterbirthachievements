$(document).ready(function () {
    $("#username-input").keypress(function(e) {
        if (e.which == 13) {
            document.location = "/" + $("#username-input").val()
        }
    });

    // Simple mobile-friendly tooltips; adopted (i.e. stolen) from
    // https://stackoverflow.com/questions/12539006/tooltips-for-mobile-browsers
    $(".tooltip").attr("data-close", false);

    // Capture mouse click events
    $(".has-tooltip").click(function () {
        var $title = $(this).find(".tooltip");
        if (!$title.attr("data-close"))
            $title.show();
        $title.attr("data-close", false);
    });

    $(".tooltip").click(function () {
        $(this).attr("data-close", true);
    });

    $(document).mouseup(function () {
        $(".tooltip").hide();
    });

    // Capture mouse enter/leave events
    $(".has-tooltip").mouseenter(function () {
        var $title = $(this).find(".tooltip");
        $title.show();
    });

    $(".has-tooltip").mouseleave(function () {
        var $title = $(this).find(".tooltip");
        $title.hide();
    });  
});