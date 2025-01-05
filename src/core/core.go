package core

import (
	"log"

	"github.com/google/uuid"
)

// APPLICATION CORE STRUCTURE

func AddToRamMemory(params map[string]interface{}, memory *map[uuid.UUID]CoreModel) {
	var model CoreModel
	new_model, err := model.Create(params)
	if err != nil || new_model == nil {
		log.Println("Error at Creating Model:", err)
		return
	}
	(*memory)[uuid.New()] = new_model
}

func test() {
}
