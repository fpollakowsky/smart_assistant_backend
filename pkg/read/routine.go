package read

import (
	"errors"
	"shome-backend/models"
	"shome-backend/server/config"
)

func Routine() (error, []models.Routine) {
	var routines []models.Routine

	result := config.DB.Find(&routines)
	if result.Error != nil {
		return result.Error, nil
	}

	if result.RowsAffected == 0 {
		return errors.New("zero devices found"), nil
	}

	return nil, routines
}
