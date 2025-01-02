package storage

type Job struct {
	Name   string
	Detail string
	Start  int64
	End    int64
	Status JobStatus
}

type JobStatus struct {
	Name  string
	Value int
}
