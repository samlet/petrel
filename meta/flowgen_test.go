package meta

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestGenFlowFromTemplate(t *testing.T) {
	flowMeta := FlowMeta{Name: "SampleFlow",
		Activities: []ActivityMeta{
			{Name: "downloadFile"},
			{Name: "processFile"},
			{Name: "uploadFile"},
		}}
	result, err := GenFlowFromTemplate(&flowMeta, "../incls/workflow_def.tmpl")
	assert.True(t, err == nil)
	println(result)
}

func TestGenFlowTestFromTemplate(t *testing.T) {
	flowMeta := FlowMeta{Name: "SampleFlow",
		Activities: []ActivityMeta{
			{Name: "downloadFile"},
			{Name: "processFile"},
			{Name: "uploadFile"},
		}}
	result, err := GenFlowFromTemplate(&flowMeta, "../incls/workflow_test.tmpl")
	assert.True(t, err == nil)
	println(result)
}

func TestConvertFlowMeta(t *testing.T) {
	names := []string{"downloadFile", "processFile"}
	flowMeta := CreateFlowMeta("Sample", names)
	result, _ := json.MarshalIndent(flowMeta, "", "  ")
	println(string(result))
}
