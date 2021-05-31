package alfin

import (
	"encoding/json"
	"io"
	"reflect"
	"strings"
	"text/template"
)

func GenTemplate(text string, tmpl string, wr io.Writer) {
	var err error
	m := make(map[string]interface{})
	if err = json.Unmarshal([]byte(text), &m); err != nil {
		panic(err)
	}

	// https://gist.github.com/alex-leonhardt/8ed3f78545706d89d466434fb6870023
	tf := template.FuncMap{
		"isInt": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
				return true
			default:
				return false
			}
		},
		"isString": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.String:
				return true
			default:
				return false
			}
		},
		"isSlice": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Slice:
				return true
			default:
				return false
			}
		},
		"isArray": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Array:
				return true
			default:
				return false
			}
		},
		"isMap": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Map:
				return true
			default:
				return false
			}
		},
		"isList":    IsList,
		"isNumber":  IsNumber,
		"isFloat":   IsFloat,
		"fieldType": FieldType,
		"title":     strings.Title,
	}

	t := template.New("hello").Funcs(tf)
	tt, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	if err = tt.Execute(wr, &m); err != nil {
		panic(err)
	}
}

func IsList(i interface{}) bool {
	v := reflect.ValueOf(i).Kind()
	return v == reflect.Array || v == reflect.Slice
}

func IsNumber(i interface{}) bool {
	v := reflect.ValueOf(i).Kind()
	switch v {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func IsInt(i interface{}) bool {
	v := reflect.ValueOf(i).Kind()
	switch v {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func IsFloat(i interface{}) bool {
	v := reflect.ValueOf(i).Kind()
	return v == reflect.Float32 || v == reflect.Float64
}

func FieldType(typeName string) string {
	var goType string
	switch typeName {
	case "id":
		goType = "string"
	case "blob", "byte-array", "object":
		goType = "[]byte"
	case "date-time", "date", "time":
		goType = "string"
	case "currency-amount", "currency-precise", "fixed-point", "floating-point":
		goType = "float64"
	case "integer":
		goType = "int32"
	case "numeric":
		goType = "int64"
	case "indicator":
		goType = "byte"
	default:
		goType = "string"
	}
	return goType
}
