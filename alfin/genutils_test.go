package alfin

import (
	"os"
	"testing"
)

func TestGenTemplate(t *testing.T) {
	text := `
{
	"msg": "hello world",
	"msgs": ["hello", "gophers", "..."],
	"msgNum": [0, 1, 2, 3, 4, 5],
	"nested": [
		{"msg": "why"},
		{"msg": "did"},
		{"msg": "the"},
		{"msg": [0, 1, 1, 3, 2]},
		{"msg": "chicken"},
		{"msg": "cross"},
		99,
		["yolo", "yolo", "yolo"]
	]
}`

	tmpl := `
{{ range $k, $v := $.msgs }}Key:{{ $k }}, Value:{{ $v }}
{{ end }}
{{ range $_, $v := $.msgNum }}Values: {{ $v }}
{{ end }}
{{ $.nested }}
{{ range $_, $v := $.nested }}
	{{ if isInt $v }}
	v is int .. {{ $v }}
	{{ end }}
	{{- if isMap $v -}}
	{{- range $k, $v := $v -}}
		k={{ $k }}, v={{ $v }}
	{{- end -}}
	{{- end -}}
	{{- if isSlice $v -}}
		{{ range $_, $s := $v -}}
			{{ $s }}
		{{- end }}
	{{- end -}}
{{ end }}
`
	GenTemplate(text, tmpl, os.Stdout)
}
