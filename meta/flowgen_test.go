package meta

import "testing"

const GEN_SAMPLE = `
/**
 * Activities used by file processing {{.Name}} workflow.
 */
const (
{{- range .Activities}}
	{{.Name}}Name = "{{.Name}}"
{{- end}}
)
`

func TestGenFlow(t *testing.T) {
	flowMeta := FlowMeta{Name: "SampleFlow",
		Activities: []ActivityMeta{
			{Name: "downloadFileActivity"},
			{Name: "processFileActivity"},
			{Name: "uploadFileActivity"},
		}}
	result := GenFlow(&flowMeta, GEN_SAMPLE)
	println(result)
}
