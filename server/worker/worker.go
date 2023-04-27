package worker

import (
	"context"
	"github.com/procyon-projects/chrono"
	"log"
	"os"
	"shome-backend/models"
	"shome-backend/pkg/read"
	"shome-backend/server/mqtt"
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
		log.New(os.Stdout, "[INFO] Error while setting up: "+err.Error(), 0)
		return
	}

	log.Println("[CRON] found " + strconv.Itoa(len(routines)) + " routines")

	for i := 0; i < len(routines); i++ {
		if routines[i].Status == true {
			for y := 0; y < len(routines[i].Payload); y++ {
				err = NewJob(
					routines[i].TriggerTime,
					routines[i].Payload[y].Device.Channel,
					routines[i].Payload[y].Device.Room,
					strconv.Itoa(routines[i].Payload[y].Value),
					routines[i].ID,
				)
				if err != nil {
					log.New(os.Stdout, "[SETUP] Error while setting up: "+err.Error(), 0)
				}
			}
			log.Println("[CRON] routine added: " + routines[i].Title)
		}
	}
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
			log.Println("[CRON] routine with ID: " + strconv.Itoa(id) + " removed")
		}
	}
}
