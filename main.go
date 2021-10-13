package main

import (
	"log"
	"os"
	"runtime"
	"shome-backend/api"
	"shome-backend/cron"
	param "shome-backend/flags"
)

func main() {
	log.Print("[INFO] REST API v0.1 - Nethcon - eHome")
	_ = cron.Cron("1", "1", "1", "1", "1", "1", true)

	if runtime.GOOS == "windows" {
		log.New(os.Stdout, "[WARN] Debug Mode", 0)
	} else {
		log.New(os.Stdout, "[INFO] Release Mode", 0)
	}

	param.HandleFlags()
	api.HandleRequests()

}
