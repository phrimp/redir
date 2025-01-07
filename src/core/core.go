package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/gen2brain/beeep"
	"github.com/google/uuid"
)

// APPLICATION CORE STRUCTURE

func AddToRamMemory(params map[string]string, memory *map[uuid.UUID]CoreModel, model CoreModel) error {
	var core_model CoreModel = model
	new_model, err := core_model.Create(params)
	if err != nil || new_model == nil {
		return fmt.Errorf("error at Creating Model: %v", err)
	}
	(*memory)[uuid.New()] = new_model
	return nil
}

func ReadCoreModel(memory map[uuid.UUID]CoreModel, filePath string) {
	// Open or create the file, truncating it if it exists
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Check if the map is empty
	if len(memory) == 0 {
		log.Println("Reading CoreModel: Current memory is empty")
		return
	}

	// Iterate over the map to determine field names from the first value
	var exampleValue CoreModel
	for _, v := range memory {
		exampleValue = v
		break
	}

	val := reflect.ValueOf(exampleValue)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		log.Fatalf("Map values must be pointers to structs. Found kind: %s", val.Kind())
		return
	}

	// Get the struct type
	elemType := val.Elem().Type()

	// Write headers
	fmt.Fprintf(file, "%-40s", "ID") // First column is the map key
	for i := 0; i < elemType.NumField(); i++ {
		fmt.Fprintf(file, "%-15s", elemType.Field(i).Name)
	}
	fmt.Fprintln(file) // Add a newline after headers

	// Write rows
	for key, value := range memory {
		val := reflect.ValueOf(value)
		if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
			log.Printf("Skipping key %v: value is not a pointer to a struct", key)
			continue
		}

		// Write the map key (UUID)
		fmt.Fprintf(file, "%-40s", key.String())

		// Write the struct fields
		row := val.Elem()
		for j := 0; j < row.NumField(); j++ {
			fmt.Fprintf(file, "%-15v", row.Field(j).Interface())
		}
		fmt.Fprintln(file) // Add a newline after each row
	}

	currentdir, _ := os.Getwd()
	message := fmt.Sprintf("Data written to file: %s\n", filepath.Join(currentdir, filepath.Base(filePath)))
	log.Println(message)
	OSNotification(message, "List", "")
}

func OSNotification(message, title, icon string) error {
	return beeep.Notify(title, message, icon)
}

func test() {
}
