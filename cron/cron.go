package cron

import "time"

type JobConfig struct {
	Name     string
	Every    time.Duration
	Schedule string
	Endpoint interface{}
}

type Job struct {
	ID       string
	Name     string
	Doc      string
	Every    time.Duration
	Schedule string
	Endpoint interface{}
}

func NewJob(id string, jobConfig JobConfig) *Job {
	panic("encore apps must be run using the encore command")
}
