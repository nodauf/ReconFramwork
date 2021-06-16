                 <div class="row d-flex  pb-4">
                  {{ range $key, $value := .Content}}
                     <div class="tabs pr-4" id="tab0{{ $key }}">
                         <h6 class="text-muted">{{ $key }}</h6>
                     </div>
                  {{ end }}
                     <div class="line"></div>
                 </div>
                  {{ range $key, $value := .Content}}
                     <fieldset id="tab0{{ $key }}1">
                     {{ if $.Html }}
                        <div id="outputCommand" class="p-3">{{toHTML $value }}</div>
                    {{ else }}
                        <div id="outputCommand" class="p-3">{{ $value }}</div>
                    {{ end }}
                    </fieldset>
                  {{ end }}
   <script>              
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