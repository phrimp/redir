package storage

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"redir/src/pkg"
)

var (
	localDataLocation = "reminder/data/data.bin"
	defaultKey        = "examplekey123456"
)

func SaveEncryptedStruct(data interface{}) error {
	key := []byte(os.Getenv("ENCRYPT_KEY"))
	if len(key) == 0 {
		key = []byte(defaultKey)
	}
	if err := SaveEncryptedStructProcess(localDataLocation, data, key); err != nil {
		return fmt.Errorf("error: %v", err)
	} else {
		log.Printf("Data successfully encrypted and saved to %s\n", localDataLocation)
		return nil
	}
}

func SaveEncryptedStructProcess(filename string, data interface{}, key []byte) error {
	plaintext, err := pkg.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal struct: %v", err)
	}
	ciphertext, nonce, err := pkg.AESGCMEncrypt(plaintext, key)
	if err != nil {
		return fmt.Errorf("failed to encrypt data: %v", err)
	}

	// Combine nonce and ciphertext
	var buffer bytes.Buffer
	buffer.Write(nonce)
	buffer.Write(ciphertext)

	// Write to file
	if err := os.WriteFile(filename, buffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
