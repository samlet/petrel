package meta

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
	"text/template"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "
	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func GenEntity(meta *EntityMeta, format string) string {
	var conv = map[string]interface{}{
		"title": strings.Title,
		"lower": strings.ToLower,
		"snake": ToSnakeCase,
		"prop":  properTitle,
	}

	b := &bytes.Buffer{}
	template.Must(template.New("").Funcs(conv).Parse(format)).Execute(b, meta)
	return b.String()
}

func GenEntityMetaFromTemplate(path string, entityName string) (string, error) {
	meta, err := GetEntityMeta(entityName)
	if err != nil {
		return "", err
	}
	tplData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	result := GenEntity(&meta.Data.Entity, string(tplData))
	return result, nil
}
