package storage_test

import (
	"fmt"
	"os"
	"reminder/src/storage"
	"testing"
)

type ExampleStruct struct {
	Name  string
	Age   int
	Email string
}

func TestSaveEncryptedStruct(t *testing.T) {
	// Define a valid encryption key
	key := []byte("examplekey123456") // Replace with a secure key

	tests := []struct {
		name       string
		data       interface{}
		filename   string
		key        []byte
		wantErr    bool
		validateFn func(filename string, key []byte, data interface{}) error
	}{
		{
			name:     "Valid data",
			data:     ExampleStruct{Name: "Alice", Age: 30, Email: "alice@example.com"},
			filename: "test_valid_data.bin",
			key:      key,
			wantErr:  false,
			validateFn: func(filename string, key []byte, data interface{}) error {
				// Verify that the file was created and contains encrypted data
				info, err := os.Stat(filename)
				if err != nil {
					return err
				}
				if info.Size() <= 0 {
					return fmt.Errorf("file is empty")
				}
				return nil
			},
		},
		{
			name:     "Invalid key length",
			data:     ExampleStruct{Name: "Bob", Age: 40, Email: "bob@example.com"},
			filename: "test_invalid_key.bin",
			key:      []byte("short"),
			wantErr:  true,
		},
		{
			name:     "Empty data",
			data:     ExampleStruct{},
			filename: "test_empty_data.bin",
			key:      key,
			wantErr:  false,
			validateFn: func(filename string, key []byte, data interface{}) error {
				// Verify the file is created but data is empty struct
				info, err := os.Stat(filename)
				if err != nil {
					return err
				}
				if info.Size() <= 0 {
					return fmt.Errorf("file is empty")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := storage.SaveEncryptedStructProcess(tt.filename, tt.data, tt.key)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SaveEncryptedStruct() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SaveEncryptedStruct() succeeded unexpectedly")
			}
			if tt.validateFn != nil {
				if err := tt.validateFn(tt.filename, tt.key, tt.data); err != nil {
					t.Errorf("Validation failed: %v", err)
				}
			}

			// Clean up test files
			if err := os.Remove(tt.filename); err != nil {
				t.Errorf("Failed to clean up test file: %v", err)
			}
		})
	}
}
