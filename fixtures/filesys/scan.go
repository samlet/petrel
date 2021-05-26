package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func TimeTrack(start time.Time, format string, v ...interface{}) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", fmt.Sprintf(format, v...), elapsed)
}

func WalkMatch(root, pattern string) ([]string, error) {
	defer TimeTrack(time.Now(), "walk '%s'", root)

	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
