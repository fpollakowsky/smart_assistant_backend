package add

import (
	"errors"
	"fmt"
	"home-ai-backend/models"
	"home-ai-backend/server/config"
	"home-ai-backend/server/worker"
	"log"
	"os"
	"strconv"
)

func Routine(routine models.Routine) error {
	result := config.DB.Create(&routine)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Println(result.RowsAffected)
		return errors.New("ADD_ROUTINE:: no rows affected")
	}

	for i := 0; i < len(routine.Payload); i++ {
		err := worker.NewJob(
			routine.TriggerTime,
			routine.Payload[i].Device.Channel,
			routine.Payload[i].Device.Room,
			strconv.Itoa(routine.Payload[i].Value),
			routine.ID,
		)
		if err != nil {
			log.New(os.Stdout, "[SETUP] Error while setting up: "+err.Error(), 0)
		}
	}

	return nil
}
