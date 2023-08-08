package worker

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/procyon-projects/chrono"
	"home-ai-backend/models"
	"home-ai-backend/pkg/read"
	"home-ai-backend/server/logging"
	"home-ai-backend/server/mqtt"
	"os"
	"strconv"
)

//
// Workers for trashcan application
//

var ScheduledTasks []models.Task

func InitializeCron() {
	// add cron jobs from db
	err, routines := read.Routine()
	if err != nil {
		logging.PrintListFail("Initialize Cron")
		os.Exit(20)
	}

	for i := 0; i < len(routines); i++ {
		if *routines[i].Status == true {
			for y := 0; y < len(routines[i].Payload); y++ {
				err = NewJob(
					routines[i].TriggerTime,
					routines[i].Payload[y].Device.Channel,
					routines[i].Payload[y].Device.Room,
					strconv.Itoa(routines[i].Payload[y].Value),
					routines[i].ID,
				)
				if err != nil {
					log.Error("Failed setting up cronjob", "Error", err.Error())
				}
			}
			log.Debug("Routine added as cronjob: " + routines[i].Title)
		}
	}
	logging.PrintListDone("Initialize Cron")
}

func task(channel, room, payload string) {
	mqtt.NewRequest(channel, room, payload)
}

func NewJob(triggerTime, channel, room, payload string, id int) error {
	// add job
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
		data := models.Task{
			ID:  id,
			Job: task,
		}
		ScheduledTasks = append(ScheduledTasks, data)
	}

	return nil
}

func RemoveWorker(id int) {
	// remove job
	for i := 0; i < len(ScheduledTasks); i++ {
		if id == ScheduledTasks[i].ID {
			ScheduledTasks[i].Job.Cancel()
			log.Debug("Routine cronjob with id " + strconv.Itoa(id) + "removed")
		}
	}
}
