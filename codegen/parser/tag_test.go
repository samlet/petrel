package main

import (
	"fmt"
	"reflect"
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