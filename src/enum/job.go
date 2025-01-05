package enum

type JobStatus int

const (
	JobCreated JobStatus = iota
	JobPending
	JobInProgress
	JobCompleted
	JobFailed
)

func (s JobStatus) String() string {
	switch s {
	case JobCreated:
		return "Created"
	case JobPending:
		return "Pending"
	case JobInProgress:
		return "In Progress"
	case JobCompleted:
		return "Completed"
	case JobFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}
