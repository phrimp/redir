package pkg

import (
	"log"
	"os"
)

func CreateMultipleDirs(locations []string) bool {
	err_num := 0
	for _, location := range locations {
		err := CreateDir(location)
		if err != nil {
			log.Println("Error Create Directory at", location)
			err_num++
		}
	}
	return err_num <= 2
}

func CreateDir(location string) error {
	return os.MkdirAll(location, 0755)
}

func InitLogFile() error {
	return nil
}
