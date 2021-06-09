<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>AdminLTE 3 | Dashboard 2</title>

  <!-- Google Font: Source Sans Pro -->
  <link rel="stylesheet"
    href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,400i,700&display=fallback">
  <!-- Font Awesome Icons -->
  <link rel="stylesheet" href="/static/plugins/fontawesome-free/css/all.min.css">
  <!-- overlayScrollbars -->
  <link rel="stylesheet" href="/static/plugins/overlayScrollbars/css/OverlayScrollbars.min.css">
  <!-- Theme style -->
  <link rel="stylesheet" href="/static/dist/css/adminlte.min.css">
   <link rel="stylesheet" href="/static/css/custom.css">

  <link rel="stylesheet" href="/static/plugins/select2/css/select2.css">

  {{ if .Select2}}
    {{ template "/recon/includes/select2-css.tpl" }}
  {{ end }}

  {{ if .DataTables}}
    {{ template "/recon/includes/datatables-css.tpl" }}
  {{ end }}

  {{ if .Toastr}}
    {{ template "/recon/includes/toastr-css.tpl" }}
  {{ end }}

   {{ if or (.CodeMirror) (.Tree)}}
     {{ template "/recon/includes/codemirror-css.tpl" }}
  {{ end }}

</head>

<body class="hold-transition dark-mode sidebar-mini layout-fixed layout-navbar-fixed layout-footer-fixed">
  <div class="wrapper">

    <!-- Preloader -->
    <div class="preloader flex-column justify-content-center align-items-center">
      <img class="animation__wobble" src="/static/dist/img/AdminLTELogo.png" alt="AdminLTELogo" height="60" width="60">
    </div>

    {{ template "recon/includes/navbar.tpl" }}

    {{ template "recon/includes/sidebar.tpl" }}

      <!-- Content Wrapper. Contains page content -->
     <div class="content-wrapper">
        {{ if .flash.error }}
         <div class="alert alert-danger">
          {{ .flash.error}}
          </div>
        {{ end }}
         {{ if .flash.success }}
         <div class="alert alert-success">
          {{ .flash.success }}
          </div>
        {{ end }}
        {{ if or (.Modal) (.Tree)}}
          {{ template "/recon/includes/modal-layout.tpl"}}
        {{ end }}

    {{ .LayoutContent}}
    </div>
  <!-- /.content-wrapper -->

    <!-- Control Sidebar -->
    <aside class="control-sidebar control-sidebar-dark">
      <!-- Control sidebar content goes here -->
    </aside>
    <!-- /.control-sidebar -->

    <!-- Main Footer -->
    <footer class="main-footer">
      <strong>Copyright &copy; 2014-2021 <a href="https://adminlte.io">AdminLTE.io</a>.</strong>
      All rights reserved.
      <div class="float-right d-none d-sm-inline-block">
        <b>Version</b> 3.1.0
      </div>
    </footer>
  </div>
  <!-- ./wrapper -->

  <!-- REQUIRED SCRIPTS -->
  <!-- jQuery -->
  <script src="/static/plugins/jquery/jquery.min.js"></script>
  <!-- Bootstrap -->
  <script src="/static/plugins/bootstrap/js/bootstrap.bundle.min.js"></script>
  <!-- overlayScrollbars -->
  <script src="/static/plugins/overlayScrollbars/js/jquery.overlayScrollbars.min.js"></script>
  <!-- AdminLTE App -->
  <script src="/static/dist/js/adminlte.js"></script>

  <!-- PAGE PLUGINS -->
  <!-- jQuery Mapael -->
  <script src="/static/plugins/jquery-mousewheel/jquery.mousewheel.js"></script>
  <script src="/static/plugins/raphael/raphael.min.js"></script>
  <script src="/static/plugins/jquery-mapael/jquery.mapael.min.js"></script>
  <script src="/static/plugins/jquery-mapael/maps/usa_states.min.js"></script>
  <!-- ChartJS -->
  <script src="/static/plugins/chart.js/Chart.min.js"></script>
  <script>
  // Enable tooltips everywhere
  $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
</script>
  {{ if .Select2}}
    {{ template "/recon/includes/select2-js.tpl" }}
  {{ end }}

  {{ if .DataTables}}
    {{ template "/recon/includes/datatables-js.tpl" }}
  {{ end }}

  {{ if .CodeMirror}}
    {{ template "/recon/includes/codemirror-js.tpl" }}
  {{ end }}

  {{ if .Toastr}}
    {{ template "/recon/includes/toastr-js.tpl" .Toastr}}
  {{ end }}

  {{ if .Modal}}
    {{ template "/recon/includes/modal-js.tpl" }}
  {{ end }}

  {{ if .Tree}}
    {{ template "/recon/includes/tree-js.tpl" .Tree }}
  {{ end }}

</html>