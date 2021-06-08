  <script src="/static/plugins/codemirror/codemirror.js"></script>

  <script src="/static/plugins/codemirror/addon/display/autorefresh.js"></script>
<script>
// Fill modal with content from link href
$("#exampleModal").on("show.bs.modal", function (e) {
    var link = $(e.relatedTarget);


    $(this).find(".modal-body").load(link.attr("href"), function () {
        $(function () {
            cm = CodeMirror.fromTextArea(document.getElementById("outputCommand"), {
                theme: "monokai",
                readOnly: true,
                autorefresh: true,
            });

            setTimeout(function () {
                cm.refresh()
            }, 200);

        })

    });

});
        
</script>