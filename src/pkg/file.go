package pkg

import (
	"fmt"
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

func ClearDirectory(dirPath string) error {
	// Read the directory contents
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// Iterate through each file and remove it
	for _, file := range files {
		filePath := dirPath + string(os.PathSeparator) + file.Name()

		// Remove file or directory
		err := os.RemoveAll(filePath)
		if err != nil {
			return fmt.Errorf("failed to remove file/directory %s: %w", filePath, err)
		}
	}

	return nil
}
