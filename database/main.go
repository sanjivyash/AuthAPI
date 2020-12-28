package database

import (
	"io/ioutil"
	"log"
)

func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("[ERROR] Error loading file\n" + err.Error())
	}

	return data
}

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}
