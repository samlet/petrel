package alfin

import (
	"encoding/json"
	"io"
	"log"
	"reflect"
	"strings"
	"text/template"
)

func (c *GenHelper) GenTemplate(text string, tmpl string, wr io.Writer) {
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
		"title":     strings.Title,
		"fieldType": FieldType,
		"paramType": ParamType,
		"ethType": EthType,
	}

	t := template.New("hello").Funcs(tf)
	tt, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	if err = tt.Execute(wr, &m); err != nil {
		panic(err)
	}

	c.Logger.Info("Generate ok.")
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
	case "date-time", "date":
		goType = "*services.DateTime"
	case "time":
		goType = "string"
	case "currency-amount", "currency-precise", "fixed-point":
		goType = "*services.Decimal"
	case "floating-point":
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

func EthType(typeName string) string {
	var ethType string
	switch typeName {
	case "id":
		ethType = "bytes20"
	case "id-long":
		ethType = "bytes32"
	case "id-vlong":
		ethType = "bytes"
	case "blob", "object":
		ethType = "bytes"
	case "byte-array":
		ethType="uint[]"
	case "date-time", "date":
		ethType = "uint256"
	case "time":
		ethType = "uint256"
	case "currency-amount", "currency-precise", "fixed-point":
		ethType = "int256"
	case "floating-point":
		ethType = "int256"
	case "integer":
		ethType = "int32"
	case "numeric":
		ethType = "int64"
	case "indicator":
		ethType = "bytes1"
	case "very-short":
		ethType = "bytes10"
	default:
		ethType = "string"
	}
	return ethType
}

func ParamType(typeName string, mode string) string {
	var goType string
	switch typeName {
	case "String":
		goType = "string"
	case "Integer":
		goType = "int32"
	case "Long":
		goType = "int64"
	case "Float":
		goType = "float32"
	case "Double":
		goType = "float64"
	case "Boolean":
		goType = "bool"
	case "Timestamp", "java.sql.Timestamp":
		goType = "*services.Timestamp"
	case "BigDecimal", "java.math.BigDecimal":
		//goType = "string"
		goType = "*services.Decimal"
	case "java.sql.Date", "java.util.Date":
		goType = "*services.DateTime"
	case "java.sql.Time":
		goType = "string"
	case "List", "java.util.List":
		goType = "[]interface{}"
	case "GenericEntity", "org.apache.ofbiz.entity.GenericValue":
		if mode == "out" {
			goType = "map[string]interface{}"
		} else {
			goType = "*services.MetaValue"
		}
	case "Map", "java.util.Map":
		goType = "map[string]interface{}"
	default:
		log.Printf("Cannot convert parameter type %s, default to interface{} \n", typeName)
		goType = "interface{}"
	}
	return goType
}

