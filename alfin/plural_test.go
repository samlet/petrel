package alfin

import (
	"fmt"
	"github.com/fatih/camelcase"
	"strings"
	"testing"
)

// ref: https://socketloop.com/tutorials/golang-convert-word-to-its-plural-form-example
func TestPluralize(t *testing.T) {
	fmt.Println("************************************")
	fmt.Println("*  convert singular to plural *")
	fmt.Println("************************************")
	// test empty string
	singular := ""
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "diagnosis"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "news"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "is"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "series"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "arm"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "status"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "dependency"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "goose"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "dolly"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "datum"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "genius"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "jones"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "tree"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "infantry"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "hero"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "zero"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "tons"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "tonus"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "temple"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "church"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "x"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "triangle"
	fmt.Println(singular + ":" + Pluralize(singular))

	singular = "colony"
	fmt.Println(singular + ":" + Pluralize(singular))

	fmt.Println("************************************")
	fmt.Println("*  sanity check tests on plurals *")
	fmt.Println("************************************")
	// what if already plural
	singular = "triangles"
	fmt.Println(singular+":", AlreadyPluralized(singular))

	singular = "shiites"
	fmt.Println(singular+":", AlreadyPluralized(singular))

	singular = "infantries"
	fmt.Println(singular+":", AlreadyPluralized(singular))

	singular = "torpedoes"
	fmt.Println(singular+":", AlreadyPluralized(singular))

	singular = "genii"
	fmt.Println(singular+":", AlreadyPluralized(singular))
	// retrieve the singular form
	fmt.Println(singular+" singular form is :", singularDictionary[singular])

	singular = "aces"
	fmt.Println(singular+":", AlreadyPluralized(singular))

	singular = "townies"
	fmt.Println(singular+":", AlreadyPluralized(singular))

	singular = "data"
	fmt.Println(singular+":", AlreadyPluralized(singular))
	// retrieve the singular form
	fmt.Println(singular+" singular form is :", singularDictionary[singular])

	singular = "joneses"
	fmt.Println(singular+":", AlreadyPluralized(singular))
}

func TestWordsInEntityName(t *testing.T) {
	ent:="ExampleItem"

	words:=camelcase.Split(ent)
	var result []string
	for i:=0;i<len(words)-1;i+=1{
		result=append(result, strings.ToLower(words[i]))
	}
	lastPart:=strings.ToLower(words[len(words)-1])
	result=append(result, Pluralize(lastPart))

	fmt.Printf("%s\n", strings.Join(result, "_"))
}