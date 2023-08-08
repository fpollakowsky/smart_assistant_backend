package remove

import (
	"errors"
	"home-ai-backend/models"
	"home-ai-backend/server/config"
	"home-ai-backend/server/worker"
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
