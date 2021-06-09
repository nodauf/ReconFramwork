{{ if .Html }}
<div id="outputCommand" class="p-3">{{toHTML .Content }}</div>
{{ else }}
<div id="outputCommand" class="p-3">{{ .Content }}</div>
{{ end }}