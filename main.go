package main

import (
	"log"
	"os"
	"runtime"
	"shome-backend/pkg/handler"
	"shome-backend/server/flags"
	"shome-backend/server/mysql"
	"shome-backend/server/worker"
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
	worker.InitializeCron()
	handler.HandleRequests()
}
