<script>
// Fill modal with content from link href
$("#exampleModal").on("show.bs.modal", function (e) {
    var link = $(e.relatedTarget);


    $(this).find(".modal-body").load(link.attr("href"), function () {
        

    });

});
        
</script>