package cron

type Duration int64

const (
	Minute Duration = 60
	Hour   Duration = 60 * Minute
)

type JobConfig struct {
	Name     string
	Every    Duration
	Schedule string
	Endpoint interface{}
}

type Job struct {
	ID       string
	Name     string
	Doc      string
	Every    Duration
	Schedule string
	Endpoint interface{}
}

func NewJob(id string, jobConfig JobConfig) *Job {
	panic("encore apps must be run using the encore command")
}
