package cron

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am running task.")
}

func AddCron() {
	gocron.Every(2).Seconds().Do(task)
	//gocron.Every(1).Day().At("16:10").Do(task)
}

func RemoveCron(task interface{}) {
	gocron.Remove(task)
}
