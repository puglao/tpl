/*
Copyright © 2022 puglao <puglao@cloudemo.site>

*/
package internal

import (
	"errors"
	"io/fs"
	"log"
	"os"
)

func IsFileExist(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
	}
	return true
}

func NewFile(filepath string) {
	// create new empty file
	_, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFile(filepath string, data []byte) {
	// create new file with content
	_, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filepath, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
