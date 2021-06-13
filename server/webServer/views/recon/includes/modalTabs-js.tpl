<script>
    // Fill modal with content from link href
    $("#exampleModal").on("show.bs.modal", function (e) {
        var link = $(e.relatedTarget);

        $(this).find(".modal-body").load(link.attr("href"), function () {});

    });
$('.modal-dialog').draggable({
    handle: ".modal-header"
});


$(document).ready(function () {

    $(".tabs").click(function () {

        $(".tabs").removeClass("active");
        $(".tabs h6").removeClass("font-weight-bold");
        $(".tabs h6").addClass("text-muted");
        $(this).children("h6").removeClass("text-muted");
        $(this).children("h6").addClass("font-weight-bold");
        $(this).addClass("active");

        current_fs = $(".active");

        next_fs = $(this).attr('id');
        next_fs = "#" + next_fs + "1";

        $("fieldset").removeClass("show");
        $(next_fs).addClass("show");

        current_fs.animate({}, {
            step: function () {
                current_fs.css({
                    'display': 'none',
                    'position': 'relative'
                });
                next_fs.css({
                    'display': 'block'
                });
            }
        });
    });

});

</script>