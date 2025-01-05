package corefunction

import (
	"redir/src/core"
	"redir/src/enum"

	"github.com/gen2brain/beeep"
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

func (j *Job) Update(core.CoreModel) (core.CoreModel, error) {
	return nil, nil
}

func OSNotification(message, title, icon string) error {
	return beeep.Notify(title, message, icon)
}
