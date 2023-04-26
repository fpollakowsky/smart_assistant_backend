package worker

import (
	"context"
	"github.com/procyon-projects/chrono"
	"log"
	"shome-backend/server/mqtt"
)

//
// Workers for trashcan application
//

func task(channel, room, payload string) {
	var client = mqtt.Connect()
	mqtt.NewRequest(client, channel, room, payload)
}

func NewJob(triggerTime, channel, room, payload string, remove bool) error {
	s := chrono.NewDefaultTaskScheduler()

	task, err := s.ScheduleWithCron(func(ctx context.Context) {
		task(channel, room, payload)
	},
		triggerTime,
		chrono.WithLocation("Europe/Berlin"),
	)

	if err != nil {
		log.Fatal(err.Error())
	} else {
		if !remove {
			log.Print("Task has been scheduled")
		} else {
			log.Print("Tasks canceled")
			task.Cancel()
		}
	}

	return nil
}

func RemoveWorker(task interface{}) {
}
