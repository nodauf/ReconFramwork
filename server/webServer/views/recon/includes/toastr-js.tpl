  <script src="/static/plugins/toastr/toastr.min.js"></script>

<script>
    $(window).ready(function() {
    toastr.options = {
  "closeButton": true,
  "positionClass": "toast-bottom-right",
}
      toastr.success({{ . }})
    });
</script>