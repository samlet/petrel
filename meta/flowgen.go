package meta

import (
	"bytes"
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
