package job

import (
	"redir/src/core"
	"redir/src/enum"
)

func (j *Job) Create(params map[string]interface{}) (core.CoreModel, error) {
	new_job := Job{
		Title:  params["title"].(string),
		Detail: params["detail"].(string),
		Start:  params["start"].(int64),
		End:    params["end"].(int64),
		Status: enum.JobCreated,
	}

	return &new_job, nil
}

func (j *Job) Update(core.CoreModel) (core.CoreModel, error) {
	return nil, nil
}
