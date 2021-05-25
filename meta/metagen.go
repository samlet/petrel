package meta

import (
	"bytes"
	"text/template"
)

const ENTITY_DEF = `
type {{.Name}} struct {
{{- range .Fields}}
	{{.Name}} {{.Type}}
{{- end}}
}
`

func GenEntity(meta *EntityMeta, format string) string {
	b := &bytes.Buffer{}
	template.Must(template.New("").Parse(format)).Execute(b, meta)
	return b.String()
}
