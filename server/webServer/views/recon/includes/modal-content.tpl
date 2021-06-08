{{ if .Html }}
{{toHTML .Html }}
{{ else }}
<textarea id="outputCommand" class="p-3">
{{ .Content }}
</textarea>
{{ end }}