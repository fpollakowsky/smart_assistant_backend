package models

import "github.com/procyon-projects/chrono"

type Task struct {
	ID  int
	Job chrono.ScheduledTask
}
