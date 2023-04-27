package remove

import (
	"errors"
	"shome-backend/models"
	"shome-backend/server/config"
	"shome-backend/server/worker"
)

func Routine(routine models.Routine) error {
	result := config.DB.Delete(&routine)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("REMOVE_ROUTINE:: no rows affected")
	}

	worker.RemoveWorker(routine.ID)

	return nil
}
