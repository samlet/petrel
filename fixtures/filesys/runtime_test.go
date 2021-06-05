package main

import (
	"runtime"
	"testing"
)

func getCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}

func TestRuntimeDir(t *testing.T) {
	println(getCurrentFilePath())
}