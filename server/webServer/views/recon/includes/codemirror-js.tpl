<!-- CodeMirror -->
  <script src="/static/plugins/codemirror/codemirror.js"></script>

  <script src="/static/plugins/codemirror/mode/yaml/yaml.js"></script>
<script>
        // CodeMirror
        $(function () {
      CodeMirror.fromTextArea(document.getElementById("codeMirrorDemo"), {
        theme: "monokai",
        lineNumbers: true,
        readOnly: true
      });
    })
    </script>