package main

import (
	"log"
	"os"
	"runtime"
	"shome-backend/api"
	param "shome-backend/flags"
)

func main() {
	log.Print("[INFO] REST API v0.1 - Nethcon - eHome")

	if runtime.GOOS == "windows" {
		log.New(os.Stdout, "[WARN] Debug Mode", 0)
	} else {
		log.New(os.Stdout, "[INFO] Release Mode", 0)
	}

	param.HandleFlags()
	api.HandleRequests()

}
