package corefunction

import (
	"redir/src/core"
	"redir/src/enum"

	"github.com/google/uuid"
)

type Job struct {
	Title  string
	Detail string
	Start  int64
	End    int64
	Status enum.JobStatus
}

var Jobs = map[uuid.UUID]core.CoreModel{}
