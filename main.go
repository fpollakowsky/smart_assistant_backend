package main

import (
	"log"
	"os"
	"runtime"
	"shome-backend/pkg/handler"
	"shome-backend/pkg/read"
	"shome-backend/server/flags"
	"shome-backend/server/mysql"
	"shome-backend/server/worker"
	"strconv"
)

func main() {
	log.Print("[INFO] REST API v1 - Nethcon - eHome")

	if runtime.GOOS == "windows" {
		log.New(os.Stdout, "[WARN] Debug Mode", 0)
	} else {
		log.New(os.Stdout, "[INFO] Release Mode", 0)
	}

	mysql.InitializeDatabase()
	flags.GetFlags()
	initializeCron()
	handler.HandleRequests()
}

func initializeCron() {

	// remove all cron jobs
	_ = worker.NewJob("* * * * * *", "1", "1", "1", true)

	// add cron jobs from db
	err, routines := read.Routine()
	if err != nil {
		log.New(os.Stdout, "[INFO] Error while setting up: "+err.Error(), 0)
		return
	}

	for i := 0; i < len(routines); i++ {
		if routines[i].Status == true {
			for y := 0; y < len(routines[i].Payload); y++ {
				err = worker.NewJob(
					routines[i].TriggerTime,
					routines[i].Payload[y].Device.Channel,
					routines[i].Payload[y].Device.Room,
					strconv.Itoa(routines[i].Payload[y].Value),
					false,
				)
				if err != nil {
					log.New(os.Stdout, "[SETUP] Error while setting up: "+err.Error(), 0)
				}
			}
		}
	}
}
