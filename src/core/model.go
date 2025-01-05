package core

import "github.com/google/uuid"

type CoreModel interface {
	Create(map[string]interface{}) (CoreModel, error)
	Search(uuid.UUID, map[uuid.UUID]interface{}) (interface{}, int, error)
	Remove(uuid.UUID, map[uuid.UUID]interface{}) (string, error)
	Update(uuid.UUID, map[uuid.UUID]interface{}, interface{}) (interface{}, error)
	Read(map[uuid.UUID]interface{}) error
}
