package add

import (
	"errors"
	"fmt"
	"shome-backend/models"
	"shome-backend/server/config"
	"shome-backend/server/worker"
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

	_ = worker.NewJob("59 12 * * 1-7", "1", "1", "1", false)

	// add cron jobs from db
	if routine.Status == true {
		for _, s := range routine.Payload {
			fmt.Println(s)
			/*
				err := worker.NewJob(
					routine.Min,
					routine.Hour,
					routine.Day,
					routine.Channel,
					routine.Room,
					routine.Payload,
					false,
				)

				if err != nil {
					log.New(os.Stdout, "[SETUP] Error while setting up: "+err.Error(), 0)
				}
			*/
		}

	}

	return nil
}
