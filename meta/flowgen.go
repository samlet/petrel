package meta

import (
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"
)

type FlowMeta struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Activities  []ActivityMeta `json:"activities"`
}

type ActivityMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GenFlow(meta *FlowMeta, format string) string {
	var conv = map[string]interface{}{
		"title": strings.Title,
		"lower": strings.ToLower,
		"snake": ToSnakeCase,
	}

	b := &bytes.Buffer{}
	template.Must(template.New("").Funcs(conv).Parse(format)).Execute(b, meta)
	return b.String()
}

func GenFlowFromTemplate(meta *FlowMeta, path string) (string, error) {
	tplData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	result := GenFlow(meta, string(tplData))
	return result, nil
}

func CreateFlowMeta(flowName string, names []string) FlowMeta {
	flowMeta := FlowMeta{Name: flowName,
		Activities: []ActivityMeta{}}
	for _, name := range names {
		flowMeta.Activities = append(flowMeta.Activities, ActivityMeta{Name: name})
	}
	return flowMeta
}
