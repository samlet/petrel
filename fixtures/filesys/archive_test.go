package main

import (
	"fmt"
	"testing"
	"time"
)

func TestZipWriter(t *testing.T) {
	//Add some files to the archive.
	var files = []Content{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	cnt := ZipWriter(files)
	fmt.Printf("result size: %d\n", int64(len(cnt)))
}

func TestDateTime(te *testing.T) {
	t := time.Now()
	fmt.Println("Same, in UTC:", t.UTC().Format(time.RFC3339))
	year, month, day := t.UTC().Date()
	fmt.Printf("%4d%02d%2d\n", year, month, day)
}
