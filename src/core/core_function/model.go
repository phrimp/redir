package corefunction

import (
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

var jobs map[uuid.UUID]Job
