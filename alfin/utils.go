package alfin

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}


func ReadJsonFile(file string, targetObj interface{}) error {
	configData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Failed to read file: %v, Error: %v", file, err)
		return err
	}

	if err := json.Unmarshal(configData, targetObj); err != nil {
		log.Printf("Error initializing object from file %s: %v", file, err)
		return err
	}
	return nil
}

