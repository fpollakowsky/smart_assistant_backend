package cron

import (
	"context"
	"github.com/procyon-projects/chrono"
	"log"
	"shome-backend/mqtt"
)

func task(channel, room, payload string) {
	var client = mqtt.Connect()
	mqtt.NewRequest(client, channel, room, payload)
}

func Cron(min, hour, day, channel, room, payload string, remove bool) error {
	s := chrono.NewDefaultTaskScheduler()

	task, err := s.ScheduleWithCron(func(ctx context.Context) {
		task(channel, room, payload)
	},
		//min +" " + hour +" " + "* *" + day,
		"0 "+min+" "+hour+" * * "+day,
		chrono.WithLocation("Europe/Berlin"))

	if err == nil {
		log.Print("Task has been scheduled")
	}

	if remove {
		task.Cancel()
	}
	return nil
}

func RemoveCron(task interface{}) {
}
