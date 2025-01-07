package job

import (
	"fmt"
	"redir/src/core"
	"redir/src/enum"
	"strconv"
)

func (j *Job) Create(params map[string]string) (core.CoreModel, error) {
	start, err := strconv.Atoi(params["start"])
	if err != nil {
		return nil, fmt.Errorf("error at converting start to int:, %v", err)
	}
	end, err := strconv.Atoi(params["end"])
	if err != nil {
		return nil, fmt.Errorf("error at converting end to int:, %v", err)
	}
	new_job := Job{
		Title:  params["title"],
		Detail: params["detail"],
		Start:  int64(start),
		End:    int64(end),
		Status: enum.JobCreated,
	}

	return &new_job, nil
}

func (j *Job) Update(core.CoreModel) (core.CoreModel, error) {
	return nil, nil
}
