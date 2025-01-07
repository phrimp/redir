package core

import (
	"fmt"

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

func OSNotification(message, title, icon string) error {
	return beeep.Notify(title, message, icon)
}

func test() {
}
