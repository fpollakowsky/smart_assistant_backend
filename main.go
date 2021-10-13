package main

import (
	"log"
	"os"
	"runtime"
	"shome-backend/api"
	"shome-backend/cron"
	param "shome-backend/flags"
	"shome-backend/mysql"
)

func main() {
	log.Print("[INFO] REST API v0.1 - Nethcon - eHome")

	initialize()

	if runtime.GOOS == "windows" {
		log.New(os.Stdout, "[WARN] Debug Mode", 0)
	} else {
		log.New(os.Stdout, "[INFO] Release Mode", 0)
	}

	param.HandleFlags()
	api.HandleRequests()

}

func initialize() {
	// remove all cron jobs
	_ = cron.Cron("1", "1", "1", "1", "1", "1", true)

	// add cron jobs from db
	routines, err := mysql.GetRoutines()
	if err != nil {
		log.New(os.Stdout, "[INFO] Error while setting up: "+err.Error(), 0)
		return
	}

	for i := 0; i < len(routines); i++ {
		if routines[i].Status == "1" {
			err := cron.Cron(
				routines[i].Min,
				routines[i].Hour,
				routines[i].Day,
				routines[i].Channel,
				routines[i].Room,
				routines[i].Payload,
				false)
			if err != nil {
				log.New(os.Stdout, "[SETUP] Error while setting up: "+err.Error(), 0)
			}
		}
	}
}
