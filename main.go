package main

import (
	"github.com/jasonlvhit/gocron"
	"log"
	"os"
	"runtime"
	"shome-backend/api"
	"shome-backend/cron"
	"shome-backend/mqtt"
)

func main() {
	log.Print("[INFO] REST API v0.1 - Nethcon - eHome")

	client := mqtt.Connect()
	// mqtt2.Sub(c)
	// mqtt2.NewRequest(c, "test","test")

	mqtt.NewRequest(client, "test", "test", "test")

	if runtime.GOOS == "windows" {
		log.New(os.Stdout, "[WARN] Debug Mode", 0)
	} else {
		log.New(os.Stdout, "[INFO] Release Mode", 0)
	}

	api.HandleRequests()

	cron.AddCron()
	<-gocron.Start()
}
