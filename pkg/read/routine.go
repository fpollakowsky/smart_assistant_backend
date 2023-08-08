package read

import (
	"errors"
	"home-ai-backend/models"
	"home-ai-backend/server/config"
)

func Routine() (error, []models.Routine) {
	var routines []models.Routine

	err := config.DB.
		Model(&models.Routine{}).
		Preload("Payload.Device").
		Preload("Payload").
		Find(&routines).Error
	if err != nil {
		return err, nil
	}

	if len(routines) == 0 {
		return errors.New("zero devices found"), nil
	}

	return nil, routines
}
