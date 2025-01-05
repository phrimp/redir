package corefunction

import (
	"redir/src/core"
	"redir/src/enum"

	"github.com/gen2brain/beeep"
	"github.com/google/uuid"
)

func (j *Job) Create(params map[string]interface{}) (core.CoreModel, error) {
	new_job := Job{
		Title:  params["title"].(string),
		Detail: params["detail"].(string),
		Start:  params["start"].(int64),
		End:    params["end"].(int64),
		Status: params["status"].(enum.JobStatus),
	}

	return &new_job, nil
}

func (j *Job) Search(uuid.UUID, map[uuid.UUID]interface{}) (interface{}, int, error)
func (j *Job) Remove(uuid.UUID, map[uuid.UUID]interface{}) (string, error)
func (j *Job) Update(uuid.UUID, map[uuid.UUID]interface{}, interface{}) (interface{}, error)
func (j *Job) Read(map[uuid.UUID]interface{}) error

func OSNotification(message, title, icon string) error {
	return beeep.Notify(title, message, icon)
}
