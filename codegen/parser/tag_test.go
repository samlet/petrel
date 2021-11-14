package main

import (
	"fmt"
	"github.com/fatih/structtag"
	"reflect"
	"sort"
	"testing"
)

type T struct {
	f1     string "f one"
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five`
	f string `one:"1" two:"2" blank:""`
}

// ref: https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e
func TestStructTag(tt *testing.T) {
	t := reflect.TypeOf(T{})
	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag) // f one
	f4, _ := t.FieldByName("f4")
	fmt.Println(f4.Tag) // f four and five
	f5, _ := t.FieldByName("f5")
	fmt.Println(f5.Tag) // f four and five

	// ...
	t = reflect.TypeOf(T{})
	f, _ := t.FieldByName("f")
	fmt.Println(f.Tag) // one:"1" two:"2"blank:""
	v, ok := f.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
	v, ok = f.Tag.Lookup("two")
	fmt.Printf("%s, %t\n", v, ok) // 2, true
	v, ok = f.Tag.Lookup("blank")
	fmt.Printf("%s, %t\n", v, ok) // , true
	v, ok = f.Tag.Lookup("five")
	fmt.Printf("%s, %t\n", v, ok) // , false
}

func TestConversion(tt *testing.T) {
	type T1 struct {
		f int `json:"foo"`
	}
	type T2 struct {
		f int `json:"bar"`
	}

	t1 := T1{10}
	var t2 T2
	t2 = T2(t1)
	fmt.Println(t2) // {10}
}

func TestAddStructTags(tt *testing.T) {
	type t struct {
		t string `json:"foo,omitempty,string" xml:"foo"`
	}

	// get field tag
	tag := reflect.TypeOf(t{}).Field(0).Tag

	// ... and start using structtag by parsing the tag
	tags, err := structtag.Parse(string(tag))
	if err != nil {
		panic(err)
	}

	// iterate over all tags
	for _, t := range tags.Tags() {
		fmt.Printf("tag: %+v\n", t)
	}

	// get a single tag
	jsonTag, err := tags.Get("json")
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonTag)         // Output: json:"foo,omitempty,string"
	fmt.Println(jsonTag.Key)     // Output: json
	fmt.Println(jsonTag.Name)    // Output: foo
	fmt.Println(jsonTag.Options) // Output: [omitempty string]

	// change existing tag
	jsonTag.Name = "foo_bar"
	jsonTag.Options = nil
	tags.Set(jsonTag)

	// add new tag
	tags.Set(&structtag.Tag{
		Key:     "hcl",
		Name:    "foo",
		Options: []string{"squash"},
	})

	// print the tags
	fmt.Println(tags) // Output: json:"foo_bar" xml:"foo" hcl:"foo,squash"

	// sort tags according to keys
	sort.Sort(tags)
	fmt.Println(tags) // Output: hcl:"foo,squash" json:"foo_bar" xml:"foo"
}